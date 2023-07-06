package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"recipes/domain"
	"strings"
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

type Recipe2 struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients []string `json:"recipeIngredient"`
	Steps       string   `json:"recipeInstructions"`
}

func main() {
	buf, _ := os.ReadFile("./cmd/load/recipes.json")
	var rs []Recipe
	err := json.Unmarshal(buf, &rs)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		var rec domain.Recipe
		rec.Title = rs[i].Name
		rec.Description = rs[i].Description
		rec.Ingredients = rs[i].Ingredients
		for _, v := range rs[i].Steps {
			curTime := rand.Intn(59) + 1
			rec.Steps = append(rec.Steps, domain.Step{
				Title: v.Text,
				Time:  curTime,
			})
			rec.TotalTime += curTime
		}
		if rec.Steps == nil {
			continue
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

	var rs2 []Recipe2
	err = json.Unmarshal(buf, &rs2)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		var rec domain.Recipe
		rec.Title = rs2[i].Name
		rec.Description = rs2[i].Description
		rec.Ingredients = rs2[i].Ingredients
		if rs2[i].Steps == "" {
			continue
		}
		steps := strings.Split(rs2[i].Steps, "\n")
		for _, v := range steps {
			curTime := rand.Intn(59) + 1
			rec.Steps = append(rec.Steps, domain.Step{
				Title: v,
				Time:  curTime,
			})
			rec.TotalTime += curTime
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
