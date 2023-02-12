package main

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
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
	author, _ := recipe.Author()
	cookTime, _ := recipe.CookTime()
	cuisine, _ := recipe.Cuisine()
	description, _ := recipe.Description()
	yields, _ := recipe.Yields()

	formattedCookTime := fmtDuration(cookTime)

	// Display basic recipe information
	println(color.Cyan + "Author: " + color.Reset + author)
	////println()
	println(color.Cyan + "Description: " + color.Reset + description)
	//println()
	println(formattedCookTime)
	//println()
	print(color.Cyan + "Cusine: " + color.Reset)
	for _, x := range cuisine {
		print(x)
	}
	println()
	println(color.Cyan + "Yields: " + color.Reset + yields)
	println()

	// Print Ingredients
	println(color.Yellow + "INGREDIENTS:" + color.Reset)
	ingredients, ok := recipe.Ingredients()
	if !ok {
		println("Unable to get recipe ingredients")
	}
	for _, i := range ingredients {
		println(i)
	}
	println()

	// Print Instructions
	println(color.Yellow + "INSTRUCTIONS:" + color.Reset)
	instructions, ok := recipe.Instructions()
	if !ok {
		println("Unable to get recipe instructions")
	}
	for _, i := range instructions {
		println(i)
	}
	NAVRecipe(mainIngredient)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	hours := fmt.Sprintf("%02d hours, ", h)
	minutes := fmt.Sprintf("%02d minutes", m)
	cookTime := fmt.Sprint(color.Cyan + "Cooktime: " + color.Reset)

	time := fmt.Sprint(cookTime + hours + minutes)

	return time
}
