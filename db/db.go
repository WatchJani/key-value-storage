package db

import "fmt"

type DB struct {
	table map[string]*Table
}

func New() *DB {
	return &DB{
		table: make(map[string]*Table),
	}
}

func (db *DB) SetTable(tableName string) {
	db.table[tableName] = CreateTable(tableName)
}

func (db *DB) GetTable(tableName string) *Table {
	return db.table[tableName]
}

func (db *DB) CreateTable(tableName string) (*Table, error) {
	if _, exist := db.table[tableName]; exist {
		return nil, fmt.Errorf("Table '%s' already exists", tableName)
	}

	db.SetTable(tableName)
	return db.GetTable(tableName), nil
}
