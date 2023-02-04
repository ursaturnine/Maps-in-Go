package main

import "fmt"

// declaring a map
// [<keys of type ...>]<values of type...>
// access values in map with <variable_name>[<key_name>]
func main() {
	//option 2 - creates empty map
	// var colors map[string]string

	//option 3 -creates empty map
	// colors := make(map[int]string)

	// option 1
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#gr33n",
		"white": "#ffffff",
	}

	// delete a key
	// delete(<variable_name>, <key_name>)
	// delete(colors, 10)

	printMap(colors)
}

// iterate over a map
// func <func-name>(<arg_name> map[<type of map>]<values of map>) {
// for <key>, <value> := range <arg_name> {
//
// }
// }
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
