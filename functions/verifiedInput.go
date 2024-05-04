package functions

import (
	"log"
)

func VerifiedInput(arguments []string) []string {
	/****************************************/
	/******* Verified Input the User ********/
	/****************************************/
	input := arguments[1]
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 || input == "" {
			log.Fatalln("Please Enter Correctly Characters")
		}
	}
	// slInput := strings.Split(input, "\n")

	// str := "hello \n\n world \n hi man"
	slInput := []string{}
	start := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\\' && input[i+1] == 'n' {
			if input[start:i] != "" {
				slInput = append(slInput, input[start:i])
			}
			slInput = append(slInput, "\\n")
			start = i + 2
		}
	}
	if input[start:] != "" {
		slInput = append(slInput, input[start:])
	}
	// slInput := strings.Split(str, "\n")
	return slInput
}
