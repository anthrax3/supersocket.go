package main

import (
	"testing"
    "fmt"
)

type IAnimal interface {
    GetName() string
    GetColor() string
	MakeSound()
}

type Animal struct {
	Name  string
    Color  string
	Legs  int8
}

func (animal Animal) MakeSound() {
    fmt.Print("...");
}

func (animal Animal) GetName() string {
    return animal.Name
}

func (animal Animal) GetColor() string {
    return animal.Color
}

type Dog struct {
	Animal
	Tall int8
}

func (dog Dog) MakeSound() {
    fmt.Print("Wang wang");
}

func TestComposition(t *testing.T) {
    var mike = Animal { Name: "Mike", Color: "Black", Legs: 2 }
   
    var henry = Dog { }
    henry.Name = "Henry"
    henry.Color = "White"
    henry.Legs = 4
    henry.Tall = 10
    
    var animals = []IAnimal { mike, henry }
  
    for _, a := range animals {
        fmt.Printf("Name: %s, Color: %s", a.GetName(), a.GetColor());
		fmt.Println()
        a.MakeSound()
        fmt.Println()
	}
}