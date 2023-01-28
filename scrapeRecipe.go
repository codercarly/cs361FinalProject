package main

import (
	"fmt"

	"github.com/kkyr/go-recipe/pkg/recipe"
)

func ScrapeRecipe(url string, recipeName string, mainIngredient string) {
	// Print out Recipe Name
	Recipe(recipeName)

	// Scrape Recipe data from website
	recipe, err := recipe.ScrapeURL(url)
	if err != nil {
		fmt.Printf("Unable to get recipe data")
	}

	// Display basic recipe information
	println(recipe.Author())
	println()
	println(recipe.CookTime())
	println()
	println(recipe.Cuisine())
	println()
	println(recipe.Description())
	println()
	println(recipe.Yields())
	println()

	// Print Ingredients
	println("INGREDIENTS:")
	ingredients, ok := recipe.Ingredients()
	if !ok {
		println("Unable to get recipe ingredients")
	}
	for _, i := range ingredients {
		println(i)
	}
	println()

	// Print Instructions
	println("INSTRUCTIONS:")
	instructions, ok := recipe.Instructions()
	if !ok {
		println("Unable to get recipe instructions")
	}
	println()
	for _, i := range instructions {
		println(i)
	}
	println()
	NAVRecipe(mainIngredient)
}
