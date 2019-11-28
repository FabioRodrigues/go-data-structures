package main

import "fmt"

type Song struct {
	name string
	next *Song
}

func getNextSong(s *Song) *Song {
	if s.next != nil {
		return s.next
	}
	return s
}

func main() {

	songThree := &Song{"waka-waka", nil}
	songTwo := &Song{"bicicleta", songThree}
	songOne := &Song{"estoy aqui", songTwo}

	s := songOne

	strangerSong := &Song{"Garota de ipanema", nil}

	actualSong := *songTwo

	songTwo.next = strangerSong
	strangerSong.next = actualSong.next

	count := 0
	for s.name != "waka-waka" || s.next != nil {
		fmt.Println(s)
		s = s.next
		count++
	}
	fmt.Println("element found at index", count)
}
