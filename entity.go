package gocleanarch

type Entity struct {
	id string
}

func (e *Entity) IsSame(entity *Entity) bool {
	return e.id != "" && (e.id == entity.id)
}

func (e *Entity) SetId(id string) {
	e.id = id
}

func (e *Entity) Id() string {
	return e.id
}
