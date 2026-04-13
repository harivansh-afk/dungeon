package entity

import "time"

type Player struct {
	Entity
	Level int
	XP    int
}

func NewPlayer() *Player {
	return &Player{
		Entity: Entity{
			Name:      "Hero",
			Symbol:    '@',
			Health:    100,
			MaxHealth: 100,
			Speed:     200 * time.Millisecond,
		},
		Level: 1,
	}
}

func (p *Player) AddXP(amount int) {
	p.XP += amount
	if p.XP >= 100*p.Level {
		p.Level++
		p.MaxHealth += 20
		p.Health = p.MaxHealth
	}
}
