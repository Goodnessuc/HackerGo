/*
Copyright © 2022 NAME HERE <ukejegoodness599@gmail.com>

*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "HackerGo",
	Short: "HackerGo is a CLI tool that aids developers stay productive and up-to-date with the technology world",
	Long: `
	HackerGo is a CLI tool that aids developers stay productive and up-to-date with the technology world.
	You'll find HackerGo useful if you're struggling with productivity; 
	You can use HackerGo to read tech news seamlessly from your CLI, and seeking for Job opportunities.
	
	You don't need to leave your IDE/Editor; The world is waiting; Keep hacking hacker 🦹🏾‍ 🦸🏼‍'`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.HackerGo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
