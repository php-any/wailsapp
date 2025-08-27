package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowDisableSizeConstraintsMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowDisableSizeConstraintsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.DisableSizeConstraints()
	return nil, nil
}

func (h *WebviewWindowDisableSizeConstraintsMethod) GetName() string { return "disableSizeConstraints" }
func (h *WebviewWindowDisableSizeConstraintsMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowDisableSizeConstraintsMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowDisableSizeConstraintsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowDisableSizeConstraintsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowDisableSizeConstraintsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
