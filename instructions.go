package main

import (
	"fmt"
	"os"
)

var UserInput string

func AppTitleIntro() {
	fmt.Println()
	fmt.Println()
	fmt.Println("*****  *****  *****  *****  *****  *****    *****  *****  *    *")
	fmt.Println("*   *  *      *        *    *   *  *        *   *  *   *  *    *")
	fmt.Println("*   *  *      *        *    *   *  *        *  *   *   *   *  * ")
	fmt.Println("*****  ****   *        *    *****  ****     ***    *   *    **  ")
	fmt.Println("**     *      *        *    *      *        *  *   *   *    **  ")
	fmt.Println("* *    *      *        *    *      *        *   *  *   *   *  * ")
	fmt.Println("*  *   *      *        *    *      *        *   *  *   *   *  * ")
	fmt.Println("*   *  *****  *****  *****  *      *****    *****  *****  *    *")
	fmt.Println()
	fmt.Println()
	fmt.Println("                 WELCOME TO YOUR RECIPE BOX!")
	fmt.Println()
	fmt.Println("                    * Pick an ingredient")
	fmt.Println("                    * Choose a recipe")
	fmt.Println("                    * Make your meal!")
}

func PickAnIngredient() {
	fmt.Println()
	fmt.Println("***** CHOOSE YOUR MAIN INGREDIENT *****")
	fmt.Println()
	fmt.Println("1. Chicken")
	fmt.Println("2. Beef")
	fmt.Println("3. Seafood")
	fmt.Println("4. Vegetable")
	fmt.Println("5. Dessert")
	fmt.Println()
	fmt.Println("6. EXIT RECIPE BOX")
	GetUserInput()

	//validInput := true
	//	for validInput (
	switch UserInput {
	case "1":
		ChooseYourRecipe("Chicken")
	case "2":
		ChooseYourRecipe("Beef")
	case "3":
		ChooseYourRecipe("Seafood")
	case "4":
		ChooseYourRecipe("Vegetable")
	case "5":
		ChooseYourRecipe("Dessert")
	case "6":
		ExitProgram()
		//validInput = false
	default:
		fmt.Println("Incorrect input, please try again")
		//validInput = true
	}
	// )
}

func ChooseYourRecipe(ingredient string) {
	fmt.Println()
	fmt.Printf("***** %s RECIPES *****", ingredient)
	fmt.Println()
}

func NAVChooseYourRecipe() {
	fmt.Println()
	fmt.Println("1. PICK DIFFERENT MAIN INGREDIENT")
	fmt.Println("2. EXIT RECIPE BOX")
}

func Recipe(recipeName string) {
	fmt.Println()
	fmt.Printf("***** RECIPE: %s *****", recipeName)
	fmt.Println()
}

func NAVRecipe() {
	fmt.Println()
	fmt.Println("1. PRINT RECIPE")
	fmt.Println("2. EMAIL RECIPE")
	fmt.Println("3. DOWNLOAD RECIPE")
	fmt.Println("4. CONVERT RECIPE TO METRIC")
	fmt.Println("5. CHOOSE A DIFFERENT RECIPE")
	fmt.Println("6. PICK DIFFERENT MAIN INGREDIENT")
	fmt.Println("7. EXIT RECIPE BOX")
}

func GetUserInput() {
	fmt.Println("(Type number and press enter to make your selection.)")
	fmt.Scan(&UserInput)
}

func ExitProgram() {
	fmt.Println()
	fmt.Println("*** EXITING RECIPE BOX ***")
	fmt.Println()
	os.Exit(0)
}
