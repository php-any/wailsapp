package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerNewMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerNewMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.New()
	return data.NewClassValue(NewWebviewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowManagerNewMethod) GetName() string            { return "new" }
func (h *WindowManagerNewMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerNewMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerNewMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowManagerNewMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowManagerNewMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
