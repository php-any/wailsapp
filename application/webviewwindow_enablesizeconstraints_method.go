package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowEnableSizeConstraintsMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowEnableSizeConstraintsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.EnableSizeConstraints()
	return nil, nil
}

func (h *WebviewWindowEnableSizeConstraintsMethod) GetName() string { return "enableSizeConstraints" }
func (h *WebviewWindowEnableSizeConstraintsMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowEnableSizeConstraintsMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowEnableSizeConstraintsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowEnableSizeConstraintsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowEnableSizeConstraintsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
