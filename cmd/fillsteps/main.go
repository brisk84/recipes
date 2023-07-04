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
		rows.Scan(&id, pq.Array(&steps))
		fmt.Println(id, steps)
		for _, v := range steps {
			time := rand.Intn(59) + 1
			q02 := `insert into steps (title, time, recipe_id) values ($1, $2, $3)`
			db.Exec(q02, v, time, id)
		}
	}
}
