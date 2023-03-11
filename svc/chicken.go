package svc

import (
	c "cs361FinalProject/constants"
	"fmt"
)

const (
	// Chicken Recipe Options
	mushroomChicken         = "Creamy Samoan Inspired Mushroom Chicken"
	lemonHerbChickenThighs  = "Lemon Herb Roasted Chicken Thighs"
	whiteBeanChickenChili   = "Creamy White Bean Chicken Chili"
	teriyakiChickenSkewers  = "Grilled Teriyaki Chicken Skewers"
	marinatedGrilledChicken = "Easy Marinated Grilled Chicken"

	// Chicken Recipe URLs
	mushroomChickenURL         = "https://minimalistbaker.com/creamy-samoan-inspired-mushroom-chicken/"
	lemonHerbChickenThighsURL  = "https://minimalistbaker.com/lemon-herb-roasted-chicken-thighs/"
	whiteBeanChickenChiliURL   = "https://minimalistbaker.com/creamy-white-bean-chicken-chili-1-pot/"
	teriyakiChickenSkewersURL  = "https://minimalistbaker.com/grilled-teriyaki-chicken-skewers-quick-marinating/"
	marinatedGrilledChickenURL = "https://minimalistbaker.com/easy-marinated-grilled-chicken-30-minutes/"
)

func ChickenRecipes() {
	fmt.Printf("%s %s\n", c.One, mushroomChicken)
	fmt.Printf("%s %s\n", c.Two, lemonHerbChickenThighs)
	fmt.Printf("%s %s\n", c.Three, whiteBeanChickenChili)
	fmt.Printf("%s %s\n", c.Four, teriyakiChickenSkewers)
	fmt.Printf("%s %s\n", c.Five, marinatedGrilledChicken)
	NAVChooseYourRecipe()
	GetUserInput()

	validInput := true
	for validInput {
		switch UserInput {
		case "1":
			ScrapeRecipe(mushroomChickenURL, mushroomChicken, c.Chicken)
		case "2":
			ScrapeRecipe(lemonHerbChickenThighsURL, lemonHerbChickenThighs, c.Chicken)
		case "3":
			ScrapeRecipe(whiteBeanChickenChiliURL, whiteBeanChickenChili, c.Chicken)
		case "4":
			ScrapeRecipe(teriyakiChickenSkewersURL, teriyakiChickenSkewers, c.Chicken)
		case "5":
			ScrapeRecipe(marinatedGrilledChickenURL, marinatedGrilledChicken, c.Chicken)
		case "6":
			PickAnIngredient()
		case "7":
			About()
			ChickenRecipes()
		case "8":
			ExitProgram()
			validInput = false
		default:
			fmt.Println("Incorrect input, please try again")
			validInput = true
		}
	}
}
