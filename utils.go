package main

import (
	"encoding/json"
	"fmt"

	"github.com/gookit/color"
)

func printColoredHeaders(headers map[string][]string) {
	// Print headers with colors
	for key, values := range headers {
		for _, value := range values {
			color.Bold.Printf("%s: ", key)
			color.Cyan.Printf("%s\n", value)
		}
	}
}

// printColoredBody prints the colored JSON response body.
func printColoredBody(body []byte) {
	// Try parsing the body as a single JSON object
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err == nil {
		printColoredJSONObject(data)
		return
	}

	// Try parsing the body as an array of JSON objects
	var arrayData []map[string]interface{}
	err = json.Unmarshal(body, &arrayData)
	if err == nil {
		for _, obj := range arrayData {
			printColoredJSONObject(obj)
			fmt.Println()
		}
		return
	}

	// If parsing fails, print the raw body as plain text
	fmt.Println(string(body))
}

func printColoredJSONObject(obj map[string]interface{}) {
	keyColor := color.New(color.FgYellow)
	valueColor := color.New(color.FgCyan)

	for key, value := range obj {
		keyColor.Printf("%s: ", key)
		valueColor.Println(value)
	}
}
