package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MessageDialogShowMethod struct {
	source *applicationsrc.MessageDialog
}

func (h *MessageDialogShowMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Show()
	return nil, nil
}

func (h *MessageDialogShowMethod) GetName() string            { return "show" }
func (h *MessageDialogShowMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MessageDialogShowMethod) GetIsStatic() bool          { return true }
func (h *MessageDialogShowMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MessageDialogShowMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MessageDialogShowMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
