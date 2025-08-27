package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewWindowEventContextClass() data.ClassStmt {
	return &WindowEventContextClass{
		source:          nil,
		construct:       &WindowEventContextConstructMethod{source: nil},
		dropZoneDetails: &WindowEventContextDropZoneDetailsMethod{source: nil},
		droppedFiles:    &WindowEventContextDroppedFilesMethod{source: nil},
	}
}

func NewWindowEventContextClassFrom(source *applicationsrc.WindowEventContext) data.ClassStmt {
	return &WindowEventContextClass{
		source:          source,
		construct:       &WindowEventContextConstructMethod{source: source},
		dropZoneDetails: &WindowEventContextDropZoneDetailsMethod{source: source},
		droppedFiles:    &WindowEventContextDroppedFilesMethod{source: source},
	}
}

type WindowEventContextClass struct {
	node.Node
	source          *applicationsrc.WindowEventContext
	construct       data.Method
	dropZoneDetails data.Method
	droppedFiles    data.Method
}

func (s *WindowEventContextClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewWindowEventContextClassFrom(&applicationsrc.WindowEventContext{}), ctx.CreateBaseContext()), nil
}
func (s *WindowEventContextClass) GetName() string                            { return "wails\\application\\WindowEventContext" }
func (s *WindowEventContextClass) GetExtend() *string                         { return nil }
func (s *WindowEventContextClass) GetImplements() []string                    { return nil }
func (s *WindowEventContextClass) AsString() string                           { return "WindowEventContext{}" }
func (s *WindowEventContextClass) GetSource() any                             { return s.source }
func (s *WindowEventContextClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *WindowEventContextClass) GetProperties() map[string]data.Property    { return nil }
func (s *WindowEventContextClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "dropZoneDetails":
		return s.dropZoneDetails, true
	case "droppedFiles":
		return s.droppedFiles, true
	}
	return nil, false
}

func (s *WindowEventContextClass) GetMethods() []data.Method {
	return []data.Method{
		s.dropZoneDetails,
		s.droppedFiles,
	}
}

func (s *WindowEventContextClass) GetConstruct() data.Method { return s.construct }

type WindowEventContextConstructMethod struct {
	source *applicationsrc.WindowEventContext
}

func (h *WindowEventContextConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WindowEventContextConstructMethod) GetName() string               { return "construct" }
func (h *WindowEventContextConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *WindowEventContextConstructMethod) GetIsStatic() bool             { return false }
func (h *WindowEventContextConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *WindowEventContextConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *WindowEventContextConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
