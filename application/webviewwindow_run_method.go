package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowRunMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowRunMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Run()
	return nil, nil
}

func (h *WebviewWindowRunMethod) GetName() string            { return "run" }
func (h *WebviewWindowRunMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowRunMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowRunMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowRunMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowRunMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
