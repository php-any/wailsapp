package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewBrowserManagerClass() data.ClassStmt {
	return &BrowserManagerClass{
		source:    nil,
		construct: &BrowserManagerConstructMethod{source: nil},
		openFile:  &BrowserManagerOpenFileMethod{source: nil},
		openURL:   &BrowserManagerOpenURLMethod{source: nil},
	}
}

func NewBrowserManagerClassFrom(source *applicationsrc.BrowserManager) data.ClassStmt {
	return &BrowserManagerClass{
		source:    source,
		construct: &BrowserManagerConstructMethod{source: source},
		openFile:  &BrowserManagerOpenFileMethod{source: source},
		openURL:   &BrowserManagerOpenURLMethod{source: source},
	}
}

type BrowserManagerClass struct {
	node.Node
	source    *applicationsrc.BrowserManager
	construct data.Method
	openFile  data.Method
	openURL   data.Method
}

func (s *BrowserManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewBrowserManagerClassFrom(&applicationsrc.BrowserManager{}), ctx.CreateBaseContext()), nil
}
func (s *BrowserManagerClass) GetName() string                            { return "wails\\application\\BrowserManager" }
func (s *BrowserManagerClass) GetExtend() *string                         { return nil }
func (s *BrowserManagerClass) GetImplements() []string                    { return nil }
func (s *BrowserManagerClass) AsString() string                           { return "BrowserManager{}" }
func (s *BrowserManagerClass) GetSource() any                             { return s.source }
func (s *BrowserManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *BrowserManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *BrowserManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "openFile":
		return s.openFile, true
	case "openURL":
		return s.openURL, true
	}
	return nil, false
}

func (s *BrowserManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.openFile,
		s.openURL,
	}
}

func (s *BrowserManagerClass) GetConstruct() data.Method { return s.construct }

type BrowserManagerConstructMethod struct {
	source *applicationsrc.BrowserManager
}

func (h *BrowserManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *BrowserManagerConstructMethod) GetName() string               { return "construct" }
func (h *BrowserManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *BrowserManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *BrowserManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *BrowserManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *BrowserManagerConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
