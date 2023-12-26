package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/rs/xid"
	"github.com/surrealdb/surrealdb.go"
	"time"
)

const branchingFactor = 2 // 100=>1M, 200=>8M, 300=>27M

type OrgUnit struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func main() {
	// Connect to SurrealDB
	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}

	// Select namespace and database
	if _, err = db.Use("default", "default"); err != nil {
		panic(err)
	}

	if _, err = db.Delete("org_unit"); err != nil {
		panic(err)
	}

	bulkCreate(db)
}

func bulkCreate(db *surrealdb.DB) {
	for i := 0; i < branchingFactor; i++ {
		time.Sleep(time.Millisecond * 500)
		id1 := create(db)

		for j := 0; j < branchingFactor; j++ {
			id2 := createWithParent(db, id1, 2)

			for k := 0; k < branchingFactor; k++ {
				_ = createWithParent(db, id2, 3)
			}
		}
	}
}

func create(db *surrealdb.DB) string {
	id := xid.New().String()

	_, err := db.Create("org_unit", OrgUnit{
		ID:    id,
		Name:  faker.Name(),
		Level: 1,
	})
	if err != nil {
		panic(err)
	}

	return id
}

func createWithParent(db *surrealdb.DB, parentID string, level int) string {
	id := xid.New().String()

	query := fmt.Sprintf(
		`
BEGIN TRANSACTION;
CREATE type::thing("org_unit", $id) CONTENT {
  name: $name,
  level: $level,
};
RELATE org_unit:%s->belongs_to->org_unit:%s;
COMMIT TRANSACTION;
`, id, parentID,
	)

	_, err := db.Query(query, map[string]any{
		"id":       id,
		"name":     faker.Name(),
		"parentID": parentID,
		"level":    level,
	})
	if err != nil {
		panic(err)
	}

	return id
}
