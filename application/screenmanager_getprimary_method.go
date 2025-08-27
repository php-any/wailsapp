package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ScreenManagerGetPrimaryMethod struct {
	source *applicationsrc.ScreenManager
}

func (h *ScreenManagerGetPrimaryMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetPrimary()
	return data.NewClassValue(NewScreenClassFrom(ret0), ctx), nil
}

func (h *ScreenManagerGetPrimaryMethod) GetName() string            { return "getPrimary" }
func (h *ScreenManagerGetPrimaryMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ScreenManagerGetPrimaryMethod) GetIsStatic() bool          { return true }
func (h *ScreenManagerGetPrimaryMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ScreenManagerGetPrimaryMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ScreenManagerGetPrimaryMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
