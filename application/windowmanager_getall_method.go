package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerGetAllMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerGetAllMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetAll()
	return data.NewAnyValue(ret0), nil
}

func (h *WindowManagerGetAllMethod) GetName() string            { return "getAll" }
func (h *WindowManagerGetAllMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerGetAllMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerGetAllMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowManagerGetAllMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowManagerGetAllMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
