package main

import "fmt"

const (
	// Vegetable Recipe Options
	veganStuffedShells     = "Easy Vegan Stuffed Shells"
	vegetableSoup          = "Simple But Perfect Vegetable Soup"
	veganBolognese         = "Vegan Bolognses with Mushrooms and Red Lentils"
	eggplantParmesean      = "Eggplant Parmesean"
	chipotleBlackBeanChili = "Chipotle Black Bean Chili"

	// Vegetable
	veganStuffedShellsURL     = "https://minimalistbaker.com/easy-vegan-stuffed-shells/"
	vegetableSoupURL          = "https://minimalistbaker.com/simple-but-perfect-vegetable-soup/"
	veganBologneseURL         = "https://minimalistbaker.com/vegan-bolognese-with-mushrooms-red-lentils/"
	eggplantParmeseanURL      = "https://www.loveandotherspices.com/air-fryer-eggplant-parmesan/"
	chipotleBlackBeanChiliURL = "https://minimalistbaker.com/1-pot-chipotle-black-bean-chili/"
)

func VegetableRecipes() {
	fmt.Printf("1. %s\n", veganStuffedShells)
	fmt.Printf("2. %s\n", vegetableSoup)
	fmt.Printf("3. %s\n", veganBolognese)
	fmt.Printf("4. %s\n", eggplantParmesean)
	fmt.Printf("5. %s\n", chipotleBlackBeanChili)
	NAVChooseYourRecipe()
	println()
	GetUserInput()

	validInput := true
	for validInput {
		switch UserInput {
		case "1":
			ScrapeRecipe(veganStuffedShellsURL, veganStuffedShells, vegetable)
		case "2":
			ScrapeRecipe(vegetableSoupURL, vegetableSoup, vegetable)
		case "3":
			ScrapeRecipe(veganBologneseURL, veganBolognese, vegetable)
		case "4":
			ScrapeRecipe(eggplantParmeseanURL, eggplantParmesean, vegetable)
		case "5":
			ScrapeRecipe(chipotleBlackBeanChiliURL, chipotleBlackBeanChili, vegetable)
		case "6":
			PickAnIngredient()
		case "7":
			About()
			VegetableRecipes()
		case "8":
			ExitProgram()
			validInput = false
		default:
			fmt.Println("Incorrect input, please try again")
			validInput = true
		}
	}
}
