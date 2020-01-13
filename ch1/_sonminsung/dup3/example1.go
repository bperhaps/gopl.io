package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dupfiles := []string{}
FILE:
	for _, filename := range os.Args[1:] {
		counts := make(map[string]int)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1 {
				dupfiles = append(dupfiles, filename)
				continue FILE
			}
		}
	}

	for _, dupfile := range dupfiles {
		fmt.Println(dupfile)
	}
}
