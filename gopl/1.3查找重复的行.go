package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 从控制台输入中读取重复的行
// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Println(line)
// 		}
// 	}
// }

// 从一系列文件中读取重复的行
// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "%v", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Println(line)
// 		}
// 	}

// }

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	// NOTE: ignoring potential errors from input.Err()
// }


func main(){
	counts := make(map[string]int)
	for _, filename := os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf("err: %v", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	
	}
	for n, line := range counts {
		if n >1 {
			fmt.Println(line)
		}
	}
}