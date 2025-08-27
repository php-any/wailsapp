package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMessageDialogOptionsClass() data.ClassStmt {
	return &MessageDialogOptionsClass{
		source:    nil,
		construct: &MessageDialogOptionsConstructMethod{source: nil},
	}
}

func NewMessageDialogOptionsClassFrom(source *applicationsrc.MessageDialogOptions) data.ClassStmt {
	return &MessageDialogOptionsClass{
		source:    source,
		construct: &MessageDialogOptionsConstructMethod{source: source},
	}
}

type MessageDialogOptionsClass struct {
	node.Node
	source    *applicationsrc.MessageDialogOptions
	construct data.Method
}

func (s *MessageDialogOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMessageDialogOptionsClassFrom(&applicationsrc.MessageDialogOptions{}), ctx.CreateBaseContext()), nil
}
func (s *MessageDialogOptionsClass) GetName() string {
	return "wails\\application\\MessageDialogOptions"
}
func (s *MessageDialogOptionsClass) GetExtend() *string      { return nil }
func (s *MessageDialogOptionsClass) GetImplements() []string { return nil }
func (s *MessageDialogOptionsClass) AsString() string        { return "MessageDialogOptions{}" }
func (s *MessageDialogOptionsClass) GetSource() any          { return s.source }
func (s *MessageDialogOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "DialogType":
		return node.NewProperty(nil, "DialogType", "public", true, data.NewAnyValue(s.source.DialogType)), true
	case "Title":
		return node.NewProperty(nil, "Title", "public", true, data.NewAnyValue(s.source.Title)), true
	case "Message":
		return node.NewProperty(nil, "Message", "public", true, data.NewAnyValue(s.source.Message)), true
	case "Buttons":
		return node.NewProperty(nil, "Buttons", "public", true, data.NewAnyValue(s.source.Buttons)), true
	case "Icon":
		return node.NewProperty(nil, "Icon", "public", true, data.NewAnyValue(s.source.Icon)), true
	}
	return nil, false
}

func (s *MessageDialogOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"DialogType": node.NewProperty(nil, "DialogType", "public", true, data.NewAnyValue(nil)),
		"Title":      node.NewProperty(nil, "Title", "public", true, data.NewAnyValue(nil)),
		"Message":    node.NewProperty(nil, "Message", "public", true, data.NewAnyValue(nil)),
		"Buttons":    node.NewProperty(nil, "Buttons", "public", true, data.NewAnyValue(nil)),
		"Icon":       node.NewProperty(nil, "Icon", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *MessageDialogOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "DialogType":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.DialogType = applicationsrc.DialogType(x)
			}
		}
	case "Title":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Title = string(sv.AsString())
		}
	case "Message":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Message = string(sv.AsString())
		}
	case "Buttons":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Buttons = v.Value.([]*applicationsrc.Button)
		}
	case "Icon":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Icon = v.Value.([]uint8)
		}
	}
}

func (s *MessageDialogOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *MessageDialogOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *MessageDialogOptionsClass) GetConstruct() data.Method { return s.construct }

type MessageDialogOptionsConstructMethod struct {
	source *applicationsrc.MessageDialogOptions
}

func (h *MessageDialogOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MessageDialogOptionsConstructMethod) GetName() string            { return "construct" }
func (h *MessageDialogOptionsConstructMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MessageDialogOptionsConstructMethod) GetIsStatic() bool          { return false }
func (h *MessageDialogOptionsConstructMethod) GetParams() []data.GetValue { return []data.GetValue{} }
func (h *MessageDialogOptionsConstructMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *MessageDialogOptionsConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
