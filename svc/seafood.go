package svc

import (
	c "cs361FinalProject/constants"
	"fmt"
)

const (
	// Seafood Recipe Options
	crispySkinSalmon = "Crispy Skin Salmon"
	teriyakiSalmon   = "Teriyaki Salmon"
	airFryerShrimp   = "Air Fryer Shrimp"
	spicySalmon      = "Spicy Salmon"
	salmonBites      = "Salmon Bites"

	// Seafood Recipe URLs
	crispySkinSalmonURL = "https://minimalistbaker.com/crispy-skin-salmon-perfect-every-time/"
	teriyakiSalmonURL   = "https://www.loveandotherspices.com/air-fryer-teriyaki-salmon/"
	airFryerShrimpURL   = "https://www.loveandotherspices.com/air-fryer-shrimp/"
	spicySalmonURL      = "https://www.loveandotherspices.com/air-fryer-salmon/"
	salmonBitesURL      = "https://www.loveandotherspices.com/air-fryer-salmon-bites/"
)

func SeafoodRecipes() {
	fmt.Printf("1. %s\n", crispySkinSalmon)
	fmt.Printf("2. %s\n", teriyakiSalmon)
	fmt.Printf("3. %s\n", airFryerShrimp)
	fmt.Printf("4. %s\n", spicySalmon)
	fmt.Printf("5. %s\n", salmonBites)
	NAVChooseYourRecipe()
	GetUserInput()

	validInput := true
	for validInput {
		switch UserInput {
		case "1":
			ScrapeRecipe(crispySkinSalmonURL, crispySkinSalmon, c.Seafood)
		case "2":
			ScrapeRecipe(teriyakiSalmonURL, teriyakiSalmon, c.Seafood)
		case "3":
			ScrapeRecipe(airFryerShrimpURL, airFryerShrimp, c.Seafood)
		case "4":
			ScrapeRecipe(spicySalmonURL, spicySalmon, c.Seafood)
		case "5":
			ScrapeRecipe(salmonBitesURL, salmonBites, c.Seafood)
		case "6":
			PickAnIngredient()
		case "7":
			About()
			SeafoodRecipes()
		case "8":
			ExitProgram()
			validInput = false
		default:
			fmt.Println("Incorrect input, please try again")
			validInput = true
		}
	}
}
