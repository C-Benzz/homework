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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletehttp",
	Long:  `Delete`,
	Run: func(cmd *cobra.Command, args []string) {
		flag, err := cmd.Flags().GetString("json")
		if err != nil {
			log.Fatal(err)
		}
		input := map[string]any{}
		json.Unmarshal([]byte(flag), &input)
		deleteHttp(input)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().String("json", "", "Construct JSON body of the Delete request.")
}

func deleteHttp(body map[string]any) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("DELETE", "https://httpbin.org/delete", &buf)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	res, err1 := client.Do(req)
	if err1 != nil {
		log.Fatal(err1)
	}
	var resp map[string]any
	json.NewDecoder(res.Body).Decode(&resp)
	fmt.Printf("%v\n", resp)
}
