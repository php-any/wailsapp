package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewEventManagerClass() data.ClassStmt {
	return &EventManagerClass{
		source:                       nil,
		construct:                    &EventManagerConstructMethod{source: nil},
		emit:                         &EventManagerEmitMethod{source: nil},
		emitEvent:                    &EventManagerEmitEventMethod{source: nil},
		off:                          &EventManagerOffMethod{source: nil},
		on:                           &EventManagerOnMethod{source: nil},
		onApplicationEvent:           &EventManagerOnApplicationEventMethod{source: nil},
		onMultiple:                   &EventManagerOnMultipleMethod{source: nil},
		registerApplicationEventHook: &EventManagerRegisterApplicationEventHookMethod{source: nil},
		reset:                        &EventManagerResetMethod{source: nil},
	}
}

func NewEventManagerClassFrom(source *applicationsrc.EventManager) data.ClassStmt {
	return &EventManagerClass{
		source:                       source,
		construct:                    &EventManagerConstructMethod{source: source},
		emit:                         &EventManagerEmitMethod{source: source},
		emitEvent:                    &EventManagerEmitEventMethod{source: source},
		off:                          &EventManagerOffMethod{source: source},
		on:                           &EventManagerOnMethod{source: source},
		onApplicationEvent:           &EventManagerOnApplicationEventMethod{source: source},
		onMultiple:                   &EventManagerOnMultipleMethod{source: source},
		registerApplicationEventHook: &EventManagerRegisterApplicationEventHookMethod{source: source},
		reset:                        &EventManagerResetMethod{source: source},
	}
}

type EventManagerClass struct {
	node.Node
	source                       *applicationsrc.EventManager
	construct                    data.Method
	emit                         data.Method
	emitEvent                    data.Method
	off                          data.Method
	on                           data.Method
	onApplicationEvent           data.Method
	onMultiple                   data.Method
	registerApplicationEventHook data.Method
	reset                        data.Method
}

func (s *EventManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewEventManagerClassFrom(&applicationsrc.EventManager{}), ctx.CreateBaseContext()), nil
}
func (s *EventManagerClass) GetName() string                            { return "wails\\application\\EventManager" }
func (s *EventManagerClass) GetExtend() *string                         { return nil }
func (s *EventManagerClass) GetImplements() []string                    { return nil }
func (s *EventManagerClass) AsString() string                           { return "EventManager{}" }
func (s *EventManagerClass) GetSource() any                             { return s.source }
func (s *EventManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *EventManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *EventManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "emit":
		return s.emit, true
	case "emitEvent":
		return s.emitEvent, true
	case "off":
		return s.off, true
	case "on":
		return s.on, true
	case "onApplicationEvent":
		return s.onApplicationEvent, true
	case "onMultiple":
		return s.onMultiple, true
	case "registerApplicationEventHook":
		return s.registerApplicationEventHook, true
	case "reset":
		return s.reset, true
	}
	return nil, false
}

func (s *EventManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.emit,
		s.emitEvent,
		s.off,
		s.on,
		s.onApplicationEvent,
		s.onMultiple,
		s.registerApplicationEventHook,
		s.reset,
	}
}

func (s *EventManagerClass) GetConstruct() data.Method { return s.construct }

type EventManagerConstructMethod struct {
	source *applicationsrc.EventManager
}

func (h *EventManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *EventManagerConstructMethod) GetName() string               { return "construct" }
func (h *EventManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *EventManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *EventManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *EventManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *EventManagerConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
