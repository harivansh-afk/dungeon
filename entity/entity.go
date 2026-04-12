package entity

import (
	"time"
)

type Entity struct {
	Position
	Name string
	Symbol rune
	Health int
	MaxHealth int
	Speed time.Duration
}

type Position struct {
	X, Y int
}

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}

func (e *Entity) TakeDamage(amount int) {
	e.Health -= amount
	if e.Health < 0 {
		e.Health = 0
	}
}

func (e *Entity) Heal(amount int) {
	e.Health += amount
	if e.Health > e.MaxHealth {
		e.Health = e.MaxHealth
	}
}

func (e *Entity) IsAlive() bool {
	return e.Health > 0
}
