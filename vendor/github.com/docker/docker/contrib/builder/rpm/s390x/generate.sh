#!/usr/bin/env bash
set -e

# This file is used to auto-generate Dockerfiles for making rpms via 'make rpm'
#
# usage: ./generate.sh [versions]
#    ie: ./generate.sh
#        to update all Dockerfiles in this directory
#    or: ./generate.sh centos-7
#        to only update centos-7/Dockerfile
#    or: ./generate.sh fedora-newversion
#        to create a new folder and a Dockerfile within it

cd "$(dirname "$(readlink -f "$BASH_SOURCE")")"

versions=( "$@" )
if [ $***REMOVED***#versions[@]***REMOVED*** -eq 0 ]; then
	versions=( */ )
fi
versions=( "$***REMOVED***versions[@]%/***REMOVED***" )

for version in "$***REMOVED***versions[@]***REMOVED***"; do
	echo "$***REMOVED***versions[@]***REMOVED***"
	distro="$***REMOVED***version%-****REMOVED***"
	suite="$***REMOVED***version##*-***REMOVED***"
	case "$distro" in
		*opensuse*)
		from="opensuse/s390x:tumbleweed"
		;;
	*clefos*)
		from="sinenomine/$***REMOVED***distro***REMOVED***"
		;;
	*)
		echo No appropriate or supported image available.
		exit 1
		;;
    esac
	installer=yum

	mkdir -p "$version"
	echo "$version -> FROM $from"
	cat > "$version/Dockerfile" <<-EOF
		#
		# THIS FILE IS AUTOGENERATED; SEE "contrib/builder/rpm/s390x/generate.sh"!
		#

		FROM $from

	EOF

	echo >> "$version/Dockerfile"

	extraBuildTags=''
	runcBuildTags=

	case "$from" in
		*clefos*)
			# Fix for RHBZ #1213602 + get "Development Tools" packages dependencies
			echo 'RUN touch /var/lib/rpm/* && yum groupinstall -y "Development Tools"' >> "$version/Dockerfile"
			;;
		*opensuse*)
			echo "RUN zypper ar https://download.opensuse.org/ports/zsystems/tumbleweed/repo/oss/ tumbleweed" >> "$version/Dockerfile"
			# get rpm-build and curl packages and dependencies
			echo 'RUN zypper --non-interactive install ca-certificates* curl gzip rpm-build' >> "$version/Dockerfile"
			;;
		*)
			echo No appropriate or supported image available.
			exit 1
			;;
	esac

	packages=(
		btrfs-progs-devel # for "btrfs/ioctl.h" (and "version.h" if possible)
		device-mapper-devel # for "libdevmapper.h"
		glibc-static
		libseccomp-devel # for "seccomp.h" & "libseccomp.so"
		libselinux-devel # for "libselinux.so"
		pkgconfig # for the pkg-config command
		selinux-policy
		selinux-policy-devel
		sqlite-devel # for "sqlite3.h"
		systemd-devel # for "sd-journal.h" and libraries
		tar # older versions of dev-tools do not have tar
		git # required for containerd and runc clone
		cmake # tini build
		vim-common # tini build
	)

	case "$from" in
		*clefos*)
			extraBuildTags+=' seccomp'
			runcBuildTags="seccomp selinux"
			;;
		*opensuse*)
			packages=( "$***REMOVED***packages[@]/libseccomp-devel***REMOVED***" )
			runcBuildTags="selinux"
			;;
		*)
			echo No appropriate or supported image available.
			exit 1
			;;
	esac

	case "$from" in
		*clefos*)
			# Same RHBZ fix needed on all yum lines
			echo "RUN touch /var/lib/rpm/* && $***REMOVED***installer***REMOVED*** install -y $***REMOVED***packages[*]***REMOVED***" >> "$version/Dockerfile"
			;;
		*opensuse*)
			packages=( "$***REMOVED***packages[@]/btrfs-progs-devel/libbtrfs-devel***REMOVED***" )
			packages=( "$***REMOVED***packages[@]/pkgconfig/pkg-config***REMOVED***" )
			packages=( "$***REMOVED***packages[@]/vim-common/vim***REMOVED***" )

			packages+=( systemd-rpm-macros ) # for use of >= opensuse:13.*

			# use zypper
			echo "RUN zypper --non-interactive install $***REMOVED***packages[*]***REMOVED***" >> "$version/Dockerfile"
			;;
		*)
			echo No appropriate or supported image available.
			exit 1
			;;
	esac

	echo >> "$version/Dockerfile"

	awk '$1 == "ENV" && $2 == "GO_VERSION" ***REMOVED*** print; exit ***REMOVED***' ../../../../Dockerfile.s390x >> "$version/Dockerfile"
	echo 'RUN curl -fsSL "https://golang.org/dl/go$***REMOVED***GO_VERSION***REMOVED***.linux-s390x.tar.gz" | tar xzC /usr/local' >> "$version/Dockerfile"
	echo 'ENV PATH $PATH:/usr/local/go/bin' >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	echo 'ENV AUTO_GOPATH 1' >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	# print build tags in alphabetical order
	buildTags=$( echo "selinux $extraBuildTags" | xargs -n1 | sort -n | tr '\n' ' ' | sed -e 's/[[:space:]]*$//' )

	echo "ENV DOCKER_BUILDTAGS $buildTags" >> "$version/Dockerfile"
	echo "ENV RUNC_BUILDTAGS $runcBuildTags" >> "$version/Dockerfile"
	# TODO: Investigate why "s390x-linux-gnu-gcc" is required
	echo "RUN ln -s /usr/bin/gcc /usr/bin/s390x-linux-gnu-gcc" >> "$version/Dockerfile"
done
