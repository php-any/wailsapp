package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowCenterMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowCenterMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Center()
	return nil, nil
}

func (h *WebviewWindowCenterMethod) GetName() string            { return "center" }
func (h *WebviewWindowCenterMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowCenterMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowCenterMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowCenterMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowCenterMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
