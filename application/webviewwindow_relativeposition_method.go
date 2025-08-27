package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowRelativePositionMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowRelativePositionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, ret1 := h.source.RelativePosition()
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WebviewWindowRelativePositionMethod) GetName() string            { return "relativePosition" }
func (h *WebviewWindowRelativePositionMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowRelativePositionMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowRelativePositionMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowRelativePositionMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowRelativePositionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
