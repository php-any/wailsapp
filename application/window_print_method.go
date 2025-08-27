package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowPrintMethod struct {
	source applicationsrc.Window
}

func (h *WindowPrintMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	if err := h.source.Print(); err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return nil, nil
}

func (h *WindowPrintMethod) GetName() string            { return "print" }
func (h *WindowPrintMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowPrintMethod) GetIsStatic() bool          { return true }
func (h *WindowPrintMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowPrintMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowPrintMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
