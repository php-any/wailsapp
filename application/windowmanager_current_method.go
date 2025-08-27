package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerCurrentMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerCurrentMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Current()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowManagerCurrentMethod) GetName() string            { return "current" }
func (h *WindowManagerCurrentMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerCurrentMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerCurrentMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowManagerCurrentMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowManagerCurrentMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
