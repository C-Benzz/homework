/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

type AutoGenerated struct {
	Args    map[string]string `json:"args"`
	Headers struct {
		AcceptEncoding string `json:"Accept-Encoding"`
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "getHttp",
	Long:  `get: send a GET request to a given URL.`,
	Run: func(cmd *cobra.Command, args []string) {

		flagQuery, err := cmd.Flags().GetStringSlice("query")
		if err != nil {
			log.Fatal(err)
		}
		flagsQuery := strings.Join(flagQuery, "&")

		flagHeader, err1 := cmd.Flags().GetStringSlice("header")
		if err1 != nil {
			log.Fatal(err1)
		}
		flagsHeader := strings.Join(flagHeader, "&")

		url := "https://httpbin.org/get"

		if flagsQuery != "" {
			getHttpWithFlagQuery(url, flagsQuery)
		} else if flagsHeader != "" {
			getHttpWithFlagHeader(url, flagsHeader)
		} else {
			getHttp(url)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringSlice("query", []string{}, "Query parameters of the GET request.")
	getCmd.PersistentFlags().StringSlice("header", []string{}, "Header")
}

func getHttp(url string) {
	res, err := http.Get(url)
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

func getHttpWithFlagQuery(url string, flag string) {
	urlGet := fmt.Sprintf("%s?%s", url, flag)
	res, err := http.Get(urlGet)
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
	var data AutoGenerated

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Printf("%s\n", data.Args)
}

func getHttpWithFlagHeader(url string, flag string) {
	urlGet := fmt.Sprintf("%s?%s", url, flag)
	res, err := http.Get(urlGet)
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
	var data AutoGenerated

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Printf("%#v\n", data.Headers)
}
