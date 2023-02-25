package svc

import (
	c "cs361FinalProject/constants"
)

func RecipeService(mainIngredient string) {
	switch mainIngredient {
	case c.Chicken:
		ChickenRecipes()
	case c.Beef:
		BeefRecipes()
	case c.Seafood:
		SeafoodRecipes()
	case c.Vegetable:
		VegetableRecipes()
	case c.Dessert:
		DessertRecipes()
	}
}
