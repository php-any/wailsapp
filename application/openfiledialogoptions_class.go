package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewOpenFileDialogOptionsClass() data.ClassStmt {
	return &OpenFileDialogOptionsClass{
		source:    nil,
		construct: &OpenFileDialogOptionsConstructMethod{source: nil},
	}
}

func NewOpenFileDialogOptionsClassFrom(source *applicationsrc.OpenFileDialogOptions) data.ClassStmt {
	return &OpenFileDialogOptionsClass{
		source:    source,
		construct: &OpenFileDialogOptionsConstructMethod{source: source},
	}
}

type OpenFileDialogOptionsClass struct {
	node.Node
	source    *applicationsrc.OpenFileDialogOptions
	construct data.Method
}

func (s *OpenFileDialogOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewOpenFileDialogOptionsClassFrom(&applicationsrc.OpenFileDialogOptions{}), ctx.CreateBaseContext()), nil
}
func (s *OpenFileDialogOptionsClass) GetName() string {
	return "wails\\application\\OpenFileDialogOptions"
}
func (s *OpenFileDialogOptionsClass) GetExtend() *string      { return nil }
func (s *OpenFileDialogOptionsClass) GetImplements() []string { return nil }
func (s *OpenFileDialogOptionsClass) AsString() string        { return "OpenFileDialogOptions{}" }
func (s *OpenFileDialogOptionsClass) GetSource() any          { return s.source }
func (s *OpenFileDialogOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "CanChooseDirectories":
		return node.NewProperty(nil, "CanChooseDirectories", "public", true, data.NewAnyValue(s.source.CanChooseDirectories)), true
	case "CanChooseFiles":
		return node.NewProperty(nil, "CanChooseFiles", "public", true, data.NewAnyValue(s.source.CanChooseFiles)), true
	case "CanCreateDirectories":
		return node.NewProperty(nil, "CanCreateDirectories", "public", true, data.NewAnyValue(s.source.CanCreateDirectories)), true
	case "ShowHiddenFiles":
		return node.NewProperty(nil, "ShowHiddenFiles", "public", true, data.NewAnyValue(s.source.ShowHiddenFiles)), true
	case "ResolvesAliases":
		return node.NewProperty(nil, "ResolvesAliases", "public", true, data.NewAnyValue(s.source.ResolvesAliases)), true
	case "AllowsMultipleSelection":
		return node.NewProperty(nil, "AllowsMultipleSelection", "public", true, data.NewAnyValue(s.source.AllowsMultipleSelection)), true
	case "HideExtension":
		return node.NewProperty(nil, "HideExtension", "public", true, data.NewAnyValue(s.source.HideExtension)), true
	case "CanSelectHiddenExtension":
		return node.NewProperty(nil, "CanSelectHiddenExtension", "public", true, data.NewAnyValue(s.source.CanSelectHiddenExtension)), true
	case "TreatsFilePackagesAsDirectories":
		return node.NewProperty(nil, "TreatsFilePackagesAsDirectories", "public", true, data.NewAnyValue(s.source.TreatsFilePackagesAsDirectories)), true
	case "AllowsOtherFileTypes":
		return node.NewProperty(nil, "AllowsOtherFileTypes", "public", true, data.NewAnyValue(s.source.AllowsOtherFileTypes)), true
	case "Filters":
		return node.NewProperty(nil, "Filters", "public", true, data.NewAnyValue(s.source.Filters)), true
	case "Window":
		return node.NewProperty(nil, "Window", "public", true, data.NewClassValue(NewWindowClassFrom(s.source.Window), runtime.NewContextToDo())), true
	case "Title":
		return node.NewProperty(nil, "Title", "public", true, data.NewAnyValue(s.source.Title)), true
	case "Message":
		return node.NewProperty(nil, "Message", "public", true, data.NewAnyValue(s.source.Message)), true
	case "ButtonText":
		return node.NewProperty(nil, "ButtonText", "public", true, data.NewAnyValue(s.source.ButtonText)), true
	case "Directory":
		return node.NewProperty(nil, "Directory", "public", true, data.NewAnyValue(s.source.Directory)), true
	}
	return nil, false
}

