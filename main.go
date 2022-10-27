package main

import (
	"fmt"
	"gorm_playground/db"

	"gorm.io/gorm"
)

type Party struct {
	ID   uint64
	Name string
	Dogs []Dog
}

type Dog struct {
	ID      uint64
	Breed   string
	Age     uint64
	PartyID uint64
}

func (Party) TableName() string {
	return "Party"
}

func (Dog) TableName() string {
	return "Dog"
}

func main() {
	db.MustInit()
	p := &Party{
		Name: "party 1",
		Dogs: []Dog{
			{
				Breed: "chihuahua",
				Age:   3,
			},
		},
	}

	tx := db.GetDB()
	_ = p.Create(tx) // a Party object and the associated dog objects will be created
	_ = p.Delete(tx) // the Println from Dog's BeforeDelete function will get triggered, but with an empty Dog object. value printed out is: &{0  0 0}
}

func (p *Party) Create(tx *gorm.DB) error {
	return tx.Create(p).Error
}

func (p *Party) Delete(tx *gorm.DB) error {
	return tx.Select("Dogs").Delete(p).Error
}

func (d *Dog) BeforeDelete(tx *gorm.DB) error {
	fmt.Println(d)
	return nil
}
