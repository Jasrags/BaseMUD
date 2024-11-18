package core

type EntityFactory struct {
}

func NewEntityFactory() *EntityFactory {
	return &EntityFactory{}
}

func (ef *EntityFactory) AddScriptListener(entityRef string, event string, listener func()) {
}

func (ef *EntityFactory) Clone(entity interface{}) interface{} {
	// {Item|Npc|Room|Area}
	return nil
}

func (ef *EntityFactory) CreateByType(area Area, entityRef string, Type interface{}) interface{} {
	// {type}
	return nil
}

func (ef *EntityFactory) CreateEntityRef(areaName string, id int) string {
	return ""
}

func (ef *EntityFactory) GetDefinition(entityRef string) interface{} {
	return nil
}

func (ef *EntityFactory) SetDefinition(entityRef string, def interface{}) {
}
