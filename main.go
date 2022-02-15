package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/c-pro/kvcli/pkg/kv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tx := kv.NewTX(nil)
	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}

		text := scanner.Text()
		if len(text) == 0 || text == "exit" || text == "quit" {
			return
		}

		parts := []string{}
		// I assume arguments do not contain spaces
		for _, p := range strings.Split(text, " ") {
			// filter out empty strings
			// that string.Split does when encounters extra separators
			if len(p) > 0 {
				parts = append(parts, p)
			}
		}

		if len(parts) == 0 {
			continue
		}

		cmd := strings.ToLower(parts[0])

		switch cmd {
		case "begin":
			tx = kv.NewTX(tx)
		case "commit":
			tx.Parent = nil
		case "rollback":
			if tx.Parent == nil {
				fmt.Println("no transaction")
				continue
			}
			tx = tx.Parent
		case "get":
			if len(parts) != 2 {
				fmt.Println("GET should have exactly one argument")
				continue
			}

			val := tx.Get(parts[1])
			if val == nil {
				fmt.Println("key not set")
				continue
			}

			fmt.Println(*val)
		case "set":
			if len(parts) != 3 {
				fmt.Println("SET should have exactly two arguments")
				continue
			}

			tx.Set(parts[1], parts[2])
		case "delete":
			if len(parts) != 2 {
				fmt.Println("DELETE should have exactly one argument")
				continue
			}

			tx.Delete(parts[1])
		}
	}
}
