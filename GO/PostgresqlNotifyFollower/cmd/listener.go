/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"Notifier/helpers"
	"Notifier/listener"
	"Notifier/listener/db"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Version = "0.1-alpha"

func preRun(cmd *cobra.Command, args []string) {
	if helpers.IsLock() {
		fmt.Fprintln(os.Stderr, "Listener is currently running")
		os.Exit(0)
	}
}

func run(cmd *cobra.Command, args []string) {
	config := listener.LoadConfig("/home/m.moosavi/Desktop/Projects/docs/GO/PostgresqlNotifyFollower/conf.ini")
	helpers.Lock()
	listener.Run(config, db.Connect(config))
}

// listenerCmd represents the listener command
var listenerCmd = &cobra.Command{
	Use:     "listener",
	Short:   "Postgresql listener",
	Long:    `Postgresql main listener here you can listen to a postgresql event and do something`,
	Version: Version,
	PreRun: func(cmd *cobra.Command, args []string) {
		preRun(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		PostStop()
	},
}

var listenerStartCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run", "enable"},
	Short:   "Postgresql start listener",
	Long:    `Postgresql main listener here you can listen to a postgresql event and do something`,
	Version: Version,
	PreRun: func(cmd *cobra.Command, args []string) {
		preRun(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		PostStop()
	},
}

func PostStop() {
	helpers.UnLock()
}

var listenerStatusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Get status of notifier",
	Long:    "",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		if helpers.IsLock() {
			fmt.Fprintln(os.Stdout, "state is Running")
			return
		}
		fmt.Fprintln(os.Stdout, "state is Stopped")
	},
}

func init() {
	rootCmd.AddCommand(listenerCmd)
	listenerCmd.AddCommand(listenerStartCmd)
	listenerCmd.AddCommand(listenerStatusCmd)
}

// Here you will define your flags and configuration settings.

// Cobra supports Persistent Flags which will work for this command
// and all subcommands, e.g.:
// listenerCmd.PersistentFlags().String("foo", "", "A help for foo")

// Cobra supports local flags which will only run when this command
// is called directly, e.g.:
// listenerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
