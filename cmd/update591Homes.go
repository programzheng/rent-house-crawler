/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/programzheng/rent-house-crawler/internal/job"
	"github.com/spf13/cobra"
)

// update591HomesCmd represents the update591Homes command
var update591HomesCmd = &cobra.Command{
	Use:   "update591Homes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		job.UpsertRent591Homes()
	},
}

func init() {
	rootCmd.AddCommand(update591HomesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// update591HomesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// update591HomesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
