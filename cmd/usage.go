package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var path string

var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Check disk usage for a given path",
	Run: func(cmd *cobra.Command, args []string) {
		total, free, used, err := getDiskUsage(path)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		fmt.Printf("\nDisk Usage for %s:\n", path)
		fmt.Println("+--------+----------------+")
		fmt.Println("| Field  | Value          |")
		fmt.Println("+--------+----------------+")
		fmt.Printf("| Total  | %10.2f GB |\n", float64(total)/1e9)
		fmt.Printf("| Used   | %10.2f GB |\n", float64(used)/1e9)
		fmt.Printf("| Free   | %10.2f GB |\n", float64(free)/1e9)
		fmt.Printf("| Usage  | %9.2f %% |\n", float64(used)/float64(total)*100)
		fmt.Println("+--------+----------------+")
	},
}

func init() {
	rootCmd.AddCommand(usageCmd)
	usageCmd.Flags().StringVarP(&path, "path", "p", "/", "Path to check disk usage")
}
