package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ButtonSetAsCancelMethod struct {
	source *applicationsrc.Button
}

func (h *ButtonSetAsCancelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.SetAsCancel()
	return data.NewClassValue(NewButtonClassFrom(ret0), ctx), nil
}

func (h *ButtonSetAsCancelMethod) GetName() string            { return "setAsCancel" }
func (h *ButtonSetAsCancelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ButtonSetAsCancelMethod) GetIsStatic() bool          { return true }
func (h *ButtonSetAsCancelMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ButtonSetAsCancelMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ButtonSetAsCancelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
