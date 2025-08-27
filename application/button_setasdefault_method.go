package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ButtonSetAsDefaultMethod struct {
	source *applicationsrc.Button
}

func (h *ButtonSetAsDefaultMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.SetAsDefault()
	return data.NewClassValue(NewButtonClassFrom(ret0), ctx), nil
}

func (h *ButtonSetAsDefaultMethod) GetName() string            { return "setAsDefault" }
func (h *ButtonSetAsDefaultMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ButtonSetAsDefaultMethod) GetIsStatic() bool          { return true }
func (h *ButtonSetAsDefaultMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ButtonSetAsDefaultMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ButtonSetAsDefaultMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
