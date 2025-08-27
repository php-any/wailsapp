package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type RectSizeMethod struct {
	source *applicationsrc.Rect
}

func (h *RectSizeMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Size()
	return data.NewClassValue(NewSizeClassFrom(&ret0), ctx), nil
}

func (h *RectSizeMethod) GetName() string            { return "size" }
func (h *RectSizeMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *RectSizeMethod) GetIsStatic() bool          { return true }
func (h *RectSizeMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *RectSizeMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *RectSizeMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
