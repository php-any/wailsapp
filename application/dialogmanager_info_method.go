package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type DialogManagerInfoMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerInfoMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Info()
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *DialogManagerInfoMethod) GetName() string            { return "info" }
func (h *DialogManagerInfoMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *DialogManagerInfoMethod) GetIsStatic() bool          { return true }
func (h *DialogManagerInfoMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *DialogManagerInfoMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *DialogManagerInfoMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