func (s *OpenFileDialogOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"CanChooseDirectories":            node.NewProperty(nil, "CanChooseDirectories", "public", true, data.NewAnyValue(nil)),
		"CanChooseFiles":                  node.NewProperty(nil, "CanChooseFiles", "public", true, data.NewAnyValue(nil)),
		"CanCreateDirectories":            node.NewProperty(nil, "CanCreateDirectories", "public", true, data.NewAnyValue(nil)),
		"ShowHiddenFiles":                 node.NewProperty(nil, "ShowHiddenFiles", "public", true, data.NewAnyValue(nil)),
		"ResolvesAliases":                 node.NewProperty(nil, "ResolvesAliases", "public", true, data.NewAnyValue(nil)),
		"AllowsMultipleSelection":         node.NewProperty(nil, "AllowsMultipleSelection", "public", true, data.NewAnyValue(nil)),
		"HideExtension":                   node.NewProperty(nil, "HideExtension", "public", true, data.NewAnyValue(nil)),
		"CanSelectHiddenExtension":        node.NewProperty(nil, "CanSelectHiddenExtension", "public", true, data.NewAnyValue(nil)),
		"TreatsFilePackagesAsDirectories": node.NewProperty(nil, "TreatsFilePackagesAsDirectories", "public", true, data.NewAnyValue(nil)),
		"AllowsOtherFileTypes":            node.NewProperty(nil, "AllowsOtherFileTypes", "public", true, data.NewAnyValue(nil)),
		"Filters":                         node.NewProperty(nil, "Filters", "public", true, data.NewAnyValue(nil)),
		"Window":                          node.NewProperty(nil, "Window", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Title":                           node.NewProperty(nil, "Title", "public", true, data.NewAnyValue(nil)),
		"Message":                         node.NewProperty(nil, "Message", "public", true, data.NewAnyValue(nil)),
		"ButtonText":                      node.NewProperty(nil, "ButtonText", "public", true, data.NewAnyValue(nil)),
		"Directory":                       node.NewProperty(nil, "Directory", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *OpenFileDialogOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "CanChooseDirectories":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.CanChooseDirectories = bool(x)
			}
		}
	case "CanChooseFiles":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.CanChooseFiles = bool(x)
			}
		}
	case "CanCreateDirectories":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.CanCreateDirectories = bool(x)
			}
		}
	case "ShowHiddenFiles":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.ShowHiddenFiles = bool(x)
			}
		}
	case "ResolvesAliases":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.ResolvesAliases = bool(x)
			}
		}
	case "AllowsMultipleSelection":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.AllowsMultipleSelection = bool(x)
			}
		}
	case "HideExtension":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.HideExtension = bool(x)
			}
		}
	case "CanSelectHiddenExtension":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.CanSelectHiddenExtension = bool(x)
			}
		}
	case "TreatsFilePackagesAsDirectories":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.TreatsFilePackagesAsDirectories = bool(x)
			}
		}
	case "AllowsOtherFileTypes":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.AllowsOtherFileTypes = bool(x)
			}
		}
	case "Filters":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Filters = v.Value.([]applicationsrc.FileFilter)
		}
	case "Window":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				s.source.Window = src.(applicationsrc.Window)
			}
		case *data.AnyValue:
			s.source.Window = v.Value.(applicationsrc.Window)
		}
	case "Title":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Title = string(sv.AsString())
		}
	case "Message":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Message = string(sv.AsString())
		}
	case "ButtonText":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.ButtonText = string(sv.AsString())
		}
	case "Directory":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Directory = string(sv.AsString())
		}
	}
}

func (s *OpenFileDialogOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *OpenFileDialogOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *OpenFileDialogOptionsClass) GetConstruct() data.Method { return s.construct }

type OpenFileDialogOptionsConstructMethod struct {
	source *applicationsrc.OpenFileDialogOptions
}

func (h *OpenFileDialogOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *OpenFileDialogOptionsConstructMethod) GetName() string { return "construct" }
func (h *OpenFileDialogOptionsConstructMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogOptionsConstructMethod) GetIsStatic() bool          { return false }
func (h *OpenFileDialogOptionsConstructMethod) GetParams() []data.GetValue { return []data.GetValue{} }
func (h *OpenFileDialogOptionsConstructMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *OpenFileDialogOptionsConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
