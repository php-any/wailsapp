package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowPrintMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowPrintMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	if err := h.source.Print(); err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return nil, nil
}

func (h *WebviewWindowPrintMethod) GetName() string            { return "print" }
func (h *WebviewWindowPrintMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowPrintMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowPrintMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowPrintMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowPrintMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
