/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "show current time",
	Long:  `show current time in two different format `,
	Run: func(cmd *cobra.Command, args []string) {
		if flag, _ := cmd.Flags().GetBool("str"); flag {
			fmt.Println(time.Now().Format("2006-01-02"))
			return
		}
		fmt.Println(time.Now())
		// fmt.Println("time called")
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().BoolP("str", "s", false, "show time in format in 2006-01-02")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
