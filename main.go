// The main package where the main function is defined
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	ipimVersion = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:     "ipim",
	Short:   "ipim is a command line tool for managing azure previlaged identity management",
	Long:    ipimLong(),
	Version: getCurrentVersion(),
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the PIM Roles",
}

var getRolesCmd = &cobra.Command{
	Use:   "ipim get roles",
	Short: "Get the PIM Roles",
}

func getCurrentVersion() string {
	return ipimVersion
}

func main() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(getRolesCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("IPIM Error: Pleae check the below error")
		fmt.Println(err)
	}

	rootCmd.Execute()
}
