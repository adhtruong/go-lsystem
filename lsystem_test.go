package lsystem

import (
	"reflect"
	"testing"
)

var system = LSystem{
	"AB",
	"",
	"A",
	map[rune]string{
		'A': "AB", 'B': "A",
	},
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func TestIterate(t *testing.T) {
	assertEqual(t, "", system.Iterate(""))
	assertEqual(t, "AAB", system.Iterate("BA"))
}

func TestGeneration(t *testing.T) {
	assertEqual(t, "A", system.Generation(0))
	assertEqual(t, "AB", system.Generation(1))
	assertEqual(t, "ABA", system.Generation(2))
	assertEqual(t, "ABAAB", system.Generation(3))
}

func TestProcess(t *testing.T) {
	var output string
	Process(
		"ABA",
		map[rune]func(){
			'A': func() {
				output += "A"
			},
			'B': func() {
				output += "B"
			},
		},
	)
	assertEqual(t, "ABA", output)
}
