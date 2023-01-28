package main

import "fmt"

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
	fmt.Printf("1. %s\n", mushroomChicken)
	fmt.Printf("2. %s\n", lemonHerbChickenThighs)
	fmt.Printf("3. %s\n", whiteBeanChickenChili)
	fmt.Printf("4. %s\n", teriyakiChickenSkewers)
	fmt.Printf("5. %s\n", marinatedGrilledChicken)
	NAVChooseYourRecipe()
	println()
	GetUserInput()

	validInput := true
	for validInput {
		switch UserInput {
		case "1":
			ScrapeRecipe(mushroomChickenURL, mushroomChicken, chicken)
		case "2":
			ScrapeRecipe(lemonHerbChickenThighsURL, lemonHerbChickenThighs, chicken)
		case "3":
			ScrapeRecipe(whiteBeanChickenChiliURL, whiteBeanChickenChili, chicken)
		case "4":
			ScrapeRecipe(teriyakiChickenSkewersURL, teriyakiChickenSkewers, chicken)
		case "5":
			ScrapeRecipe(marinatedGrilledChickenURL, marinatedGrilledChicken, chicken)
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
