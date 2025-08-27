package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type RectOriginMethod struct {
	source *applicationsrc.Rect
}

func (h *RectOriginMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Origin()
	return data.NewClassValue(NewPointClassFrom(&ret0), ctx), nil
}

func (h *RectOriginMethod) GetName() string            { return "origin" }
func (h *RectOriginMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *RectOriginMethod) GetIsStatic() bool          { return true }
func (h *RectOriginMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *RectOriginMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *RectOriginMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
