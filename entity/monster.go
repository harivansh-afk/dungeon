package entity

import "time"

type Monster struct {
	Entity
	Kind string
}

func NewMonster(kind string) *Monster {
	m := &Monster{Kind: kind}
	switch kind {
	case "rat":
		m.Name="Rat"
		m.Symbol='R'
		m.Health=30
		m.MaxHealth=30
		m.Speed=150*time.Millisecond
	case "goblin":
		m.Name="Goblin"
		m.Symbol='G'
		m.Health=60
		m.MaxHealth=60
		m.Speed=200*time.Millisecond
	case "dragon":
		m.Name="Dragon"
		m.Symbol='D'
		m.Health=100
		m.MaxHealth=100
		m.Speed=400*time.Millisecond
	}
	return m
}
