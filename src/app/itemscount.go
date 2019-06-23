package main

import "fmt"

func itemscount(c *Config) {

	item := len(c.URLRanges) + len(c.IPRanges)
	fmt.Println(item)
	wg.Add(item)
}
