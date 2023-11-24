package main

import (
	"fmt"
	"log"
	"root/db"
)

func main() {
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

	user.Insert([]byte(`{
        "FirstName", "Janko",
        "LastName", "Kondić",
        "phone": "+386 311 063"
    }`), "janis:D12383_asd")

	user.Insert([]byte(`{
        "FirstName", "Marko",
        "LastName", "Markovic",
        "phone": "+386 311 063"
    }`), "janis:D12383_asdrg")

	data, err := user.GetData("janis:D12383_asdrg")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
