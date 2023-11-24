package db

import (
	"os"
	"path/filepath"
)

func CreateDirIfNotExist(name string) error {
	var err error

	if _, err = os.Stat(name); os.IsNotExist(err) {
		err = os.Mkdir(name, os.ModePerm)
	}

	return err
}

func CreateFileIfNotExist(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
}

func CreateDir(name string) error {
	return os.Mkdir(name, os.ModePerm)
}

func CreateFile(name string) (*os.File, error) {
	return os.Create(name)
}

func WriteData(data []byte, name, tableName string) {
	file, err := CreateFileIfNotExist(filepath.Join(LOG_PATH, tableName, name+".bin"))

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}

func ReadFromFile(store Store, tableName string) ([]byte, error) {
	file, err := CreateFileIfNotExist(filepath.Join(LOG_PATH, tableName, store.name+".bin"))
	defer file.Close()

	// Postavi se na početak određenog raspona
	_, err = file.Seek(int64(store.position), 0)
	if err != nil {
		return nil, err
	}

	// Čitanje podataka iz određenog raspona
	data := make([]byte, store.length)
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
