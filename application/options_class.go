package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	slog "log/slog"
)

func NewOptionsClass() data.ClassStmt {
	return &OptionsClass{
		source:    nil,
		construct: &OptionsConstructMethod{source: nil},
	}
}

func NewOptionsClassFrom(source *applicationsrc.Options) data.ClassStmt {
	return &OptionsClass{
		source:    source,
		construct: &OptionsConstructMethod{source: source},
	}
}

type OptionsClass struct {
	node.Node
	source    *applicationsrc.Options
	construct data.Method
}

func (s *OptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewOptionsClassFrom(&applicationsrc.Options{}), ctx.CreateBaseContext()), nil
}
func (s *OptionsClass) GetName() string         { return "wails\\application\\Options" }
func (s *OptionsClass) GetExtend() *string      { return nil }
func (s *OptionsClass) GetImplements() []string { return nil }
func (s *OptionsClass) AsString() string        { return "Options{}" }
func (s *OptionsClass) GetSource() any          { return s.source }
func (s *OptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Name":
		return node.NewProperty(nil, "Name", "public", true, data.NewAnyValue(s.source.Name)), true
	case "Description":
		return node.NewProperty(nil, "Description", "public", true, data.NewAnyValue(s.source.Description)), true
	case "Icon":
		return node.NewProperty(nil, "Icon", "public", true, data.NewAnyValue(s.source.Icon)), true
	case "Mac":
		return node.NewProperty(nil, "Mac", "public", true, data.NewClassValue(NewMacOptionsClassFrom(&s.source.Mac), runtime.NewContextToDo())), true
	case "Windows":
		return node.NewProperty(nil, "Windows", "public", true, data.NewClassValue(NewWindowsOptionsClassFrom(&s.source.Windows), runtime.NewContextToDo())), true
	case "Linux":
		return node.NewProperty(nil, "Linux", "public", true, data.NewClassValue(NewLinuxOptionsClassFrom(&s.source.Linux), runtime.NewContextToDo())), true
	case "Services":
		return node.NewProperty(nil, "Services", "public", true, data.NewAnyValue(s.source.Services)), true
	case "MarshalError":
		return node.NewProperty(nil, "MarshalError", "public", true, data.NewAnyValue(s.source.MarshalError)), true
	case "BindAliases":
		return node.NewProperty(nil, "BindAliases", "public", true, data.NewAnyValue(s.source.BindAliases)), true
	case "Logger":
		return node.NewProperty(nil, "Logger", "public", true, data.NewAnyValue(s.source.Logger)), true
	case "LogLevel":
		return node.NewProperty(nil, "LogLevel", "public", true, data.NewAnyValue(s.source.LogLevel)), true
	case "Assets":
		return node.NewProperty(nil, "Assets", "public", true, data.NewClassValue(NewAssetOptionsClassFrom(&s.source.Assets), runtime.NewContextToDo())), true
	case "Flags":
		return node.NewProperty(nil, "Flags", "public", true, data.NewAnyValue(s.source.Flags)), true
	case "PanicHandler":
		return node.NewProperty(nil, "PanicHandler", "public", true, data.NewAnyValue(s.source.PanicHandler)), true
	case "DisableDefaultSignalHandler":
		return node.NewProperty(nil, "DisableDefaultSignalHandler", "public", true, data.NewAnyValue(s.source.DisableDefaultSignalHandler)), true
	case "KeyBindings":
		return node.NewProperty(nil, "KeyBindings", "public", true, data.NewAnyValue(s.source.KeyBindings)), true
	case "OnShutdown":
		return node.NewProperty(nil, "OnShutdown", "public", true, data.NewAnyValue(s.source.OnShutdown)), true
	case "PostShutdown":
		return node.NewProperty(nil, "PostShutdown", "public", true, data.NewAnyValue(s.source.PostShutdown)), true
	case "ShouldQuit":
		return node.NewProperty(nil, "ShouldQuit", "public", true, data.NewAnyValue(s.source.ShouldQuit)), true
	case "RawMessageHandler":
		return node.NewProperty(nil, "RawMessageHandler", "public", true, data.NewAnyValue(s.source.RawMessageHandler)), true
	case "WarningHandler":
		return node.NewProperty(nil, "WarningHandler", "public", true, data.NewAnyValue(s.source.WarningHandler)), true
	case "ErrorHandler":
		return node.NewProperty(nil, "ErrorHandler", "public", true, data.NewAnyValue(s.source.ErrorHandler)), true
	case "FileAssociations":
		return node.NewProperty(nil, "FileAssociations", "public", true, data.NewAnyValue(s.source.FileAssociations)), true
	case "SingleInstance":
		return node.NewProperty(nil, "SingleInstance", "public", true, data.NewClassValue(NewSingleInstanceOptionsClassFrom(s.source.SingleInstance), runtime.NewContextToDo())), true
	}
	return nil, false
}

