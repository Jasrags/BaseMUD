package core

type AttributeFormula struct {
	Requires []string
	Formula  func()
}

func NewAttributeFormula() *AttributeFormula {
	return &AttributeFormula{}
}
