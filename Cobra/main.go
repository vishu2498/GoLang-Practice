//'Cobra' Tool is used for building CLI applications for GoLang.
//It is basically used for making custom commands for performing some operations.
//It has been used for many huge applications like Kubernetes & Docker etc.

package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {
	//We define a variable that uses 'Command' struct from 'cobra' package.
	//This 'Command' Struct contains all the fields that are useful for running our custom commands.
	//A command is a command for our application. E.g.  'go run ...' - 'run' is the command.
	cmd := &cobra.Command{
		Use: "egmessage", //the custom command that we need to fire make our application run/execute something
		//Aliases are slices of strings which means that using aliases, many commands can be fired instead of just one.
		Short: "Custom message from Vishu", //It is the short description shown in the 'help' output.
	}

	cmd.AddCommand(printCustomMsg()) //This calls the function that will return the information that will be displayed on the console.

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
