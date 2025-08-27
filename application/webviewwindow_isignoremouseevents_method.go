package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowIsIgnoreMouseEventsMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowIsIgnoreMouseEventsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsIgnoreMouseEvents()
	return data.NewBoolValue(ret0), nil
}

func (h *WebviewWindowIsIgnoreMouseEventsMethod) GetName() string { return "isIgnoreMouseEvents" }
func (h *WebviewWindowIsIgnoreMouseEventsMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowIsIgnoreMouseEventsMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowIsIgnoreMouseEventsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowIsIgnoreMouseEventsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowIsIgnoreMouseEventsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
