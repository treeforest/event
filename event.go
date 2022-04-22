package event

type Event interface {
	Type() string
	Source() interface{}
	Value() interface{}
}

type event struct {
	typ       string
	source    interface{}
	value     interface{}
	prevValue interface{}
}

func NewEvent(typ string, value interface{}) *event {
	return &event{
		typ:   typ,
		value: value,
	}
}

func (e *event) Type() string {
	return e.typ
}

func (e *event) Source() interface{} {
	return e.source
}

func (e *event) Value() interface{} {
	return e.value
}
