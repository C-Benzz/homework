package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "httpcli",
	Short: "Http cli",
	Long:  `HTTP Client Command-line program`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// NOTE: What is this toggle flag?
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
