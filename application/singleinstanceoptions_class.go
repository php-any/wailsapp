package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewSingleInstanceOptionsClass() data.ClassStmt {
	return &SingleInstanceOptionsClass{
		source:    nil,
		construct: &SingleInstanceOptionsConstructMethod{source: nil},
	}
}

func NewSingleInstanceOptionsClassFrom(source *applicationsrc.SingleInstanceOptions) data.ClassStmt {
	return &SingleInstanceOptionsClass{
		source:    source,
		construct: &SingleInstanceOptionsConstructMethod{source: source},
	}
}

type SingleInstanceOptionsClass struct {
	node.Node
	source    *applicationsrc.SingleInstanceOptions
	construct data.Method
}

func (s *SingleInstanceOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewSingleInstanceOptionsClassFrom(&applicationsrc.SingleInstanceOptions{}), ctx.CreateBaseContext()), nil
}
func (s *SingleInstanceOptionsClass) GetName() string {
	return "wails\\application\\SingleInstanceOptions"
}
func (s *SingleInstanceOptionsClass) GetExtend() *string      { return nil }
func (s *SingleInstanceOptionsClass) GetImplements() []string { return nil }
func (s *SingleInstanceOptionsClass) AsString() string        { return "SingleInstanceOptions{}" }
func (s *SingleInstanceOptionsClass) GetSource() any          { return s.source }
func (s *SingleInstanceOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "UniqueID":
		return node.NewProperty(nil, "UniqueID", "public", true, data.NewAnyValue(s.source.UniqueID)), true
	case "OnSecondInstanceLaunch":
		return node.NewProperty(nil, "OnSecondInstanceLaunch", "public", true, data.NewAnyValue(s.source.OnSecondInstanceLaunch)), true
	case "AdditionalData":
		return node.NewProperty(nil, "AdditionalData", "public", true, data.NewAnyValue(s.source.AdditionalData)), true
	case "ExitCode":
		return node.NewProperty(nil, "ExitCode", "public", true, data.NewAnyValue(s.source.ExitCode)), true
	case "EncryptionKey":
		return node.NewProperty(nil, "EncryptionKey", "public", true, data.NewAnyValue(s.source.EncryptionKey)), true
	}
	return nil, false
}

func (s *SingleInstanceOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"UniqueID":               node.NewProperty(nil, "UniqueID", "public", true, data.NewAnyValue(nil)),
		"OnSecondInstanceLaunch": node.NewProperty(nil, "OnSecondInstanceLaunch", "public", true, data.NewAnyValue(nil)),
		"AdditionalData":         node.NewProperty(nil, "AdditionalData", "public", true, data.NewAnyValue(nil)),
		"ExitCode":               node.NewProperty(nil, "ExitCode", "public", true, data.NewAnyValue(nil)),
		"EncryptionKey":          node.NewProperty(nil, "EncryptionKey", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *SingleInstanceOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "UniqueID":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.UniqueID = string(sv.AsString())
		}
	case "OnSecondInstanceLaunch":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.OnSecondInstanceLaunch = v.Value.(func(applicationsrc.SecondInstanceData))
		}
	case "AdditionalData":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.AdditionalData = v.Value.(map[string]string)
		}
	case "ExitCode":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.ExitCode = int(x)
			}
		}
	case "EncryptionKey":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.EncryptionKey = v.Value.([32]uint8)
		}
	}
}

func (s *SingleInstanceOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *SingleInstanceOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *SingleInstanceOptionsClass) GetConstruct() data.Method { return s.construct }

type SingleInstanceOptionsConstructMethod struct {
	source *applicationsrc.SingleInstanceOptions
}

func (h *SingleInstanceOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *SingleInstanceOptionsConstructMethod) GetName() string { return "construct" }
func (h *SingleInstanceOptionsConstructMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SingleInstanceOptionsConstructMethod) GetIsStatic() bool          { return false }
func (h *SingleInstanceOptionsConstructMethod) GetParams() []data.GetValue { return []data.GetValue{} }
func (h *SingleInstanceOptionsConstructMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *SingleInstanceOptionsConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
