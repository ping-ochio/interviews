package main

import "fmt"

func main() {

	// Make a new empty map where key is "string"
	// and all values are "[]string"
	data := make(map[string][]string)

	// Insert every peer, key:value
	data["Peter"] = []string{"music", "swimming", "cook"}
	data["Jhon"] = []string{"dance", "run", "biking"}
	data["Mary"] = []string{"hunting", "programming", "computers"}

	// Get all keys
	for k, v := range data {

		// Print each key
		fmt.Println("User:", k)

		// each "v" value is an "[]string"
		for i, val := range v {
			fmt.Println("\t", i, val)

		}
	}
}
