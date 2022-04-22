package event

import (
	"reflect"
	"sync"
)

type Dispatcher interface {
	AddEventListener(string, Listener)
	RemoveEventListener(string, Listener)
	DispatchEvent(Event)
}

type eventDispatcher struct {
	sync.RWMutex
	source    interface{}
	listeners map[string]eventListeners
}

type Listener func(Event)

type eventListeners []Listener

func NewEventDispatcher(source interface{}) Dispatcher {
	return &eventDispatcher{
		source:    source,
		listeners: make(map[string]eventListeners),
	}
}

func (d *eventDispatcher) AddEventListener(typ string, listener Listener) {
	d.Lock()
	defer d.Unlock()
	d.listeners[typ] = append(d.listeners[typ], listener)
}

func (d *eventDispatcher) RemoveEventListener(typ string, listener Listener) {
	d.Lock()
	defer d.Unlock()

	ptr := reflect.ValueOf(listener).Pointer()

	listeners := d.listeners[typ]
	for i, l := range listeners {
		if reflect.ValueOf(l).Pointer() == ptr {
			d.listeners[typ] = append(listeners[:i], listeners[i+1:]...)
		}
	}
}

func (d *eventDispatcher) DispatchEvent(e Event) {
	d.RLock()
	defer d.RUnlock()

	if ev, ok := e.(*event); ok {
		ev.source = d.source
	}

	// 派发事件给监听者
	for _, l := range d.listeners[e.Type()] {
		l(e)
	}
}
