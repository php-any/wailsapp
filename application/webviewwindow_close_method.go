package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowCloseMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowCloseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Close()
	return nil, nil
}

func (h *WebviewWindowCloseMethod) GetName() string            { return "close" }
func (h *WebviewWindowCloseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowCloseMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowCloseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowCloseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowCloseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
