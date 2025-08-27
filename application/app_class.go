package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	slog "log/slog"
)

func NewAppClass() data.ClassStmt {
	return &AppClass{
		source:          nil,
		construct:       &AppConstructMethod{source: nil},
		capabilities:    &AppCapabilitiesMethod{source: nil},
		config:          &AppConfigMethod{source: nil},
		context:         &AppContextMethod{source: nil},
		getPID:          &AppGetPIDMethod{source: nil},
		hide:            &AppHideMethod{source: nil},
		newMenu:         &AppNewMenuMethod{source: nil},
		onShutdown:      &AppOnShutdownMethod{source: nil},
		quit:            &AppQuitMethod{source: nil},
		registerService: &AppRegisterServiceMethod{source: nil},
		run:             &AppRunMethod{source: nil},
		setIcon:         &AppSetIconMethod{source: nil},
		show:            &AppShowMethod{source: nil},
	}
}

func NewAppClassFrom(source *applicationsrc.App) data.ClassStmt {
	return &AppClass{
		source:          source,
		construct:       &AppConstructMethod{source: source},
		capabilities:    &AppCapabilitiesMethod{source: source},
		config:          &AppConfigMethod{source: source},
		context:         &AppContextMethod{source: source},
		getPID:          &AppGetPIDMethod{source: source},
		hide:            &AppHideMethod{source: source},
		newMenu:         &AppNewMenuMethod{source: source},
		onShutdown:      &AppOnShutdownMethod{source: source},
		quit:            &AppQuitMethod{source: source},
		registerService: &AppRegisterServiceMethod{source: source},
		run:             &AppRunMethod{source: source},
		setIcon:         &AppSetIconMethod{source: source},
		show:            &AppShowMethod{source: source},
	}
}

type AppClass struct {
	node.Node
	source          *applicationsrc.App
	construct       data.Method
	capabilities    data.Method
	config          data.Method
	context         data.Method
	getPID          data.Method
	hide            data.Method
	newMenu         data.Method
	onShutdown      data.Method
	quit            data.Method
	registerService data.Method
	run             data.Method
	setIcon         data.Method
	show            data.Method
}

func (s *AppClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewAppClassFrom(&applicationsrc.App{}), ctx.CreateBaseContext()), nil
}
func (s *AppClass) GetName() string         { return "wails\\application\\App" }
func (s *AppClass) GetExtend() *string      { return nil }
func (s *AppClass) GetImplements() []string { return nil }
func (s *AppClass) AsString() string        { return "App{}" }
func (s *AppClass) GetSource() any          { return s.source }
func (s *AppClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Window":
		return node.NewProperty(nil, "Window", "public", true, data.NewClassValue(NewWindowManagerClassFrom(s.source.Window), runtime.NewContextToDo())), true
	case "ContextMenu":
		return node.NewProperty(nil, "ContextMenu", "public", true, data.NewClassValue(NewContextMenuManagerClassFrom(s.source.ContextMenu), runtime.NewContextToDo())), true
	case "KeyBinding":
		return node.NewProperty(nil, "KeyBinding", "public", true, data.NewClassValue(NewKeyBindingManagerClassFrom(s.source.KeyBinding), runtime.NewContextToDo())), true
	case "Browser":
		return node.NewProperty(nil, "Browser", "public", true, data.NewClassValue(NewBrowserManagerClassFrom(s.source.Browser), runtime.NewContextToDo())), true
	case "Env":
		return node.NewProperty(nil, "Env", "public", true, data.NewClassValue(NewEnvironmentManagerClassFrom(s.source.Env), runtime.NewContextToDo())), true
	case "Dialog":
		return node.NewProperty(nil, "Dialog", "public", true, data.NewClassValue(NewDialogManagerClassFrom(s.source.Dialog), runtime.NewContextToDo())), true
	case "Event":
		return node.NewProperty(nil, "Event", "public", true, data.NewClassValue(NewEventManagerClassFrom(s.source.Event), runtime.NewContextToDo())), true
	case "Menu":
		return node.NewProperty(nil, "Menu", "public", true, data.NewClassValue(NewMenuManagerClassFrom(s.source.Menu), runtime.NewContextToDo())), true
	case "Screen":
		return node.NewProperty(nil, "Screen", "public", true, data.NewClassValue(NewScreenManagerClassFrom(s.source.Screen), runtime.NewContextToDo())), true
	case "Clipboard":
		return node.NewProperty(nil, "Clipboard", "public", true, data.NewClassValue(NewClipboardManagerClassFrom(s.source.Clipboard), runtime.NewContextToDo())), true
	case "SystemTray":
		return node.NewProperty(nil, "SystemTray", "public", true, data.NewClassValue(NewSystemTrayManagerClassFrom(s.source.SystemTray), runtime.NewContextToDo())), true
	case "Logger":
		return node.NewProperty(nil, "Logger", "public", true, data.NewAnyValue(s.source.Logger)), true
	}
	return nil, false
}

