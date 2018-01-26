// +build linux

package mount

import (
	"os"
	"path"
	"testing"

	"golang.org/x/sys/unix"
)

// nothing is propagated in or out
func TestSubtreePrivate(t *testing.T) ***REMOVED***
	tmp := path.Join(os.TempDir(), "mount-tests")
	if err := os.MkdirAll(tmp, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer os.RemoveAll(tmp)

	var (
		sourceDir   = path.Join(tmp, "source")
		targetDir   = path.Join(tmp, "target")
		outside1Dir = path.Join(tmp, "outside1")
		outside2Dir = path.Join(tmp, "outside2")

		outside1Path      = path.Join(outside1Dir, "file.txt")
		outside2Path      = path.Join(outside2Dir, "file.txt")
		outside1CheckPath = path.Join(targetDir, "a", "file.txt")
		outside2CheckPath = path.Join(sourceDir, "b", "file.txt")
	)
	if err := os.MkdirAll(path.Join(sourceDir, "a"), 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.MkdirAll(path.Join(sourceDir, "b"), 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.Mkdir(targetDir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.Mkdir(outside1Dir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.Mkdir(outside2Dir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***

	if err := createFile(outside1Path); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := createFile(outside2Path); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***

	// mount the shared directory to a target
	if err := Mount(sourceDir, targetDir, "none", "bind,rw"); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(targetDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// next, make the target private
	if err := MakePrivate(targetDir); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(targetDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// mount in an outside path to a mounted path inside the _source_
	if err := Mount(outside1Dir, path.Join(sourceDir, "a"), "none", "bind,rw"); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(path.Join(sourceDir, "a")); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// check that this file _does_not_ show in the _target_
	if _, err := os.Stat(outside1CheckPath); err != nil && !os.IsNotExist(err) ***REMOVED***
		t.Fatal(err)
	***REMOVED*** else if err == nil ***REMOVED***
		t.Fatalf("%q should not be visible, but is", outside1CheckPath)
	***REMOVED***

	// next mount outside2Dir into the _target_
	if err := Mount(outside2Dir, path.Join(targetDir, "b"), "none", "bind,rw"); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(path.Join(targetDir, "b")); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// check that this file _does_not_ show in the _source_
	if _, err := os.Stat(outside2CheckPath); err != nil && !os.IsNotExist(err) ***REMOVED***
		t.Fatal(err)
	***REMOVED*** else if err == nil ***REMOVED***
		t.Fatalf("%q should not be visible, but is", outside2CheckPath)
	***REMOVED***
***REMOVED***

// Testing that when a target is a shared mount,
// then child mounts propagate to the source
func TestSubtreeShared(t *testing.T) ***REMOVED***
	tmp := path.Join(os.TempDir(), "mount-tests")
	if err := os.MkdirAll(tmp, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer os.RemoveAll(tmp)

	var (
		sourceDir  = path.Join(tmp, "source")
		targetDir  = path.Join(tmp, "target")
		outsideDir = path.Join(tmp, "outside")

		outsidePath     = path.Join(outsideDir, "file.txt")
		sourceCheckPath = path.Join(sourceDir, "a", "file.txt")
	)

	if err := os.MkdirAll(path.Join(sourceDir, "a"), 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.Mkdir(targetDir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.Mkdir(outsideDir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***

	if err := createFile(outsidePath); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***

	// mount the source as shared
	if err := MakeShared(sourceDir); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(sourceDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// mount the shared directory to a target
	if err := Mount(sourceDir, targetDir, "none", "bind,rw"); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(targetDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// mount in an outside path to a mounted path inside the target
	if err := Mount(outsideDir, path.Join(targetDir, "a"), "none", "bind,rw"); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(path.Join(targetDir, "a")); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// NOW, check that the file from the outside directory is available in the source directory
	if _, err := os.Stat(sourceCheckPath); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
***REMOVED***

// testing that mounts to a shared source show up in the slave target,
// and that mounts into a slave target do _not_ show up in the shared source
func TestSubtreeSharedSlave(t *testing.T) ***REMOVED***
	tmp := path.Join(os.TempDir(), "mount-tests")
	if err := os.MkdirAll(tmp, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer os.RemoveAll(tmp)

	var (
		sourceDir   = path.Join(tmp, "source")
		targetDir   = path.Join(tmp, "target")
		outside1Dir = path.Join(tmp, "outside1")
		outside2Dir = path.Join(tmp, "outside2")

		outside1Path      = path.Join(outside1Dir, "file.txt")
		outside2Path      = path.Join(outside2Dir, "file.txt")
		outside1CheckPath = path.Join(targetDir, "a", "file.txt")
		outside2CheckPath = path.Join(sourceDir, "b", "file.txt")
	)
	if err := os.MkdirAll(path.Join(sourceDir, "a"), 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.MkdirAll(path.Join(sourceDir, "b"), 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.Mkdir(targetDir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.Mkdir(outside1Dir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.Mkdir(outside2Dir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***

	if err := createFile(outside1Path); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := createFile(outside2Path); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***

	// mount the source as shared
	if err := MakeShared(sourceDir); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(sourceDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// mount the shared directory to a target
	if err := Mount(sourceDir, targetDir, "none", "bind,rw"); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(targetDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// next, make the target slave
	if err := MakeSlave(targetDir); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(targetDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// mount in an outside path to a mounted path inside the _source_
	if err := Mount(outside1Dir, path.Join(sourceDir, "a"), "none", "bind,rw"); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(path.Join(sourceDir, "a")); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// check that this file _does_ show in the _target_
	if _, err := os.Stat(outside1CheckPath); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***

	// next mount outside2Dir into the _target_
	if err := Mount(outside2Dir, path.Join(targetDir, "b"), "none", "bind,rw"); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(path.Join(targetDir, "b")); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// check that this file _does_not_ show in the _source_
	if _, err := os.Stat(outside2CheckPath); err != nil && !os.IsNotExist(err) ***REMOVED***
		t.Fatal(err)
	***REMOVED*** else if err == nil ***REMOVED***
		t.Fatalf("%q should not be visible, but is", outside2CheckPath)
	***REMOVED***
***REMOVED***

func TestSubtreeUnbindable(t *testing.T) ***REMOVED***
	tmp := path.Join(os.TempDir(), "mount-tests")
	if err := os.MkdirAll(tmp, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer os.RemoveAll(tmp)

	var (
		sourceDir = path.Join(tmp, "source")
		targetDir = path.Join(tmp, "target")
	)
	if err := os.MkdirAll(sourceDir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	if err := os.MkdirAll(targetDir, 0777); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***

	// next, make the source unbindable
	if err := MakeUnbindable(sourceDir); err != nil ***REMOVED***
		t.Fatal(err)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(sourceDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()

	// then attempt to mount it to target. It should fail
	if err := Mount(sourceDir, targetDir, "none", "bind,rw"); err != nil && err != unix.EINVAL ***REMOVED***
		t.Fatal(err)
	***REMOVED*** else if err == nil ***REMOVED***
		t.Fatalf("%q should not have been bindable", sourceDir)
	***REMOVED***
	defer func() ***REMOVED***
		if err := Unmount(targetDir); err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
	***REMOVED***()
***REMOVED***

func createFile(path string) error ***REMOVED***
	f, err := os.Create(path)
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	f.WriteString("hello world!")
	return f.Close()
***REMOVED***
