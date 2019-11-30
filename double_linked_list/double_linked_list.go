package main

import (
	"fmt"
	"time"
)

//Song ...
type Song struct {
	name     string
	previous *Song
	next     *Song
}

//NewSong ..
func NewSong(name string) *Song {
	return &Song{name: name}
}

//AddNextSong ...
func (s *Song) AddNextSong(next *Song) *Song {
	s.next = next
	next.previous = s
	return next
}

func getNextSong(s *Song) *Song {
	if s.next != nil {
		return s.next
	}
	return s
}

func main() {

	firstSong := NewSong("waka-waka")

	firstSong.
		AddNextSong(NewSong("bicicleta")).
		AddNextSong(NewSong("estoy aqui")).
		AddNextSong(NewSong("garota de Ipanema"))

	currentSong := firstSong
	previousSong := currentSong

	for currentSong != nil {
		fmt.Println("Playing: ")
		fmt.Println(currentSong.name)
		fmt.Println("*****")

		time.Sleep(1 * time.Second)
		previousSong = currentSong
		currentSong = currentSong.next
	}

	fmt.Println("||||||||||||||||||||||")
	fmt.Println("GOING BACK")
	fmt.Println("||||||||||||||||||||||")

	currentSong = previousSong

	for currentSong != nil {
		fmt.Println("Playing: ")
		fmt.Println(currentSong.name)
		fmt.Println("*****")

		time.Sleep(1 * time.Second)
		currentSong = currentSong.previous
	}
}
