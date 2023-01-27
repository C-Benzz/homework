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
		// NOTE: Will this function work without --json?
		flag, err := cmd.Flags().GetString("json")
		if err != nil {
			log.Fatal(err)
		}

		input := map[string]any{}
		json.Unmarshal([]byte(flag), &input)
		putHttp(input)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.PersistentFlags().String("json", "", "Construct JSON body of the PUT request.")
}

func putHttp(body map[string]any) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "https://httpbin.org/put", &buf)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	res, err1 := client.Do(req) // NOTE: You don't need to declare another error (err1)
	if err1 != nil {
		log.Fatal(err1)
	}
	var resp map[string]any
	json.NewDecoder(res.Body).Decode(&resp) // NOTE: Should handle the error returned from the Decode function
	fmt.Printf("%v\n", resp)
}
