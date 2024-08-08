package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{
		Use:   "greeter",
		Short: "Greeter basic CLI",
		Long:  "Greeter is a friendly command line application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Greater!")
		},
	}

	greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "Greeter basic CLI",
		Long:  "Greeter is a friendly command line application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!", args[0])
		},
	}

	rootCmd.AddCommand(greetCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
