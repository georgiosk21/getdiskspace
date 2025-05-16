//go:build darwin || linux

package cmd

import "syscall"

func getDiskUsage(path string) (uint64, uint64, uint64, error) {
	// Declare a Statfs_t struct to hold filesystem statistics
	var stat syscall.Statfs_t

	// Get filesystem statistics for the given path
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0, 0, 0, err // Return error if syscall fails
	}

	// Calculate total, free, and used bytes
	total := stat.Blocks * uint64(stat.Bsize) // Total data blocks * block size
	free := stat.Bfree * uint64(stat.Bsize)   // Free blocks available to non-root users
	used := total - free                      // Used space = total - free

	return total, free, used, nil
}
