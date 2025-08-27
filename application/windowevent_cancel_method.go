package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowEventCancelMethod struct {
	source *applicationsrc.WindowEvent
}

func (h *WindowEventCancelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Cancel()
	return nil, nil
}

func (h *WindowEventCancelMethod) GetName() string            { return "cancel" }
func (h *WindowEventCancelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowEventCancelMethod) GetIsStatic() bool          { return true }
func (h *WindowEventCancelMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowEventCancelMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowEventCancelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
