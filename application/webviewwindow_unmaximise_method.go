package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowUnMaximiseMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowUnMaximiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.UnMaximise()
	return nil, nil
}

func (h *WebviewWindowUnMaximiseMethod) GetName() string            { return "unMaximise" }
func (h *WebviewWindowUnMaximiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowUnMaximiseMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowUnMaximiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowUnMaximiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowUnMaximiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
