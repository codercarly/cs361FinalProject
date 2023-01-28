package main

import "fmt"

const (
	// Beef Recipe Options
	garlicSteakBites       = "Garlic Steak Bites"
	cheeseStuffedMeatballs = "Cheese Stuffed Meatballs"
	beefNachos             = "Beef Nachos"
	roastBeef              = "Roast Beef"
	stuffedPeppers         = "Stuffed Peppers with Rice & Ground Beef"

	// Beef Recipe URLs
	garlicSteakBitesURL       = "https://www.loveandotherspices.com/air-fryer-garlic-steak-bites/"
	cheeseStuffedMeatballsURL = "https://www.loveandotherspices.com/air-fryer-cheese-stuffed-meatballs/"
	beefNachosURL             = "https://www.loveandotherspices.com/air-fryer-nachos/"
	roastBeefURL              = "https://www.loveandotherspices.com/easy-air-fryer-roast-beef/"
	stuffedPeppersURL         = "https://www.loveandotherspices.com/stuffed-peppers-with-rice-ground-beef/"
)

func BeefRecipes() {
	fmt.Printf("1. %s\n", garlicSteakBites)
	fmt.Printf("2. %s\n", cheeseStuffedMeatballs)
	fmt.Printf("3. %s\n", beefNachos)
	fmt.Printf("4. %s\n", roastBeef)
	fmt.Printf("5. %s\n", stuffedPeppers)
	NAVChooseYourRecipe()
	println()
	GetUserInput()

	validInput := true
	for validInput {
		switch UserInput {
		case "1":
			ScrapeRecipe(garlicSteakBitesURL, garlicSteakBites, beef)
		case "2":
			ScrapeRecipe(cheeseStuffedMeatballsURL, cheeseStuffedMeatballs, beef)
		case "3":
			ScrapeRecipe(beefNachosURL, beefNachos, beef)
		case "4":
			ScrapeRecipe(roastBeefURL, roastBeef, beef)
		case "5":
			ScrapeRecipe(stuffedPeppersURL, stuffedPeppers, beef)
		case "6":
			PickAnIngredient()
		case "7":
			About()
			BeefRecipes()
		case "8":
			ExitProgram()
			validInput = false
		default:
			fmt.Println("Incorrect input, please try again")
			validInput = true
		}
	}
}
