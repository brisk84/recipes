package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func main() {
	uri := "postgresql://localhost/recipes?user=postgres&password=sqlRec1pe58&sslmode=disable"
	db, _ := sql.Open("postgres", uri)
	q01 := `select id, steps from recipes`
	rows, _ := db.Query(q01)
	for rows.Next() {
		var id string
		var steps []string
		err := rows.Scan(&id, pq.Array(&steps))
		if err != nil {
			panic(err)
		}
		fmt.Println(id, steps)
		for _, v := range steps {
			time := rand.Intn(59) + 1
			q02 := `insert into steps (title, time, recipe_id) values ($1, $2, $3)`
			_, err := db.Exec(q02, v, time, id)
			if err != nil {
				panic(err)
			}
		}
	}
}
