package db

import (
	"fmt"
	"strings"

	. "restapi.com/hello/helpers"
	p "restapi.com/hello/person"
)

type DB struct {
	m      map[int]*p.Person
	id_ctr int
}

func New() *DB {
	return &DB{map[int]*p.Person{}, 0}
}

func (db *DB) Create(p *p.Person) int {
	db.id_ctr++
	db.m[db.id_ctr] = p
	return db.id_ctr
}

func (db *DB) Read(id int) (*p.Person, error) {
	if v, ok := db.m[id]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("id %d not found", id)
}

func (db *DB) Update(id int, p *p.Person) error {
	if _, ok := db.m[id]; ok {
		db.m[id] = p
		return nil
	}
	return fmt.Errorf("id %d not found", id)
}

func (db *DB) Delete(id int) error {
	if _, ok := db.m[id]; ok {
		delete(db.m, id)
		return nil
	}
	return fmt.Errorf("id %d not found", id)
}

func (db *DB) Search(sub string) []any {
	res := []any{}
	for id, p := range db.m {
		if include(p.FirstName, sub) || include(p.LastName, sub) {
			res = append(res, JMap{
				"id":     id,
				"person": p,
			})
		}
	}
	return res
}

func include(str string, sub string) bool {
	return strings.Contains(strings.ToUpper(str), strings.ToUpper(sub))
}
