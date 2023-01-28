package main

func RecipeService(mainIngredient string) {
	switch mainIngredient {
	case chicken:
		ChickenRecipes()
	case beef:
		BeefRecipes()
	case seafood:
		SeafoodRecipes()
	case vegetable:
		VegetableRecipes()
	case dessert:
		DessertRecipes()
	}
}
