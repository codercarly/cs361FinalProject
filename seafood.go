package main

import "fmt"

const (
	// Seafood Recipe Options
	crispySkinSalmon   = "Crispy Skin Salmon"
	teriyakiSalmon     = "Teriyaki Salmon"
	airFryerShrimp     = "Air Fryer Shrimp"
	prawnsCoconutCurry = "Prawns in Mild Coconut Curry"
	salmonBites        = "Salmon Bites"

	// Seafood Recipe URLs
	crispySkinSalmonURL   = "https://minimalistbaker.com/crispy-skin-salmon-perfect-every-time/"
	teriyakiSalmonURL     = "https://www.loveandotherspices.com/air-fryer-teriyaki-salmon/"
	airFryerShrimpURL     = "https://www.loveandotherspices.com/air-fryer-shrimp/"
	prawnsCoconutCurryURL = "https://www.loveandotherspices.com/prawns-in-mild-coconut-curry/"
	salmonBitesURL        = "https://www.loveandotherspices.com/air-fryer-salmon-bites/"
)

func SeafoodRecipes() {
	fmt.Printf("1. %s\n", crispySkinSalmon)
	fmt.Printf("2. %s\n", teriyakiSalmon)
	fmt.Printf("3. %s\n", airFryerShrimp)
	fmt.Printf("4. %s\n", prawnsCoconutCurry)
	fmt.Printf("5. %s\n", salmonBites)
	NAVChooseYourRecipe()
	println()
	GetUserInput()

	validInput := true
	for validInput {
		switch UserInput {
		case "1":
			ScrapeRecipe(crispySkinSalmonURL, crispySkinSalmon, seafood)
		case "2":
			ScrapeRecipe(teriyakiSalmonURL, teriyakiSalmon, seafood)
		case "3":
			ScrapeRecipe(airFryerShrimpURL, airFryerShrimp, seafood)
		case "4":
			ScrapeRecipe(prawnsCoconutCurryURL, prawnsCoconutCurry, seafood)
		case "5":
			ScrapeRecipe(salmonBitesURL, salmonBites, seafood)
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
