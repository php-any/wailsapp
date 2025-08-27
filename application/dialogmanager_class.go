package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewDialogManagerClass() data.ClassStmt {
	return &DialogManagerClass{
		source:              nil,
		construct:           &DialogManagerConstructMethod{source: nil},
		error:               &DialogManagerErrorMethod{source: nil},
		info:                &DialogManagerInfoMethod{source: nil},
		openFile:            &DialogManagerOpenFileMethod{source: nil},
		openFileWithOptions: &DialogManagerOpenFileWithOptionsMethod{source: nil},
		question:            &DialogManagerQuestionMethod{source: nil},
		saveFile:            &DialogManagerSaveFileMethod{source: nil},
		saveFileWithOptions: &DialogManagerSaveFileWithOptionsMethod{source: nil},
		warning:             &DialogManagerWarningMethod{source: nil},
	}
}

func NewDialogManagerClassFrom(source *applicationsrc.DialogManager) data.ClassStmt {
	return &DialogManagerClass{
		source:              source,
		construct:           &DialogManagerConstructMethod{source: source},
		error:               &DialogManagerErrorMethod{source: source},
		info:                &DialogManagerInfoMethod{source: source},
		openFile:            &DialogManagerOpenFileMethod{source: source},
		openFileWithOptions: &DialogManagerOpenFileWithOptionsMethod{source: source},
		question:            &DialogManagerQuestionMethod{source: source},
		saveFile:            &DialogManagerSaveFileMethod{source: source},
		saveFileWithOptions: &DialogManagerSaveFileWithOptionsMethod{source: source},
		warning:             &DialogManagerWarningMethod{source: source},
	}
}

type DialogManagerClass struct {
	node.Node
	source              *applicationsrc.DialogManager
	construct           data.Method
	error               data.Method
	info                data.Method
	openFile            data.Method
	openFileWithOptions data.Method
	question            data.Method
	saveFile            data.Method
	saveFileWithOptions data.Method
	warning             data.Method
}

func (s *DialogManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewDialogManagerClassFrom(&applicationsrc.DialogManager{}), ctx.CreateBaseContext()), nil
}
func (s *DialogManagerClass) GetName() string                            { return "wails\\application\\DialogManager" }
func (s *DialogManagerClass) GetExtend() *string                         { return nil }
func (s *DialogManagerClass) GetImplements() []string                    { return nil }
func (s *DialogManagerClass) AsString() string                           { return "DialogManager{}" }
func (s *DialogManagerClass) GetSource() any                             { return s.source }
func (s *DialogManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *DialogManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *DialogManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "error":
		return s.error, true
	case "info":
		return s.info, true
	case "openFile":
		return s.openFile, true
	case "openFileWithOptions":
		return s.openFileWithOptions, true
	case "question":
		return s.question, true
	case "saveFile":
		return s.saveFile, true
	case "saveFileWithOptions":
		return s.saveFileWithOptions, true
	case "warning":
		return s.warning, true
	}
	return nil, false
}

func (s *DialogManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.error,
		s.info,
		s.openFile,
		s.openFileWithOptions,
		s.question,
		s.saveFile,
		s.saveFileWithOptions,
		s.warning,
	}
}

func (s *DialogManagerClass) GetConstruct() data.Method { return s.construct }

type DialogManagerConstructMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *DialogManagerConstructMethod) GetName() string               { return "construct" }
func (h *DialogManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *DialogManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *DialogManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *DialogManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *DialogManagerConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
