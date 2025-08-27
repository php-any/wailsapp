package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowIsIgnoreMouseEventsMethod struct {
	source applicationsrc.Window
}

func (h *WindowIsIgnoreMouseEventsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsIgnoreMouseEvents()
	return data.NewBoolValue(ret0), nil
}

func (h *WindowIsIgnoreMouseEventsMethod) GetName() string            { return "isIgnoreMouseEvents" }
func (h *WindowIsIgnoreMouseEventsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowIsIgnoreMouseEventsMethod) GetIsStatic() bool          { return true }
func (h *WindowIsIgnoreMouseEventsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowIsIgnoreMouseEventsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowIsIgnoreMouseEventsMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
