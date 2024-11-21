package core

type Attributes struct {
	Attributes map[string]*Attribute
}

func NewAttributes() *Attributes {
	return &Attributes{
		Attributes: make(map[string]*Attribute),
	}
}

func (a *Attributes) Add(attribute *Attribute) {
	a.Attributes[attribute.Name] = attribute
}

func (a *Attributes) ClearDeltas() {
	for _, attribute := range a.Attributes {
		attribute.Delta = 0
	}
}

func (a *Attributes) GetAttributes() map[string]*Attribute {
	return a.Attributes
}

func (a *Attributes) Serialize() string {
	return ""
}
