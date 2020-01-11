package main

import "fmt"

type Ball struct {
	Radius int
}

func (b Ball) Bounce() {
	fmt.Println(b, b.Radius)
}

type Basketball struct {
	Radius int
	*Ball
	Weight int
}

type Action interface {
	Bounce()
}

type VolleyBall struct {
	Action
	Weight int
}

func main() {
	b := Ball{40}
	bb := Basketball{Radius: 60, Weight: 100, Ball: &b}

	bb.Bounce()
	fmt.Println(bb.Radius)

	// embed interface
	vb := VolleyBall{&Ball{50}, 100}
	vb.Bounce()
}
