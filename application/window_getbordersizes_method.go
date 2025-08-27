package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowGetBorderSizesMethod struct {
	source applicationsrc.Window
}

func (h *WindowGetBorderSizesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetBorderSizes()
	return data.NewClassValue(NewLRTBClassFrom(ret0), ctx), nil
}

func (h *WindowGetBorderSizesMethod) GetName() string            { return "getBorderSizes" }
func (h *WindowGetBorderSizesMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowGetBorderSizesMethod) GetIsStatic() bool          { return true }
func (h *WindowGetBorderSizesMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowGetBorderSizesMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowGetBorderSizesMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
