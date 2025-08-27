package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowUnMinimiseMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowUnMinimiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.UnMinimise()
	return nil, nil
}

func (h *WebviewWindowUnMinimiseMethod) GetName() string            { return "unMinimise" }
func (h *WebviewWindowUnMinimiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowUnMinimiseMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowUnMinimiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowUnMinimiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowUnMinimiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
