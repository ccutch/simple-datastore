package simpledb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type datadump struct {
	Tables map[string]Schema `json:"tables"`
	Data   map[string][]Node `json:"data"`
}

type Datastore struct {
	File   string
	tables []*Table
}

func Open(file string) *Datastore {
	db := &Datastore{
		File:   file,
		tables: []*Table{},
	}
	r, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return db
		} else {
			panic(err)
		}
	}
	var data datadump
	err = json.Unmarshal(r, &data)
	if err != nil {
		return db
	}
	for name, fields := range data.Tables {
		t := &Table{
			store:  db,
			name:   name,
			fields: fields,
		}
		db.tables = append(db.tables, t)
	}
	for name, data := range data.Data {
		t := db.GetTable(name)
		t.data = data
	}
	return db
}

func (db *Datastore) GetTable(name string) *Table {
	for _, t := range db.tables {
		if t.name == name {
			return t
		}
	}
	return nil
}

// Write data from tables to database file
func (db *Datastore) Dump() error {
	data := datadump{
		Tables: map[string]Schema{},
		Data:   map[string][]Node{},
	}
	for _, t := range db.tables {
		data.Tables[t.name] = t.fields
		data.Data[t.name] = t.data
	}
	c, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(db.File, c, 0644)
	return err
}

func (db *Datastore) MustDump() {
	err := db.Dump()
	if err != nil {
		panic(err)
	}
}

// Create new table in datastore
func (db *Datastore) DefineTable(name string, schema Schema) (*Table, error) {
	t := db.GetTable(name)
	if t != nil {
		return nil, errors.New("A table with that name already exists")
	}
	t = &Table{
		store:  db,
		name:   name,
		fields: schema,
		data:   []Node{},
	}
	db.tables = append(db.tables, t)
	return t, nil
}

// Get value of all tables in database
func (db *Datastore) Tables() []Table {
	tables := []Table{}
	for _, t := range db.tables {
		tables = append(tables, *t)
	}
	return tables
}
