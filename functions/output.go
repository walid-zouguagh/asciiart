package functions

import (
	"fmt"
)

func Output(verifiedInput []string, contentFileStandard map[rune][]string) {
	for _, word := range verifiedInput {
		line := ""
		i := 0
		for i < 8 {
			for _, char := range word {
				line += contentFileStandard[char][i]
			}
			i++
			line += "\n"
		}
		fmt.Print(line)
	}
}

// hello world
