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

	//chickenRecipes := [5]string{mushroomChicken, lemonHerbChickenThighs, whiteBeanChickenChili, teriyakiChickenSkewers, marinatedGrilledChicken}
	fmt.Printf("1. %s", mushroomChicken)
	fmt.Printf("2. %s", lemonHerbChickenThighs)
	fmt.Printf("3. %s", whiteBeanChickenChili)
	fmt.Printf("4. %s", teriyakiChickenSkewers)
	fmt.Printf("5. %s", marinatedGrilledChicken)
	println()
}
