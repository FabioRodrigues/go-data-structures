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

	strangerSong := &Song{"Garota de ipanema", nil}
	actualSong := *songTwo

	songTwo.next = strangerSong
	strangerSong.next = actualSong.next

	s := songOne
	i := 0
	for i < 4 {
		fmt.Println(s)
		s = s.next
		i++
	}
}
