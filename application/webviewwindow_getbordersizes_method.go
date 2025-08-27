package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowGetBorderSizesMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowGetBorderSizesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetBorderSizes()
	return data.NewClassValue(NewLRTBClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowGetBorderSizesMethod) GetName() string            { return "getBorderSizes" }
func (h *WebviewWindowGetBorderSizesMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowGetBorderSizesMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowGetBorderSizesMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowGetBorderSizesMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowGetBorderSizesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
