// The main package where the main function is defined
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipim",
	Short: "ipim is a command line tool for managing azure previlaged identity management",
	// Long:
	//ipimLong,
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the PIM Roles",
}

func main() {

	rootCmd.AddCommand(getCmd)

	fmt.Println("Hello, World!")
	rootCmd.Execute()
}
