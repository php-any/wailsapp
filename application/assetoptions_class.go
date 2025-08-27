package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	http "net/http"
)

func NewAssetOptionsClass() data.ClassStmt {
	return &AssetOptionsClass{
		source:    nil,
		construct: &AssetOptionsConstructMethod{source: nil},
	}
}

func NewAssetOptionsClassFrom(source *applicationsrc.AssetOptions) data.ClassStmt {
	return &AssetOptionsClass{
		source:    source,
		construct: &AssetOptionsConstructMethod{source: source},
	}
}

type AssetOptionsClass struct {
	node.Node
	source    *applicationsrc.AssetOptions
	construct data.Method
}

func (s *AssetOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewAssetOptionsClassFrom(&applicationsrc.AssetOptions{}), ctx.CreateBaseContext()), nil
}
func (s *AssetOptionsClass) GetName() string         { return "wails\\application\\AssetOptions" }
func (s *AssetOptionsClass) GetExtend() *string      { return nil }
func (s *AssetOptionsClass) GetImplements() []string { return nil }
func (s *AssetOptionsClass) AsString() string        { return "AssetOptions{}" }
func (s *AssetOptionsClass) GetSource() any          { return s.source }
func (s *AssetOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Handler":
		return node.NewProperty(nil, "Handler", "public", true, data.NewAnyValue(s.source.Handler)), true
	case "Middleware":
		return node.NewProperty(nil, "Middleware", "public", true, data.NewAnyValue(s.source.Middleware)), true
	case "DisableLogging":
		return node.NewProperty(nil, "DisableLogging", "public", true, data.NewAnyValue(s.source.DisableLogging)), true
	}
	return nil, false
}

func (s *AssetOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Handler":        node.NewProperty(nil, "Handler", "public", true, data.NewAnyValue(nil)),
		"Middleware":     node.NewProperty(nil, "Middleware", "public", true, data.NewAnyValue(nil)),
		"DisableLogging": node.NewProperty(nil, "DisableLogging", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *AssetOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Handler":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Handler = v.Value.(http.Handler)
		}
	case "Middleware":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Middleware = v.Value.(applicationsrc.Middleware)
		}
	case "DisableLogging":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DisableLogging = bool(x)
			}
		}
	}
}

func (s *AssetOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *AssetOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *AssetOptionsClass) GetConstruct() data.Method { return s.construct }

type AssetOptionsConstructMethod struct {
	source *applicationsrc.AssetOptions
}

func (h *AssetOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *AssetOptionsConstructMethod) GetName() string               { return "construct" }
func (h *AssetOptionsConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *AssetOptionsConstructMethod) GetIsStatic() bool             { return false }
func (h *AssetOptionsConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *AssetOptionsConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *AssetOptionsConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
