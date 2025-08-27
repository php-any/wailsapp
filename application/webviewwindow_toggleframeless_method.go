package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowToggleFramelessMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowToggleFramelessMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ToggleFrameless()
	return nil, nil
}

func (h *WebviewWindowToggleFramelessMethod) GetName() string            { return "toggleFrameless" }
func (h *WebviewWindowToggleFramelessMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowToggleFramelessMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowToggleFramelessMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowToggleFramelessMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowToggleFramelessMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
