package main

import (
	"log"
	"root/db"
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	stormDB := db.New()

	user, err := stormDB.CreateTable("user")
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < b.N; i++ {
		user.Insert([]byte(`{
        "FirstName", "Janko",
        "LastName", "Kondić",
        "phone": "+386 311 063"
    }`), "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5")
	}
}

func BenchmarkRead(b *testing.B) {
	stormDB := db.New()

	user, err := stormDB.CreateTable("user")
	if err != nil {
		log.Println(err)
		return
	}

	user.Insert([]byte(`{
        "FirstName", "Janko",
        "LastName", "Kondić",
        "phone": "+386 311 063"
    }`), "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5")

	for i := 0; i < b.N; i++ {
		user.GetData("a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5")
	}
}
