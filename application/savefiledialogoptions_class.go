package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewSaveFileDialogOptionsClass() data.ClassStmt {
	return &SaveFileDialogOptionsClass{
		source:    nil,
		construct: &SaveFileDialogOptionsConstructMethod{source: nil},
	}
}

func NewSaveFileDialogOptionsClassFrom(source *applicationsrc.SaveFileDialogOptions) data.ClassStmt {
	return &SaveFileDialogOptionsClass{
		source:    source,
		construct: &SaveFileDialogOptionsConstructMethod{source: source},
	}
}

type SaveFileDialogOptionsClass struct {
	node.Node
	source    *applicationsrc.SaveFileDialogOptions
	construct data.Method
}

func (s *SaveFileDialogOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewSaveFileDialogOptionsClassFrom(&applicationsrc.SaveFileDialogOptions{}), ctx.CreateBaseContext()), nil
}
func (s *SaveFileDialogOptionsClass) GetName() string {
	return "wails\\application\\SaveFileDialogOptions"
}
func (s *SaveFileDialogOptionsClass) GetExtend() *string      { return nil }
func (s *SaveFileDialogOptionsClass) GetImplements() []string { return nil }
func (s *SaveFileDialogOptionsClass) AsString() string        { return "SaveFileDialogOptions{}" }
func (s *SaveFileDialogOptionsClass) GetSource() any          { return s.source }
func (s *SaveFileDialogOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "CanCreateDirectories":
		return node.NewProperty(nil, "CanCreateDirectories", "public", true, data.NewAnyValue(s.source.CanCreateDirectories)), true
	case "ShowHiddenFiles":
		return node.NewProperty(nil, "ShowHiddenFiles", "public", true, data.NewAnyValue(s.source.ShowHiddenFiles)), true
	case "CanSelectHiddenExtension":
		return node.NewProperty(nil, "CanSelectHiddenExtension", "public", true, data.NewAnyValue(s.source.CanSelectHiddenExtension)), true
	case "AllowOtherFileTypes":
		return node.NewProperty(nil, "AllowOtherFileTypes", "public", true, data.NewAnyValue(s.source.AllowOtherFileTypes)), true
	case "HideExtension":
		return node.NewProperty(nil, "HideExtension", "public", true, data.NewAnyValue(s.source.HideExtension)), true
	case "TreatsFilePackagesAsDirectories":
		return node.NewProperty(nil, "TreatsFilePackagesAsDirectories", "public", true, data.NewAnyValue(s.source.TreatsFilePackagesAsDirectories)), true
	case "Title":
		return node.NewProperty(nil, "Title", "public", true, data.NewAnyValue(s.source.Title)), true
	case "Message":
		return node.NewProperty(nil, "Message", "public", true, data.NewAnyValue(s.source.Message)), true
	case "Directory":
		return node.NewProperty(nil, "Directory", "public", true, data.NewAnyValue(s.source.Directory)), true
	case "Filename":
		return node.NewProperty(nil, "Filename", "public", true, data.NewAnyValue(s.source.Filename)), true
	case "ButtonText":
		return node.NewProperty(nil, "ButtonText", "public", true, data.NewAnyValue(s.source.ButtonText)), true
	case "Filters":
		return node.NewProperty(nil, "Filters", "public", true, data.NewAnyValue(s.source.Filters)), true
	case "Window":
		return node.NewProperty(nil, "Window", "public", true, data.NewClassValue(NewWindowClassFrom(s.source.Window), runtime.NewContextToDo())), true
	}
	return nil, false
}

func (s *SaveFileDialogOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"CanCreateDirectories":            node.NewProperty(nil, "CanCreateDirectories", "public", true, data.NewAnyValue(nil)),
		"ShowHiddenFiles":                 node.NewProperty(nil, "ShowHiddenFiles", "public", true, data.NewAnyValue(nil)),
		"CanSelectHiddenExtension":        node.NewProperty(nil, "CanSelectHiddenExtension", "public", true, data.NewAnyValue(nil)),
		"AllowOtherFileTypes":             node.NewProperty(nil, "AllowOtherFileTypes", "public", true, data.NewAnyValue(nil)),
		"HideExtension":                   node.NewProperty(nil, "HideExtension", "public", true, data.NewAnyValue(nil)),
		"TreatsFilePackagesAsDirectories": node.NewProperty(nil, "TreatsFilePackagesAsDirectories", "public", true, data.NewAnyValue(nil)),
		"Title":                           node.NewProperty(nil, "Title", "public", true, data.NewAnyValue(nil)),
		"Message":                         node.NewProperty(nil, "Message", "public", true, data.NewAnyValue(nil)),
		"Directory":                       node.NewProperty(nil, "Directory", "public", true, data.NewAnyValue(nil)),
		"Filename":                        node.NewProperty(nil, "Filename", "public", true, data.NewAnyValue(nil)),
		"ButtonText":                      node.NewProperty(nil, "ButtonText", "public", true, data.NewAnyValue(nil)),
		"Filters":                         node.NewProperty(nil, "Filters", "public", true, data.NewAnyValue(nil)),
		"Window":                          node.NewProperty(nil, "Window", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
	}
}

func (s *SaveFileDialogOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
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
	case "CanSelectHiddenExtension":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.CanSelectHiddenExtension = bool(x)
			}
		}
	case "AllowOtherFileTypes":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.AllowOtherFileTypes = bool(x)
			}
		}
	case "HideExtension":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.HideExtension = bool(x)
			}
		}
	case "TreatsFilePackagesAsDirectories":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.TreatsFilePackagesAsDirectories = bool(x)
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
	case "Directory":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Directory = string(sv.AsString())
		}
	case "Filename":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Filename = string(sv.AsString())
		}
	case "ButtonText":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.ButtonText = string(sv.AsString())
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
	}
}

func (s *SaveFileDialogOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *SaveFileDialogOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *SaveFileDialogOptionsClass) GetConstruct() data.Method { return s.construct }

type SaveFileDialogOptionsConstructMethod struct {
	source *applicationsrc.SaveFileDialogOptions
}

func (h *SaveFileDialogOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *SaveFileDialogOptionsConstructMethod) GetName() string { return "construct" }
func (h *SaveFileDialogOptionsConstructMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogOptionsConstructMethod) GetIsStatic() bool          { return false }
func (h *SaveFileDialogOptionsConstructMethod) GetParams() []data.GetValue { return []data.GetValue{} }
func (h *SaveFileDialogOptionsConstructMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *SaveFileDialogOptionsConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
