package main

func ipimLong() string {
	return `Welcome to the IPIM!
	  ___    ___   ___   __  __
	 |_ _|  | _ \ |_ _| |  \/  |
	  | |   |  _/  | |  | |\/| | 
	 |___|  |_|   |___| |_|  |_|
	
	Usage : ipim [global options] <subcommands> [arguments...]
	
	IPIM is a tool to help elevate yourself to the azure pim roles. 
	its a thin wrapper around azure cli and Graph API to help you elevate your permissions to perform your tasks.
	
	You must have the azure cli installed and configured to use this tool,
	login to azure using the command 'az login' before using this tool.
	
	Main Commands:
	
	get-pim 		Get the PIM Roles
	config		 	Generate the config YAML file for all the available roles, and later can be altered to assign the roles
	admin 			PIM Admin Commands
	status 			Get the current active roles
	
	Global Options (Use these options with the main commands/sub commands):
	
	-h / --help 		List all the available commands
	-v / --version 		Get the version of the IPIM
	-i / --info 		Get the information about IPIM
	
	
	For more information, please open support ticket at CET kanban board.`
}
