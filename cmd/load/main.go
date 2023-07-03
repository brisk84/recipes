package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"recipes/domain"
)

type Step struct {
	Type string `json:"@type"`
	Text string `json:"text"`
}

type Recipe struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients []string `json:"recipeIngredient"`
	Steps       []Step   `json:"recipeInstructions"`
}

func main() {
	buf, _ := os.ReadFile("./cmd/load/recipes.json")
	var rs []Recipe
	json.Unmarshal(buf, &rs)

	for i := 0; i < 100; i++ {
		var rec domain.Recipe
		rec.Title = rs[i].Name
		rec.Description = rs[i].Description
		rec.Ingredients = rs[i].Ingredients
		for _, v := range rs[i].Steps {
			rec.Steps = append(rec.Steps, v.Text)
		}

		jd, err := json.Marshal(rec)
		if err != nil {
			log.Fatal(err)
		}

		_, err = http.Post("http://localhost:8000/api/recipe/c/create", "application/json",
			bytes.NewBuffer(jd))
		if err != nil {
			log.Fatal(err)
		}
	}
}
