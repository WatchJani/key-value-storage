package db

import (
	"path/filepath"
)

type Table struct {
	name    string
	segment Segment
	index   map[string]Store
}

func CreateTable(name string) *Table {
	segment := NewSegment(name, 0)
	ErrorChecker(CreateDirIfNotExist(filepath.Join(LOG_PATH, name))) //make folder if not exist

	return &Table{
		name:    name,
		index:   make(map[string]Store),
		segment: segment,
	}
}

func (t *Table) SetData(key string, store Store) {
	t.index[key] = store
}

func (t Table) GetKey(key string) Store {
	return t.index[key]
}

func (t *Table) GetData(key string) ([]byte, error) {
	return ReadFromFile(t.GetKey(key), t.name)
}

func (t *Table) Insert(data []byte, key string) {
	store := NewStore(len(data), &t.segment) //update data about position in our segment
	t.SetData(key, store)                    //save new key
	WriteData(data, t.segment.name, t.name)
}
