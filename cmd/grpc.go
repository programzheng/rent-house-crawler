/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/programzheng/rent-house-crawler/config"
	"github.com/programzheng/rent-house-crawler/internal/server"
	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		port := ""
		if len(args) == 0 {
			port = config.Cfg.GetString("grpc-port")
		} else {
			port = args[0]
		}
		server.RunGrpcServer(port)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
