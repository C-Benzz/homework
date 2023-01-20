/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "getHttp",
	Long:  `get: send a GET request to a given URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		getHttp()
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getHttp() {
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
