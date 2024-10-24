/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package show

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showAllCmd = &cobra.Command{
	Use:   "all",
	Short: "List all available roles",
	Long: `
Usage: ipim [global options] show all [options] [args]

This command has subcommands for advanced PIM management.

You can use this command to list all the eligible role assignments using 
Microsoft Entra PIM for Azure resources, Entra ID Roles and Azure Entra ID Groups. 
With Microsoft Entra PIM, your end users must activate an eligible role assignment 
to get permission to perform certain actions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showAll called")
	},
}

// azureRolesCmd represents the azureRoles command
var azureRolesCmd = &cobra.Command{
	Use:   "resources",
	Short: "List the available PIM roles for azure resources",
	Long: `
Usage: ipim [global options] show resources [options] [args]

This command has subcommands for advanced PIM management.

You can use this command to list the eligible role assignments using Microsoft Entra PIM for Azure resources. 
With Microsoft Entra PIM, your end users must activate an eligible role assignment 
to get permission to perform certain actions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("azureRoles called")
	},
}

var entraRolesCmd = &cobra.Command{
	Use:   "entra",
	Short: "List PIM roles for EntraID/AzureAD Roles",
	Long: `
Usage: ipim [global options] show entra [options] [args]

This command has subcommands for advanced PIM management.

You can use this command to list the eligible role assignments using Microsoft Entra PIM for Entra ID Roles. 
With Microsoft Entra PIM, your end users must activate an eligible role assignment 
to get permission to perform certain actions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("entraRoles called")
	},
}

var adGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List availalbe PIM roles for AD groups",
	Long: `
Usage: ipim [global options] show groups [options] [args]

This command has subcommands for advanced PIM management.

You can use this command to list the eligible role assignments using Microsoft Entra PIM for Azure Entra ID Groups. 
With Microsoft Entra PIM, your end users must activate an eligible role assignment 
to get permission to perform certain actions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adGroups called")
	},
}

func init() {
	ShowCmd.AddCommand(showAllCmd)
	ShowCmd.AddCommand(azureRolesCmd)
	ShowCmd.AddCommand(entraRolesCmd)
	ShowCmd.AddCommand(adGroupsCmd)
}
