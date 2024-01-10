package main

import (
	"fmt"
	"time"
)

func main() {

	answer := 42
	fmt.Println(&answer)
	fmt.Println(&answer)

	address := &answer
	fmt.Println(*address)

	fmt.Printf("address это %T\n", address)

	canada := "Canada"

	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &canada
	fmt.Println(*home)

	var administrator *string

	scoles := "Chistopher J. Scolese"
	administrator = &scoles
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	bolden = "Charles Frank Bolder Jr."
	fmt.Println(*administrator)

	*administrator = "Maj. Gen. Charles Frank Bolder Jr."
	fmt.Println(bolden)

	major := administrator
	*major = "Major General Charles Frank Bolder Jr."
	fmt.Println(bolden)

	fmt.Println(administrator == major)

	lightfood := "Rebort M. Lightfoot Jr."
	administrator = &lightfood
	fmt.Println(administrator == major)

	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	fmt.Println(bolden)

	charles = "Charles Bolden"
	fmt.Println(charles == bolden)
	fmt.Println(&charles == &bolden)

	type person struct {
		name, superpower string
		age              int
	}

	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)

	superpowers := &[3]string{"flight", "invisibility", "super strenght"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])

	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)

	fmt.Printf("%+v\n", rebecca)

	terry := &person{
		name: "Terry",
		age:  15,
	}

	terry.birthday()
	fmt.Printf("%+v\n", terry)

	nathan := person{
		name: "Nathan",
		age:  17,
	}

	nathan.birthday()
	fmt.Printf("%+v\n", nathan)

	const layout = "Mon, Jan 2, 2006"
	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)
	fmt.Println(day.Format(layout))
	fmt.Println(tomorrow.Format(layout))

	player := character{name: "Matthias"}
	levelUp(&player.stats)

	fmt.Printf("%+v\n", player.stats)
	var board [8][8]rune
	reset(&board)

	fmt.Printf("%c", board[0][0])
}

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

type person struct {
	name, superpower string
	age              int
}

func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

type character struct {
	name  string
	stats stats
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}
