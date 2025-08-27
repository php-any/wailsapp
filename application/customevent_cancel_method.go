package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type CustomEventCancelMethod struct {
	source *applicationsrc.CustomEvent
}

func (h *CustomEventCancelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Cancel()
	return nil, nil
}

func (h *CustomEventCancelMethod) GetName() string            { return "cancel" }
func (h *CustomEventCancelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *CustomEventCancelMethod) GetIsStatic() bool          { return true }
func (h *CustomEventCancelMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *CustomEventCancelMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *CustomEventCancelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
