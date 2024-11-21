package core

type Attribute struct {
	Name    string
	Base    float32
	Delta   float32
	Formula *AttributeFormula
	// Metadata string
}

func NewAttribute(name string, base, delta float32, formula *AttributeFormula, metadata string) *Attribute {
	return &Attribute{
		Name:    name,
		Base:    base,
		Delta:   delta,
		Formula: formula,
		// Metadata: metadata,
	}
}

func (a *Attribute) Lower(amount int) {
	// Lower current value
}

func (a *Attribute) Raise(amount int) {
	// Raise current value
}

func (a *Attribute) SetBase(amount int) {
	// Change the base value
}

func (a *Attribute) SetDelta(amount int) {
	// Bypass raise/lower, directly setting the delta
}
