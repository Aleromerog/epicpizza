package epicpizza

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/kkyr/go-recipe/pkg/recipe"
)

func RandomPizzaRecipe() (map[string][]string, error) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	url := data[rand.Intn(len(data))]

	recipe, err := recipe.ScrapeURL(url)
	if err != nil {
		err := fmt.Sprintf("%e %s", err, url)
		return nil, errors.New(err)
	}

	ingredients, ok_ingredients := recipe.Ingredients()
	instructions, ok_instructions := recipe.Instructions()

	if !ok_ingredients || !ok_instructions {
		err := fmt.Sprintf("Couldn't fetch recipe from site :( %s", url)
		return nil, errors.New(err)
	}

	m := make(map[string][]string)
	m["ingredients"] = ingredients
	m["instructions"] = instructions

	return m, nil
}

var data = []string{
	"https://tasty.co/recipe/pizza-dough",
	"https://sallysbakingaddiction.com/homemade-pizza-crust-recipe/",
	"https://www.loveandlemons.com/homemade-pizza/",
}
