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

var outputFile string

// checkcontentCmd represents the checkcontent command
var checkcontentCmd = &cobra.Command{
	Use:   "checkcontent -file <your file name>",
	Short: "Fetch the content of a given URL and store it in a file. ",
	Long: `You can directly get the content or store it in a file :
	./getfromurl checkcontent
	./getfromurl checkcontent -file <your file name>`,
	Args: cobra.ExactArgs(1),

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

		content := utils.CheckContent(resp)

		if outputFile != "" {
			err := utils.WriteInFile(outputFile, content)
			if err != nil {
				fmt.Printf("error writing to file\n")
				os.Exit(1)
			}
			fmt.Printf("Content written to %s\n", outputFile)
		} else {
			fmt.Print("Print content:\n")
			fmt.Print(content)
		}
	},
}

func init() {
	checkcontentCmd.Flags().StringVarP(&outputFile, "file", "f", "", "File to write the content to")
	rootCmd.AddCommand(checkcontentCmd)
}
