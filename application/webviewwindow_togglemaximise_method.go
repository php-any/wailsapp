package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowToggleMaximiseMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowToggleMaximiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ToggleMaximise()
	return nil, nil
}

func (h *WebviewWindowToggleMaximiseMethod) GetName() string            { return "toggleMaximise" }
func (h *WebviewWindowToggleMaximiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowToggleMaximiseMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowToggleMaximiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowToggleMaximiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowToggleMaximiseMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
