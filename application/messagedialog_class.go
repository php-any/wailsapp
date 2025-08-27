package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMessageDialogClass() data.ClassStmt {
	return &MessageDialogClass{
		source:           nil,
		construct:        &MessageDialogConstructMethod{source: nil},
		addButton:        &MessageDialogAddButtonMethod{source: nil},
		addButtons:       &MessageDialogAddButtonsMethod{source: nil},
		attachToWindow:   &MessageDialogAttachToWindowMethod{source: nil},
		setCancelButton:  &MessageDialogSetCancelButtonMethod{source: nil},
		setDefaultButton: &MessageDialogSetDefaultButtonMethod{source: nil},
		setIcon:          &MessageDialogSetIconMethod{source: nil},
		setMessage:       &MessageDialogSetMessageMethod{source: nil},
		setTitle:         &MessageDialogSetTitleMethod{source: nil},
		show:             &MessageDialogShowMethod{source: nil},
	}
}

func NewMessageDialogClassFrom(source *applicationsrc.MessageDialog) data.ClassStmt {
	return &MessageDialogClass{
		source:           source,
		construct:        &MessageDialogConstructMethod{source: source},
		addButton:        &MessageDialogAddButtonMethod{source: source},
		addButtons:       &MessageDialogAddButtonsMethod{source: source},
		attachToWindow:   &MessageDialogAttachToWindowMethod{source: source},
		setCancelButton:  &MessageDialogSetCancelButtonMethod{source: source},
		setDefaultButton: &MessageDialogSetDefaultButtonMethod{source: source},
		setIcon:          &MessageDialogSetIconMethod{source: source},
		setMessage:       &MessageDialogSetMessageMethod{source: source},
		setTitle:         &MessageDialogSetTitleMethod{source: source},
		show:             &MessageDialogShowMethod{source: source},
	}
}

type MessageDialogClass struct {
	node.Node
	source           *applicationsrc.MessageDialog
	construct        data.Method
	addButton        data.Method
	addButtons       data.Method
	attachToWindow   data.Method
	setCancelButton  data.Method
	setDefaultButton data.Method
	setIcon          data.Method
	setMessage       data.Method
	setTitle         data.Method
	show             data.Method
}

func (s *MessageDialogClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMessageDialogClassFrom(&applicationsrc.MessageDialog{}), ctx.CreateBaseContext()), nil
}
func (s *MessageDialogClass) GetName() string         { return "wails\\application\\MessageDialog" }
func (s *MessageDialogClass) GetExtend() *string      { return nil }
func (s *MessageDialogClass) GetImplements() []string { return nil }
func (s *MessageDialogClass) AsString() string        { return "MessageDialog{}" }
func (s *MessageDialogClass) GetSource() any          { return s.source }
func (s *MessageDialogClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "MessageDialogOptions":
		return node.NewProperty(nil, "MessageDialogOptions", "public", true, data.NewClassValue(NewMessageDialogOptionsClassFrom(&s.source.MessageDialogOptions), runtime.NewContextToDo())), true
	}
	return nil, false
}

func (s *MessageDialogClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"MessageDialogOptions": node.NewProperty(nil, "MessageDialogOptions", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
	}
}

func (s *MessageDialogClass) SetProperty(name string, value data.Value) {
	switch name {
	case "MessageDialogOptions":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MessageDialogOptions); ok {
					s.source.MessageDialogOptions = *ptr
				} else {
					s.source.MessageDialogOptions = src.(applicationsrc.MessageDialogOptions)
				}
			}
		case *data.AnyValue:
			s.source.MessageDialogOptions = v.Value.(applicationsrc.MessageDialogOptions)
		}
	}
}

func (s *MessageDialogClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "addButton":
		return s.addButton, true
	case "addButtons":
		return s.addButtons, true
	case "attachToWindow":
		return s.attachToWindow, true
	case "setCancelButton":
		return s.setCancelButton, true
	case "setDefaultButton":
		return s.setDefaultButton, true
	case "setIcon":
		return s.setIcon, true
	case "setMessage":
		return s.setMessage, true
	case "setTitle":
		return s.setTitle, true
	case "show":
		return s.show, true
	}
	return nil, false
}

func (s *MessageDialogClass) GetMethods() []data.Method {
	return []data.Method{
		s.addButton,
		s.addButtons,
		s.attachToWindow,
		s.setCancelButton,
		s.setDefaultButton,
		s.setIcon,
		s.setMessage,
		s.setTitle,
		s.show,
	}
}

func (s *MessageDialogClass) GetConstruct() data.Method { return s.construct }

type MessageDialogConstructMethod struct {
	source *applicationsrc.MessageDialog
}

func (h *MessageDialogConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MessageDialogConstructMethod) GetName() string               { return "construct" }
func (h *MessageDialogConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MessageDialogConstructMethod) GetIsStatic() bool             { return false }
func (h *MessageDialogConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MessageDialogConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MessageDialogConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
