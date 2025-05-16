//go:build windows

package cmd

import (
	"syscall"
	"unsafe"
)

func getDiskUsage(path string) (uint64, uint64, uint64, error) {
	// Variables to store disk space info: available, total, and free bytes
	lpFree := uint64(0)
	lpTotal := uint64(0)
	lpAvailable := uint64(0)

	// Load Windows kernel32.dll and find the GetDiskFreeSpaceExW procedure
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	getDiskFreeSpaceExW := kernel32.MustFindProc("GetDiskFreeSpaceExW")

	// Convert path string to UTF-16 pointer for Windows API call
	pathPtr, _ := syscall.UTF16PtrFromString(path)

	// Call GetDiskFreeSpaceExW Windows API to get disk space info
	_, _, callErr := getDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(pathPtr)),
		uintptr(unsafe.Pointer(&lpAvailable)),
		uintptr(unsafe.Pointer(&lpTotal)),
		uintptr(unsafe.Pointer(&lpFree)),
	)

	// Return error if the API call failed
	if callErr != syscall.Errno(0) {
		return 0, 0, 0, callErr
	}

	// Calculate used space = total - free
	used := lpTotal - lpFree
	return lpTotal, lpFree, used, nil
}
