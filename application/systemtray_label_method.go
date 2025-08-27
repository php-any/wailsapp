package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayLabelMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayLabelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Label()
	return data.NewStringValue(ret0), nil
}

func (h *SystemTrayLabelMethod) GetName() string            { return "label" }
func (h *SystemTrayLabelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayLabelMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayLabelMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *SystemTrayLabelMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *SystemTrayLabelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
