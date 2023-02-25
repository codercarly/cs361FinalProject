package svc

import (
	"fmt"

	"github.com/TwiN/go-color"
)

func About() {
	fmt.Println(color.Green + "----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------" + color.Reset)
	println(color.Cyan + "ABOUT THE RECIPE BOX (v1.0.1)" + color.Reset)
	println()
	println("The Recipe Box is a quick way to find, print, and email recipes.")
	println()
	println(color.Yellow + "This application needs an active internet connection." + color.Reset)
	println()
	println(color.Cyan + "HOW TO:" + color.Reset)
	println("	* Choose a main ingredient from a list of ingredients.")
	println("	* Choose a recipe from a list of recipes.")
	println("	* View recipe and instructions.")
	println("	* Print, email, and download recipe.")
	println()
	println(color.Cyan + "NAVIGATION:" + color.Reset)
	println("	* Select an option by typing a number + enter.")
	println("	* Each step will provide additional navigation options.")
	fmt.Println(color.Green + "----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------" + color.Reset)
}
