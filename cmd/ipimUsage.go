package cmd

// Custom usage message with ASCII art aligned to the left
func IpimUsage() string {
	return `
  ___    ___   ___   __  __
 |_ _|  | _ \ |_ _| |  \/  |
  | |   |  _/  | |  | |\/| | 
 |___|  |_|   |___| |_|  |_|
	
Usage : ipim [global options] <subcommands> [arguments...]

IPIM is a tool to help elevate yourself to the azure pim roles. 
It's a thin wrapper around azure CLI and Graph API to help you elevate your permissions to perform your tasks.

You must have the Azure CLI installed and configured to use this tool.
Login to Azure using the command 'az login' before using this tool.

Main Commands:

show 		List all the available roles
activate        Activate the roles
config	        Generate the config YAML file for all the available roles, and later can be altered to assign the roles
admin           PIM Admin Commands
status 	        Get the current active roles

Global Options (Use these options with the main commands/subcommands):

-h / --help 		List all the available commands
-v / --version 		Get the version of the IPIM
-i / --info 		Get information about IPIM


For more information, please open a support ticket on the CET kanban board.
`
}
