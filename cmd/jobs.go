// Package cmd /*
package cmd

import (
	"Hacker-Go/scraper"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// randomCmd.PersistentFlags().String("term", "", "A search term for a dad joke.")

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You didn't specify the additional arguments\nUse the --help flag for comprehensive help on how to use this tool")
		}

		if args[0] == "all" {
			scraper.AllJobsFromPage()
		} else {
			numberStringConv, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalln("Please specify a number from 1 - 30")
			}
			if numberStringConv > 30 {
				log.Fatalln("Please specify a number from 1 - 30")
			} else {
				scraper.NumberOfJobs(numberStringConv)
			}

		}

		//fmt.Println("jobs called")
	},
}

func init() {
	rootCmd.AddCommand(jobsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
