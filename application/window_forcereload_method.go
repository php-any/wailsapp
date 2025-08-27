package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowForceReloadMethod struct {
	source applicationsrc.Window
}

func (h *WindowForceReloadMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ForceReload()
	return nil, nil
}

func (h *WindowForceReloadMethod) GetName() string            { return "forceReload" }
func (h *WindowForceReloadMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowForceReloadMethod) GetIsStatic() bool          { return true }
func (h *WindowForceReloadMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowForceReloadMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowForceReloadMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
