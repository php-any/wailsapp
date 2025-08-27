package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowCenterMethod struct {
	source applicationsrc.Window
}

func (h *WindowCenterMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Center()
	return nil, nil
}

func (h *WindowCenterMethod) GetName() string            { return "center" }
func (h *WindowCenterMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowCenterMethod) GetIsStatic() bool          { return true }
func (h *WindowCenterMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowCenterMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowCenterMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
