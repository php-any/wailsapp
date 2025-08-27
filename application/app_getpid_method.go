package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppGetPIDMethod struct {
	source *applicationsrc.App
}

func (h *AppGetPIDMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetPID()
	return data.NewIntValue(ret0), nil
}

func (h *AppGetPIDMethod) GetName() string            { return "getPID" }
func (h *AppGetPIDMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppGetPIDMethod) GetIsStatic() bool          { return true }
func (h *AppGetPIDMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppGetPIDMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppGetPIDMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
