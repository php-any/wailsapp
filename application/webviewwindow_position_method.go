package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowPositionMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowPositionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, ret1 := h.source.Position()
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WebviewWindowPositionMethod) GetName() string            { return "position" }
func (h *WebviewWindowPositionMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowPositionMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowPositionMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowPositionMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowPositionMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
