<<<<<<< HEAD
// +build solaris
// +build !appengine
=======
//go:build solaris && !appengine
// +build solaris,!appengine
>>>>>>> grw_branch

package isatty

import (
	"golang.org/x/sys/unix"
)

// IsTerminal returns true if the given file descriptor is a terminal.
<<<<<<< HEAD
// see: http://src.illumos.org/source/xref/illumos-gate/usr/src/lib/libbc/libc/gen/common/isatty.c
func IsTerminal(fd uintptr) bool {
	var termio unix.Termio
	err := unix.IoctlSetTermio(int(fd), unix.TCGETA, &termio)
=======
// see: https://src.illumos.org/source/xref/illumos-gate/usr/src/lib/libc/port/gen/isatty.c
func IsTerminal(fd uintptr) bool {
	_, err := unix.IoctlGetTermio(int(fd), unix.TCGETA)
>>>>>>> grw_branch
	return err == nil
}

// IsCygwinTerminal return true if the file descriptor is a cygwin or msys2
// terminal. This is also always false on this environment.
func IsCygwinTerminal(fd uintptr) bool {
	return false
}
