package svc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func SendFileToText(rd RecipeData) (string, error) {

	recipeName := rd.Title
	fileName := recipeName + ".txt"
	// Create filename based on recipe title.
	// os.Create will create or truncate the named file.
	// If file already exists, it is truncated, if the file does not exist, it is created.
	_, err := os.Create(filepath.Join("/Users/carly/cs361FinalProject/microservice", filepath.Base(fileName)))
	if err != nil {
		log.Fatal(err)
	}

	// Open file and create buffer to write to file line by line
	file, _ := os.OpenFile(filepath.Join("/Users/carly/cs361FinalProject/microservice", filepath.Base(fileName)), os.O_RDWR|os.O_TRUNC, 0644)
	defer file.Close()

	buffer := bufio.NewWriter(file)

	err = os.WriteFile(filepath.Join("/Users/carly/cs361FinalProject/microservice", filepath.Base(fileName)), []byte(rd.Title), 0)

	_, err = fmt.Fprintf(file, "RECIPE: %v\n\n", rd.Title)
	_, err = fmt.Fprintf(file, "CREATED BY: %v\n\n", rd.Author)
	_, err = fmt.Fprintf(file, "DESCRIPTION: %v\n\n", rd.Description)
	_, err = fmt.Fprintf(file, "CUISINE: %v\n\n", rd.Cuisine)
	_, err = fmt.Fprintf(file, "YIELDS: %v\n\n", rd.Yields)
	_, err = fmt.Fprintf(file, "COOK TIME: %s\n\n", rd.CookTime)

	_, err = fmt.Fprintf(file, "INGREDIENTS:\n")
	for _, i := range rd.Ingredients {
		_, err = fmt.Fprintf(file, "%v\n", i)
	}
	_, err = fmt.Fprintf(file, "\n")

	_, err = fmt.Fprintf(file, "INSTRUCTIONS:\n")
	for x, i := range rd.Instructions {
		_, err = fmt.Fprintf(file, "%v. %v\n", x+1, i)
	}
	_, err = fmt.Fprintf(file, "\n\n")
	_, err = fmt.Fprintf(file, "RECIPE FOUND IN YOUR RECIPE BOX!\n\n")
	_, err = fmt.Fprintf(file, "END\n")

	buffer.Flush()
	fmt.Println("Converting your recipe now...")
	return fileName, nil
}