func (s *OptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Name":                        node.NewProperty(nil, "Name", "public", true, data.NewAnyValue(nil)),
		"Description":                 node.NewProperty(nil, "Description", "public", true, data.NewAnyValue(nil)),
		"Icon":                        node.NewProperty(nil, "Icon", "public", true, data.NewAnyValue(nil)),
		"Mac":                         node.NewProperty(nil, "Mac", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Windows":                     node.NewProperty(nil, "Windows", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Linux":                       node.NewProperty(nil, "Linux", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Services":                    node.NewProperty(nil, "Services", "public", true, data.NewAnyValue(nil)),
		"MarshalError":                node.NewProperty(nil, "MarshalError", "public", true, data.NewAnyValue(nil)),
		"BindAliases":                 node.NewProperty(nil, "BindAliases", "public", true, data.NewAnyValue(nil)),
		"Logger":                      node.NewProperty(nil, "Logger", "public", true, data.NewAnyValue(nil)),
		"LogLevel":                    node.NewProperty(nil, "LogLevel", "public", true, data.NewAnyValue(nil)),
		"Assets":                      node.NewProperty(nil, "Assets", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Flags":                       node.NewProperty(nil, "Flags", "public", true, data.NewAnyValue(nil)),
		"PanicHandler":                node.NewProperty(nil, "PanicHandler", "public", true, data.NewAnyValue(nil)),
		"DisableDefaultSignalHandler": node.NewProperty(nil, "DisableDefaultSignalHandler", "public", true, data.NewAnyValue(nil)),
		"KeyBindings":                 node.NewProperty(nil, "KeyBindings", "public", true, data.NewAnyValue(nil)),
		"OnShutdown":                  node.NewProperty(nil, "OnShutdown", "public", true, data.NewAnyValue(nil)),
		"PostShutdown":                node.NewProperty(nil, "PostShutdown", "public", true, data.NewAnyValue(nil)),
		"ShouldQuit":                  node.NewProperty(nil, "ShouldQuit", "public", true, data.NewAnyValue(nil)),
		"RawMessageHandler":           node.NewProperty(nil, "RawMessageHandler", "public", true, data.NewAnyValue(nil)),
		"WarningHandler":              node.NewProperty(nil, "WarningHandler", "public", true, data.NewAnyValue(nil)),
		"ErrorHandler":                node.NewProperty(nil, "ErrorHandler", "public", true, data.NewAnyValue(nil)),
		"FileAssociations":            node.NewProperty(nil, "FileAssociations", "public", true, data.NewAnyValue(nil)),
		"SingleInstance":              node.NewProperty(nil, "SingleInstance", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
	}
}

func (s *OptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Name":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Name = string(sv.AsString())
		}
	case "Description":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Description = string(sv.AsString())
		}
	case "Icon":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Icon = v.Value.([]uint8)
		}
	case "Mac":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MacOptions); ok {
					s.source.Mac = *ptr
				} else {
					s.source.Mac = src.(applicationsrc.MacOptions)
				}
			}
		case *data.AnyValue:
			s.source.Mac = v.Value.(applicationsrc.MacOptions)
		}
	case "Windows":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.WindowsOptions); ok {
					s.source.Windows = *ptr
				} else {
					s.source.Windows = src.(applicationsrc.WindowsOptions)
				}
			}
		case *data.AnyValue:
			s.source.Windows = v.Value.(applicationsrc.WindowsOptions)
		}
	case "Linux":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.LinuxOptions); ok {
					s.source.Linux = *ptr
				} else {
					s.source.Linux = src.(applicationsrc.LinuxOptions)
				}
			}
		case *data.AnyValue:
			s.source.Linux = v.Value.(applicationsrc.LinuxOptions)
		}
	case "Services":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Services = v.Value.([]applicationsrc.Service)
		}
	case "MarshalError":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.MarshalError = v.Value.(func(error) []uint8)
		}
	case "BindAliases":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.BindAliases = v.Value.(map[uint32]uint32)
		}
	case "Logger":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Logger = v.Value.(*slog.Logger)
		}
	case "LogLevel":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.LogLevel = v.Value.(slog.Level)
		}
	case "Assets":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.AssetOptions); ok {
					s.source.Assets = *ptr
				} else {
					s.source.Assets = src.(applicationsrc.AssetOptions)
				}
			}
		case *data.AnyValue:
			s.source.Assets = v.Value.(applicationsrc.AssetOptions)
		}
	case "Flags":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Flags = v.Value.(map[string]interface{})
		}
	case "PanicHandler":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.PanicHandler = v.Value.(func(*applicationsrc.PanicDetails))
		}
	case "DisableDefaultSignalHandler":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DisableDefaultSignalHandler = bool(x)
			}
		}
	case "KeyBindings":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.KeyBindings = v.Value.(map[string]func(applicationsrc.Window))
		}
	case "OnShutdown":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.OnShutdown = v.Value.(func())
		}
	case "PostShutdown":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.PostShutdown = v.Value.(func())
		}
	case "ShouldQuit":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.ShouldQuit = v.Value.(func() bool)
		}
	case "RawMessageHandler":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.RawMessageHandler = v.Value.(func(applicationsrc.Window, string))
		}
	case "WarningHandler":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.WarningHandler = v.Value.(func(string))
		}
	case "ErrorHandler":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.ErrorHandler = v.Value.(func(error))
		}
	case "FileAssociations":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.FileAssociations = v.Value.([]string)
		}
	case "SingleInstance":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.SingleInstanceOptions); ok {
					s.source.SingleInstance = ptr
				}
			}
		case *data.AnyValue:
			s.source.SingleInstance = v.Value.(*applicationsrc.SingleInstanceOptions)
		}
	}
}

func (s *OptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *OptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *OptionsClass) GetConstruct() data.Method { return s.construct }

type OptionsConstructMethod struct {
	source *applicationsrc.Options
}

func (h *OptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *OptionsConstructMethod) GetName() string               { return "construct" }
func (h *OptionsConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *OptionsConstructMethod) GetIsStatic() bool             { return false }
func (h *OptionsConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *OptionsConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *OptionsConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
