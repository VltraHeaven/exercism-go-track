package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerPrep int) (prepTime int) {
    switch layerPrep {
    	case 0:
    	layerPrep = 2
    }
	return len(layers) * layerPrep
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {  
    for _, v := range layers {
        switch v {
        	case "noodles":
            noodles += 50
        	case "sauce":
        	sauce += 0.2
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngredients []string, myIngredients []string) {
    copy(myIngredients[len(myIngredients)-1:], friendIngredients[len(friendIngredients)-1:])
}


// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) (totalAmounts []float64) {
    for _, v := range amounts {
        totalAmounts = append(totalAmounts, v / 2 * float64(portions))
    }
	return
} 
