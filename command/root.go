package main

import "github.com/spf13/cobra"

type RootCommand struct {
	Use string
	Short: "",
	Long: "",
}

var rootCommand = &cobra.Command{
	Use: "",
	Short: "",
	Long: "",
}

func Execute()  {
	rootCommand.Execute()
	rootCommand.Flags().StringVar()
}
