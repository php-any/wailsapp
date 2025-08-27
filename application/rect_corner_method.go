package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type RectCornerMethod struct {
	source *applicationsrc.Rect
}

func (h *RectCornerMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Corner()
	return data.NewClassValue(NewPointClassFrom(&ret0), ctx), nil
}

func (h *RectCornerMethod) GetName() string            { return "corner" }
func (h *RectCornerMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *RectCornerMethod) GetIsStatic() bool          { return true }
func (h *RectCornerMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *RectCornerMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *RectCornerMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
