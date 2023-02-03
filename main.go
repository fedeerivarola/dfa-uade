package main

import "fmt"

func main() {
	println("L = { x / x ∈ {0,1}*, existe la subcadena 00 o 11 }")

	s1 := Symbol{value: "0"}
	s2 := Symbol{value: "1"}
	symbols := make([]Symbol, 2)
	symbols[0] = s1
	symbols[1] = s2

	alphabet := Σ{symbols}

	dfa := NewAutomaton(alphabet)
	rbys := make(map[string][]Restriction)

	rbys[s1.value] = []Restriction{ContainsSubstringRestriction{"subcadena 00", "00"}}
	rbys[s2.value] = []Restriction{ContainsSubstringRestriction{"subcadena 11", "11"}}

	stackMemory := make([]Symbol, 1000)
	dfa.initialize(dfa.s, stackMemory)

	for i := 0; i < len(dfa.K); i++ {
		fmt.Printf("Estado %d	|\n", i)
		for j := 0; j < len(dfa.K[i].transitions); j++ {
			//δ := dfa.K[i].transitions[j]
			//fmt.Printf("δ: %s x %s → %s", δ.q., δ.z, δ.result)
			//TODO NAME STATES
		}
	}
}
