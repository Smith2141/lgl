package main

import (
	"fmt"
	// "log"
	"reflect"
)

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу,
// то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []any) {
	for _, obj := range objects {
		// fmt.Println("before: ", obj)
		CastTo(spell, obj)
		// fmt.Println("after: ", obj)
	}
}

// реализуйте эту функцию.
func CastTo(spell Spell, object any) {
	if val, ok := object.(CastReceiver); ok {
		val.ReceiveSpell(spell)
	} else {
		val := reflect.ValueOf(object)
		if val.Kind() != reflect.Pointer {
			return
		}

		val = val.Elem()

		for fieldIndex := 0; fieldIndex < val.NumField(); fieldIndex++ {
			field := val.Field(fieldIndex) // field — тоже Value

			if val.Type().Field(fieldIndex).Name == spell.Char() {
				field.SetInt(field.Int() + int64(spell.Value()))
			}
		}
	}
}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}

func main() {

	player := &Player{
		name:   "Player_1",
		health: 100,
	}

	enemies := []any{
		&Zombie{Health: 1000},
		&Zombie{Health: 1000},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Daemon{Health: 1000},
		&Daemon{Health: 1000},
		&Wall{Durability: 100},
	}

	CastToAll(newSpell("fire", "Health", -50), append(enemies, player))
	CastToAll(newSpell("heal", "Health", 190), append(enemies, player))

	fmt.Println(player)
}
