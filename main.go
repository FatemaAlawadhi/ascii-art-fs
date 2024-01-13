package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ascii-art-fs/pkg"
)

func main() {
	// Read the input
	input := os.Args
	if len(input) > 3 {
		fmt.Println("Please Enter Two Arguments Only")
	} else if len(input) < 2 {
		fmt.Println("Please Enter An Argument")
	} else {
		Text := string(input[1])
		TextArray := strings.Split(Text, "\\n")
		var FileType string
		if len(input) == 2 {
			FileType = "standard.txt"
		} else {
			if input[2] == "standard" {
				FileType = "standard.txt"
			} else if input [2] == "shadow" {
				FileType = "shadow.txt"
			} else if input[2] == "thinkertoy" {
				FileType = "thinkertoy.txt"
			}else {
				FileType = "standard.txt"
			}
		}
		
		// Open the file
		file, err := os.Open(FileType)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// Create a scanner for the file
		scanner := bufio.NewScanner(file)
		// Read each line of the file
		var Banner []string
		for scanner.Scan() {
			Banner = append(Banner, scanner.Text())
		}
		// To Print
		pkg.PrintString(Banner, TextArray)
	}
}
