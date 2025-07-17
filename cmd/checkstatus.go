/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"getfromurl/utils"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// checkstatusCmd represents the checkstatus command
var checkstatusCmd = &cobra.Command{
	Use:   "checkstatus",
	Short: "Check status code of a given URL",
	Long:  `./getfromurl checkstatus`,
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		client := http.Client{
			Timeout: 1 * time.Second,
		}

		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		utils.CheckStatus(resp.StatusCode)
	},
}

func init() {
	rootCmd.AddCommand(checkstatusCmd)
}
