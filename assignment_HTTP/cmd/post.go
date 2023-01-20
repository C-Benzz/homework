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
		postHttp()
		fmt.Println("post called")
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
}

type Post struct {
	Name string `json:"name"`
}

func postHttp() {
	p := Post{
		Name: "Hello",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("https://httpbin.org/post", "", &buf)
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)
}
