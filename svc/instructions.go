package svc

import (
	"fmt"
	"log"
	"os"

	c "cs361FinalProject/constants"
	m "cs361FinalProject/microservice"

	"github.com/TwiN/go-color"
)

var UserInput string

func AppTitleIntro() {
	fmt.Println(color.Green + "----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------" + color.Reset + color.Cyan)
	fmt.Println("*****  *****  *****  *****  *****  *****    *****  *****  *    *")
	fmt.Println("*   *  *      *        *    *   *  *        *   *  *   *  *    *")
	fmt.Println("*   *  *      *        *    *   *  *        *  *   *   *   *  * ")
	fmt.Println("*****  ****   *        *    *****  ****     ***    *   *    **  ")
	fmt.Println("**     *      *        *    *      *        *  *   *   *    **  ")
	fmt.Println("* *    *      *        *    *      *        *   *  *   *   *  * ")
	fmt.Println("*  *   *      *        *    *      *        *   *  *   *   *  * ")
	fmt.Println("*   *  *****  *****  *****  *      *****    *****  *****  *    *")
	fmt.Println(color.Green + "----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------" + color.Reset)
	fmt.Println(color.Cyan)
	fmt.Println("		WELCOME TO YOUR RECIPE BOX!")
	fmt.Println()
	fmt.Println("		* Pick an ingredient")
	fmt.Println("		* Choose a recipe")
	fmt.Println("		* Make your meal!" + color.Reset)
	fmt.Println()
}

func PickAnIngredient() {
	fmt.Println(color.Yellow + "***** CHOOSE YOUR MAIN INGREDIENT *****" + color.Reset)
	fmt.Println("1. CHICKEN")
	fmt.Println("2. BEEF")
	fmt.Println("3. SEAFOOD")
	fmt.Println("4. VEGETABLE")
	fmt.Println("5. DESSERT")
	fmt.Println()
	fmt.Println("6. ABOUT RECIPE BOX")
	fmt.Println("7. EXIT RECIPE BOX")

	validInput := true
	for validInput {
		GetUserInput()
		switch UserInput {
		case "1":
			ChooseYourRecipe(c.Chicken)
			RecipeService(c.Chicken)
		case "2":
			ChooseYourRecipe(c.Beef)
			RecipeService(c.Beef)
		case "3":
			ChooseYourRecipe(c.Seafood)
			RecipeService(c.Seafood)
		case "4":
			ChooseYourRecipe(c.Vegetable)
			RecipeService(c.Vegetable)
		case "5":
			ChooseYourRecipe(c.Dessert)
			RecipeService(c.Dessert)
		case "6":
			About()
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

func ChooseYourRecipe(ingredient string) {
	fmt.Println(color.Yellow)
	fmt.Printf("***** %s RECIPES *****", ingredient)
	fmt.Println(color.Reset)
}

func NAVChooseYourRecipe() {
	fmt.Println()
	fmt.Println("6. PICK DIFFERENT MAIN INGREDIENT")
	fmt.Println("7. ABOUT RECIPE BOX")
	fmt.Println("8. EXIT RECIPE BOX")
}

func Recipe(recipeName string) {
	fmt.Println(color.Yellow)
	fmt.Printf("***** RECIPE: %s *****", recipeName)
	fmt.Println(color.Reset)
}

func NAVRecipe(mainIngredient string, recipeData RecipeData) {
	fmt.Println()
	fmt.Println("1. PRINT RECIPE")
	fmt.Println("2. EMAIL RECIPE")
	fmt.Println("3. DOWNLOAD RECIPE")
	fmt.Println("4. CONVERT RECIPE TO METRIC")
	fmt.Println("5. CHOOSE A DIFFERENT RECIPE")
	fmt.Println("6. PICK DIFFERENT MAIN INGREDIENT")
	fmt.Println("7. ABOUT RECIPE BOX")
	fmt.Println("8. EXIT RECIPE BOX")

	validInput := true
	for validInput {
		GetUserInput()
		switch UserInput {
		case "1":
			fmt.Println("TODO: PRINT")
		case "2":
			fmt.Println("TODO: EMAIL")
		case "3":
			fmt.Println("DOWNLOAD")
			fileName, err := SendFileToText(recipeData)
			if err != nil {
				log.Fatal(err)
			}
			m.SendToPDF(fileName)
		case "4":
			fmt.Println("TODO: CONVERT")
		case "5":
			ChooseYourRecipe(mainIngredient)
			RecipeService(mainIngredient)
		case "6":
			PickAnIngredient()
		case "7":
			About()
			NAVRecipe(mainIngredient, recipeData)
		case "8":
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
	fmt.Println(color.Red)
	fmt.Println("*** EXITING RECIPE BOX ***")
	fmt.Println(color.Reset)
	os.Exit(0)
}