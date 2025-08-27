package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ScreenOriginMethod struct {
	source *applicationsrc.Screen
}

func (h *ScreenOriginMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Origin()
	return data.NewClassValue(NewPointClassFrom(&ret0), ctx), nil
}

func (h *ScreenOriginMethod) GetName() string            { return "origin" }
func (h *ScreenOriginMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ScreenOriginMethod) GetIsStatic() bool          { return true }
func (h *ScreenOriginMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ScreenOriginMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ScreenOriginMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
