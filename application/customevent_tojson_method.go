package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type CustomEventToJSONMethod struct {
	source *applicationsrc.CustomEvent
}

func (h *CustomEventToJSONMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.ToJSON()
	return data.NewStringValue(ret0), nil
}

func (h *CustomEventToJSONMethod) GetName() string            { return "toJSON" }
func (h *CustomEventToJSONMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *CustomEventToJSONMethod) GetIsStatic() bool          { return true }
func (h *CustomEventToJSONMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *CustomEventToJSONMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *CustomEventToJSONMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
