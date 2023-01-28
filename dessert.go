package main

import "fmt"

const (
	// Dessert Recipe Options
	veganChocolateCookies  = "Gluten Free Chocolate Crinkle Cookies"
	cakeMixCookies         = "Yellow Cake Mix Cookies"
	lemonBlueBerryCookies  = "Lemon Blueberry Cookies"
	ketoPeantButterCookies = "Keto Peanut Butter Cookies"
	pineappleWithCinnamon  = "Air Fryer Pineapple with Cinnamon"

	veganChocolateCookiesURL  = "https://minimalistbaker.com/gluten-free-chocolate-crinkle-cookies-vegan/"
	cakeMixCookiesURL         = "https://www.loveandotherspices.com/yellow-cake-mix-cookies/"
	lemonBlueBerryCookiesURL  = "https://www.loveandotherspices.com/lemon-blueberry-cookies/"
	ketoPeantButterCookiesURL = "https://www.loveandotherspices.com/keto-peanut-butter-cookies/"
	pineappleWithCinnamonURL  = "https://www.loveandotherspices.com/air-fryer-pineapple/"
)

func DessertRecipes() {
	fmt.Printf("1. %s\n", veganChocolateCookies)
	fmt.Printf("2. %s\n", cakeMixCookies)
	fmt.Printf("3. %s\n", lemonBlueBerryCookies)
	fmt.Printf("4. %s\n", ketoPeantButterCookies)
	fmt.Printf("5. %s\n", pineappleWithCinnamon)
	NAVChooseYourRecipe()
	println()
	GetUserInput()

	validInput := true
	for validInput {
		switch UserInput {
		case "1":
			ScrapeRecipe(veganChocolateCookiesURL, veganChocolateCookies, dessert)
		case "2":
			ScrapeRecipe(cakeMixCookiesURL, cakeMixCookies, dessert)
		case "3":
			ScrapeRecipe(lemonBlueBerryCookiesURL, lemonBlueBerryCookies, dessert)
		case "4":
			ScrapeRecipe(ketoPeantButterCookiesURL, ketoPeantButterCookies, dessert)
		case "5":
			ScrapeRecipe(pineappleWithCinnamonURL, pineappleWithCinnamon, dessert)
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
