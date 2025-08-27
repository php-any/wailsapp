package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowOpenDevToolsMethod struct {
	source applicationsrc.Window
}

func (h *WindowOpenDevToolsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.OpenDevTools()
	return nil, nil
}

func (h *WindowOpenDevToolsMethod) GetName() string            { return "openDevTools" }
func (h *WindowOpenDevToolsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowOpenDevToolsMethod) GetIsStatic() bool          { return true }
func (h *WindowOpenDevToolsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowOpenDevToolsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowOpenDevToolsMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
