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

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "postHttp",
	Long:  `post: send a POST request to a given URL`,
	Run: func(cmd *cobra.Command, args []string) {
		flagJSON, err := cmd.Flags().GetString("json")
		if err != nil {
			log.Fatal(err)
		}
		input := map[string]interface{}{}
		json.Unmarshal([]byte(flagJSON), &input) //change json to map
		postHttp(input)
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.PersistentFlags().String("json", "", "Construct JSON body of the POST request.")
}

type Post struct {
	Name string `json:"name"`
}

var p = fmt.Println

func postHttp(body map[string]interface{}) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		log.Fatal(err)
	}
	p(body)
	res, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatal(err)
	}
	var resp map[string]any

	json.NewDecoder(res.Body).Decode(&resp)

	fmt.Printf("%s\n", resp)
}
