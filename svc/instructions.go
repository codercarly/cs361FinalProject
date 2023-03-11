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
	fmt.Printf("%s %s\n", c.One, c.Chicken)
	fmt.Printf("%s %s\n", c.Two, c.Beef)
	fmt.Printf("%s %s\n", c.Three, c.Seafood)
	fmt.Printf("%s %s\n", c.Four, c.Vegetable)
	fmt.Printf("%s %s\n", c.Five, c.Dessert)
	fmt.Println()
	fmt.Printf(color.Green +"%s %s\n", c.Six, c.About)
	fmt.Printf("%s %s\n" +color.Reset, c.Seven, c.Exit)

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
			fmt.Printf("%s", c.Incorrect)
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
	fmt.Printf(color.Green +"%s %s\n", c.Six, c.PickDifferentIng)
	fmt.Printf("%s %s\n", c.Seven, c.About)
	fmt.Printf("%s %s\n" + color.Reset, c.Eight, c.Exit)
}

func Recipe(recipeName string) {
	fmt.Println(color.Yellow)
	fmt.Printf("***** RECIPE: %s *****", recipeName)
	fmt.Println(color.Reset)
}

func NAVRecipe(mainIngredient string, recipeData RecipeData) {
	fmt.Println()
	fmt.Printf(color.Green +"%s %s\n", c.One, c.DownloadRecipe)
	fmt.Printf("%s %s\n", c.Two, c.ChooseDifferentRec)
	fmt.Printf("%s %s\n", c.Three, c.PickDifferentIng)
	fmt.Printf("%s %s\n", c.Four, c.About)
	fmt.Printf("%s %s\n" +color.Reset, c.Five, c.Exit)

	validInput := true
	for validInput {
		GetUserInput()
		switch UserInput {
		case "1":
			fmt.Println("DOWNLOAD")
			fileName, err := SendFileToText(recipeData)
			if err != nil {
				log.Fatal(err)
			}
			m.CallMicroserviceClient(fileName)
		case "2":
			ChooseYourRecipe(mainIngredient)
			RecipeService(mainIngredient)
		case "3":
			PickAnIngredient()
		case "4":
			About()
			NAVRecipe(mainIngredient, recipeData)
		case "5":
			ExitProgram()
			validInput = false
		default:
			fmt.Printf("%s\n", c.Incorrect)
			validInput = true
		}
	}
}

func GetUserInput() {
	fmt.Printf("%s\n", c.UserInputReq)
	fmt.Scan(&UserInput)
}

func ExitProgram() {
	fmt.Println(color.Red)
	fmt.Printf("%s\n", c.ExitingRecipeBox)
	fmt.Println(color.Reset)
	os.Exit(0)
}
