package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowHeightMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowHeightMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Height()
	return data.NewIntValue(ret0), nil
}

func (h *WebviewWindowHeightMethod) GetName() string            { return "height" }
func (h *WebviewWindowHeightMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowHeightMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowHeightMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowHeightMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowHeightMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
