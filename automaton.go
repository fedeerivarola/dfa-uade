package main

import "fmt"

type State struct {
	//transition δ
	stack       []Symbol
	transitions []δ
	err         bool
}

func NewState(stack []Symbol) *State {
	return &State{stack: stack, transitions: nil, err: false}
}

type Stack interface {
	[]Symbol
	add(symbol Symbol)
	remove(symbol Symbol)
}

type Symbol struct {
	value string
}

func (s State) printStack() string {
	textStack := ""
	for i := 0; i < len(s.stack); i++ {
		textStack += fmt.Sprintf("Pos %d: %s | ", i, s.stack[i])
	}

	return textStack
}

type Σ struct {
	symbols []Symbol //symbol sets
}

type Automaton struct {
	K            []State //states set
	alphabet     Σ
	s            State   //Initial state
	F            []State //final states set
	transitions  []δ
	restrictions map[string][]Restriction
}

func NewAutomaton(alphabet Σ) *Automaton {
	return &Automaton{
		K:        nil,
		alphabet: alphabet,
		s: State{
			stack:       nil,
			transitions: nil},
		F:           nil,
		transitions: nil,
	}
}

func (a Automaton) initialize(state State, stack []Symbol) []Symbol {
	symbols := a.alphabet.symbols
	actualStack := stack

	for i := 0; i < len(symbols); i++ {
		//get symbol 'z'
		z := symbols[i]
		//get restrictions for symbol 'z'
		conditionsBySymbol := a.restrictions[z.value]

		//get symbol 'y' for evaluate transition
		for j := 0; j < len(symbols); j++ {
			//get symbol 'y'
			y := symbols[j]

			//get restrictions for symbol
			for c := 0; c < len(conditionsBySymbol); c++ {
				restriction := conditionsBySymbol[c]

				//validate restriction
				isAccepted, _ := restriction.apply(actualStack, y)

				//Add new state
				//Add transition to state
				newState := NewState(actualStack)
				a.K = append(a.K, *newState)

				if isAccepted {
					state.stack = actualStack
					state.transitions = append(state.transitions, δ{
						q:      &state,
						z:      y,
						result: newState,
					})

					//call recursive for initialize new state
					a.initialize(*newState, actualStack)
				} else {
					//error state
					newState.err = true
					state.transitions = append(state.transitions, δ{
						q:      &state,
						z:      y,
						result: newState,
					})
				}
			}
		}

		//run state with new stack
		go a.initialize(state, append(actualStack, z))
	}
}
