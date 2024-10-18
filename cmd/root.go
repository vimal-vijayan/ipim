package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const (
	ipimVersion = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:     "ipim",
	Short:   "IPIM is a CLI tool to manage the PIM roles in Azure",
	Long:    "IPIM is a CLI tool to manage the PIM roles in Azure",
	Version: ipimVersion,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(ipimUsage())
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("The command executed with error: ", err)
	}
}