func (s *AppClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Window":      node.NewProperty(nil, "Window", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"ContextMenu": node.NewProperty(nil, "ContextMenu", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"KeyBinding":  node.NewProperty(nil, "KeyBinding", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Browser":     node.NewProperty(nil, "Browser", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Env":         node.NewProperty(nil, "Env", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Dialog":      node.NewProperty(nil, "Dialog", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Event":       node.NewProperty(nil, "Event", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Menu":        node.NewProperty(nil, "Menu", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Screen":      node.NewProperty(nil, "Screen", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Clipboard":   node.NewProperty(nil, "Clipboard", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"SystemTray":  node.NewProperty(nil, "SystemTray", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Logger":      node.NewProperty(nil, "Logger", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *AppClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Window":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.WindowManager); ok {
					s.source.Window = ptr
				}
			}
		case *data.AnyValue:
			s.source.Window = v.Value.(*applicationsrc.WindowManager)
		}
	case "ContextMenu":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.ContextMenuManager); ok {
					s.source.ContextMenu = ptr
				}
			}
		case *data.AnyValue:
			s.source.ContextMenu = v.Value.(*applicationsrc.ContextMenuManager)
		}
	case "KeyBinding":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.KeyBindingManager); ok {
					s.source.KeyBinding = ptr
				}
			}
		case *data.AnyValue:
			s.source.KeyBinding = v.Value.(*applicationsrc.KeyBindingManager)
		}
	case "Browser":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.BrowserManager); ok {
					s.source.Browser = ptr
				}
			}
		case *data.AnyValue:
			s.source.Browser = v.Value.(*applicationsrc.BrowserManager)
		}
	case "Env":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.EnvironmentManager); ok {
					s.source.Env = ptr
				}
			}
		case *data.AnyValue:
			s.source.Env = v.Value.(*applicationsrc.EnvironmentManager)
		}
	case "Dialog":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.DialogManager); ok {
					s.source.Dialog = ptr
				}
			}
		case *data.AnyValue:
			s.source.Dialog = v.Value.(*applicationsrc.DialogManager)
		}
	case "Event":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.EventManager); ok {
					s.source.Event = ptr
				}
			}
		case *data.AnyValue:
			s.source.Event = v.Value.(*applicationsrc.EventManager)
		}
	case "Menu":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MenuManager); ok {
					s.source.Menu = ptr
				}
			}
		case *data.AnyValue:
			s.source.Menu = v.Value.(*applicationsrc.MenuManager)
		}
	case "Screen":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.ScreenManager); ok {
					s.source.Screen = ptr
				}
			}
		case *data.AnyValue:
			s.source.Screen = v.Value.(*applicationsrc.ScreenManager)
		}
	case "Clipboard":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.ClipboardManager); ok {
					s.source.Clipboard = ptr
				}
			}
		case *data.AnyValue:
			s.source.Clipboard = v.Value.(*applicationsrc.ClipboardManager)
		}
	case "SystemTray":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.SystemTrayManager); ok {
					s.source.SystemTray = ptr
				}
			}
		case *data.AnyValue:
			s.source.SystemTray = v.Value.(*applicationsrc.SystemTrayManager)
		}
	case "Logger":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Logger = v.Value.(*slog.Logger)
		}
	}
}

func (s *AppClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "capabilities":
		return s.capabilities, true
	case "config":
		return s.config, true
	case "context":
		return s.context, true
	case "getPID":
		return s.getPID, true
	case "hide":
		return s.hide, true
	case "newMenu":
		return s.newMenu, true
	case "onShutdown":
		return s.onShutdown, true
	case "quit":
		return s.quit, true
	case "registerService":
		return s.registerService, true
	case "run":
		return s.run, true
	case "setIcon":
		return s.setIcon, true
	case "show":
		return s.show, true
	}
	return nil, false
}

func (s *AppClass) GetMethods() []data.Method {
	return []data.Method{
		s.capabilities,
		s.config,
		s.context,
		s.getPID,
		s.hide,
		s.newMenu,
		s.onShutdown,
		s.quit,
		s.registerService,
		s.run,
		s.setIcon,
		s.show,
	}
}

func (s *AppClass) GetConstruct() data.Method { return s.construct }

type AppConstructMethod struct {
	source *applicationsrc.App
}

func (h *AppConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *AppConstructMethod) GetName() string               { return "construct" }
func (h *AppConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *AppConstructMethod) GetIsStatic() bool             { return false }
func (h *AppConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *AppConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *AppConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
