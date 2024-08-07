package serial

import (
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetCommState    = modkernel32.NewProc("GetCommState")
	procSetCommState    = modkernel32.NewProc("SetCommState")
	procGetCommTimeouts = modkernel32.NewProc("GetCommTimeouts")
	procSetCommTimeouts = modkernel32.NewProc("SetCommTimeouts")
)

func GetCommState(handle syscall.Handle, dcb *c_DCB) (err error) {
	//r1, _, e1 := syscall.Syscall(procGetCommState.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(dcb)), 0)
	r1, _, e1 := syscall.SyscallN(procGetCommState.Addr(), uintptr(handle), uintptr(unsafe.Pointer(dcb)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetCommState(handle syscall.Handle, dcb *c_DCB) (err error) {
	//r1, _, e1 := syscall.Syscall(procSetCommState.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(dcb)), 0)
	r1, _, e1 := syscall.SyscallN(procSetCommState.Addr(), uintptr(handle), uintptr(unsafe.Pointer(dcb)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetCommTimeouts(handle syscall.Handle, timeouts *c_COMMTIMEOUTS) (err error) {
	//r1, _, e1 := syscall.Syscall(procGetCommTimeouts.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)
	r1, _, e1 := syscall.SyscallN(procGetCommTimeouts.Addr(), uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetCommTimeouts(handle syscall.Handle, timeouts *c_COMMTIMEOUTS) (err error) {
	//r1, _, e1 := syscall.Syscall(procSetCommTimeouts.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)
	r1, _, e1 := syscall.SyscallN(procSetCommTimeouts.Addr(), uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
