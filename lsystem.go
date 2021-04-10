package lsystem

type LSystem struct {
	Variables string
	Constants string
	Axiom     string
	Rules     map[rune]string
}

func (lsystem LSystem) Generation(iterations int) string {
	output := lsystem.Axiom
	for i := 0; i < iterations; i++ {
		output = lsystem.Iterate(output)
	}
	return string(output)
}

func (lsystem LSystem) Iterate(input string) string {
	var output string
	for _, symbol := range input {
		if rule, ok := lsystem.Rules[symbol]; ok {
			output += rule
			continue
		}
		output += string(symbol)
	}
	return output
}

func Process(output string, operations map[rune]func()) {
	for _, symbol := range output {
		operation := operations[rune(symbol)]
		if operation == nil {
			continue
		}
		operation()
	}
}
