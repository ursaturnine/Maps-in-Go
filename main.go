package main

import "fmt"

// declaring a map
// [<keys of type ...>]<values of type...>
// access values in map with <variable_name>[<key_name>]
func main() {
	//option 2 - creates empty amp
	// var colors map[string]string

	//option 3 -creates empty map
	colors := make(map[int]string)

	// option 1
	// colors := map[string]string{
	// 	"red":   "#ff000",
	// 	"green": "#gr33n",
	// }

	colors[10] = "#ffffff"

	// delete a key
	// delete(<variable_name>, <key_name>)
	delete(colors, 10)

	fmt.Println(colors)
}
