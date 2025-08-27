package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowHideMethod struct {
	source applicationsrc.Window
}

func (h *WindowHideMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Hide()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowHideMethod) GetName() string            { return "hide" }
func (h *WindowHideMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowHideMethod) GetIsStatic() bool          { return true }
func (h *WindowHideMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowHideMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowHideMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
