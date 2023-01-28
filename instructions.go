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
	fmt.Println("1. CHICKEN")
	fmt.Println("2. BEEF")
	fmt.Println("3. SEAFOOD")
	fmt.Println("4. VEGETABLE")
	fmt.Println("5. DESSERT")
	fmt.Println()
	fmt.Println("6. EXIT RECIPE BOX")

	validInput := true
	for validInput {
		GetUserInput()
		switch UserInput {
		case "1":
			ChooseYourRecipe(chicken)
			RecipeService(chicken)
		case "2":
			ChooseYourRecipe(beef)
			RecipeService(beef)
		case "3":
			ChooseYourRecipe(seafood)
			RecipeService(seafood)
		case "4":
			ChooseYourRecipe(vegetable)
			RecipeService(vegetable)
		case "5":
			ChooseYourRecipe(dessert)
			RecipeService(dessert)
		case "6":
			ExitProgram()
			validInput = false
		default:
			fmt.Println("Incorrect input, please try again")
			validInput = true
		}
	}
}

func ChooseYourRecipe(ingredient string) {
	fmt.Println()
	fmt.Printf("***** %s RECIPES *****", ingredient)
	fmt.Println()
}

func NAVChooseYourRecipe() {
	fmt.Println()
	fmt.Println("6. PICK DIFFERENT MAIN INGREDIENT")
	fmt.Println("7. EXIT RECIPE BOX")
}

func Recipe(recipeName string) {
	fmt.Println()
	fmt.Printf("***** RECIPE: %s *****", recipeName)
	fmt.Println()
}

func NAVRecipe(mainIngredient string) {
	fmt.Println()
	fmt.Println("1. PRINT RECIPE")
	fmt.Println("2. EMAIL RECIPE")
	fmt.Println("3. DOWNLOAD RECIPE")
	fmt.Println("4. CONVERT RECIPE TO METRIC")
	fmt.Println("5. CHOOSE A DIFFERENT RECIPE")
	fmt.Println("6. PICK DIFFERENT MAIN INGREDIENT")
	fmt.Println("7. EXIT RECIPE BOX")

	validInput := true
	for validInput {
		GetUserInput()
		switch UserInput {
		case "1":
			fmt.Println("TODO: PRINT")
		case "2":
			fmt.Println("TODO: EMAIL")
		case "3":
			fmt.Println("TODO: DOWNLOAD")
		case "4":
			fmt.Println("TODO: CONVERT")
		case "5":
			ChooseYourRecipe(mainIngredient)
			RecipeService(mainIngredient)
		case "6":
			PickAnIngredient()
		case "7":
			ExitProgram()
			validInput = false
		default:
			fmt.Println("Incorrect input, please try again")
			validInput = true
		}
	}
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
