package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type DialogManagerWarningMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerWarningMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Warning()
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *DialogManagerWarningMethod) GetName() string            { return "warning" }
func (h *DialogManagerWarningMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *DialogManagerWarningMethod) GetIsStatic() bool          { return true }
func (h *DialogManagerWarningMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *DialogManagerWarningMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *DialogManagerWarningMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
