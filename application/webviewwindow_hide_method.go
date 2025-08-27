package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowHideMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowHideMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Hide()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowHideMethod) GetName() string            { return "hide" }
func (h *WebviewWindowHideMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowHideMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowHideMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowHideMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowHideMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
