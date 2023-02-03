package main

// δ(q,z)->result
type δ struct {
	q      *State
	z      Symbol
	result *State
}

type Restriction interface {
	apply(stack []Symbol, z Symbol) (bool, error)
}

func (t δ) execute(s Symbol) *State {
	//result, err := t.r.apply(t.q.stack, t.z)
	if s == t.z {
		return t.result
	}

	return nil
}
