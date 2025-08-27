package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowGetScreenMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowGetScreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, err := h.source.GetScreen()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return data.NewClassValue(NewScreenClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowGetScreenMethod) GetName() string            { return "getScreen" }
func (h *WebviewWindowGetScreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowGetScreenMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowGetScreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowGetScreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowGetScreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
