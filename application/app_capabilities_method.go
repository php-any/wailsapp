package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppCapabilitiesMethod struct {
	source *applicationsrc.App
}

func (h *AppCapabilitiesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Capabilities()
	return data.NewAnyValue(ret0), nil
}

func (h *AppCapabilitiesMethod) GetName() string            { return "capabilities" }
func (h *AppCapabilitiesMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppCapabilitiesMethod) GetIsStatic() bool          { return true }
func (h *AppCapabilitiesMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppCapabilitiesMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppCapabilitiesMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
