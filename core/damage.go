package core

type Damage struct {
}

func NewDamage() *Damage {
	return &Damage{}
}

func (d *Damage) Commit(target Character) {
	// Character#event:hit
	// Character#event:damaged

	panic("unimplemented")
}

func (d *Damage) Evaluate(target Character) int {
	panic("unimplemented")
}
