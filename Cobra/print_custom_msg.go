package main

import (
	"github.com/spf13/cobra"
)

//This function will be called by 'cobra' package since we have defined this with "AddCommand()"
func printCustomMsg() *cobra.Command {
	return &cobra.Command{
		Use: "egmessage", //our custom command

		//'Run' is the actual work function. Most commands only implement this field.
		//'RunE' does the same work of 'Run' but returns an error
		RunE: func(cmd *cobra.Command, args []string) error { //the return type of the anonymous function is 'error' due to usage of 'RunE'
			cmd.Println("This is a custom message from Vishu fired via CLI") //custom print message on the console
			return nil
		},
	}
}
