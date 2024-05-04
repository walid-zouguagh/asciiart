package main

import (
	"log"
	"os"

	"asciiart/functions"
)

func main() {
	/***********************************/
	/**** Read And Chekin Arguments ****/
	/***********************************/
	arguments := os.Args
	if len(arguments) != 2 {
		log.Fatalln("Please Enter Correctly Arguments")
	}
	verifiedInput := functions.VerifiedInput(arguments)
	// VerifiedInput(arguments)

	/**********************************************/
	/******* Read Content the File Standard *******/
	/**********************************************/
	contentFileStandard := functions.ReadFileStandard("./sources/standard.txt")

	/**********************/
	/******* Output *******/
	/**********************/
	functions.Output(verifiedInput, contentFileStandard)

	// input := "hello \n\n world \n hi man"
	// slInput := strings.Split(input, "\n")
	// fmt.Println(slInput)
}
