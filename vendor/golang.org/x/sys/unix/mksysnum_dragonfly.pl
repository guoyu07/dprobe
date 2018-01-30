#!/usr/bin/env perl
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
#
# Generate system call table for DragonFly from master list
# (for example, /usr/src/sys/kern/syscalls.master).

use strict;

if($ENV***REMOVED***'GOARCH'***REMOVED*** eq "" || $ENV***REMOVED***'GOOS'***REMOVED*** eq "") ***REMOVED***
	print STDERR "GOARCH or GOOS not defined in environment\n";
	exit 1;
***REMOVED***

my $command = "mksysnum_dragonfly.pl " . join(' ', @ARGV);

print <<EOF;
// $command
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build $ENV***REMOVED***'GOARCH'***REMOVED***,$ENV***REMOVED***'GOOS'***REMOVED***

package unix

const (
EOF

while(<>)***REMOVED***
	if(/^([0-9]+)\s+STD\s+(***REMOVED*** \S+\s+(\w+).*)$/)***REMOVED***
		my $num = $1;
		my $proto = $2;
		my $name = "SYS_$3";
		$name =~ y/a-z/A-Z/;

		# There are multiple entries for enosys and nosys, so comment them out.
		if($name =~ /^SYS_E?NOSYS$/)***REMOVED***
			$name = "// $name";
		***REMOVED***
		if($name eq 'SYS_SYS_EXIT')***REMOVED***
			$name = 'SYS_EXIT';
		***REMOVED***

		print "	$name = $num;  // $proto\n";
	***REMOVED***
***REMOVED***

print <<EOF;
)
EOF