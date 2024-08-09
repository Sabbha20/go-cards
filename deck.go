package main

import "fmt"

// Created a custom type deck
type deck []string

func (d deck) print(){
	for i, card :=range d{
		fmt.Println(i, card)
	}
}