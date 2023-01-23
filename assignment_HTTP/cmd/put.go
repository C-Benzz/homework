/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "putHttp",
	Long:  `put: send a PUT request to a given URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		flag, err := cmd.Flags().GetStringSlice("json")
		if err != nil {
			log.Fatal(err)
		}
		putHttp(flag)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.PersistentFlags().StringSlice("json", []string{}, "Construct JSON body of the PUT request.")
}

func putHttp(body []string) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.NewRequest("PUT", "https://httpbin.org/put", &buf)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	// if res.Response.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", res.Body)
}
