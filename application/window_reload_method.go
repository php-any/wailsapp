package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowReloadMethod struct {
	source applicationsrc.Window
}

func (h *WindowReloadMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Reload()
	return nil, nil
}

func (h *WindowReloadMethod) GetName() string            { return "reload" }
func (h *WindowReloadMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowReloadMethod) GetIsStatic() bool          { return true }
func (h *WindowReloadMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowReloadMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowReloadMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
