package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EnvironmentManagerGetAccentColorMethod struct {
	source *applicationsrc.EnvironmentManager
}

func (h *EnvironmentManagerGetAccentColorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetAccentColor()
	return data.NewStringValue(ret0), nil
}

func (h *EnvironmentManagerGetAccentColorMethod) GetName() string { return "getAccentColor" }
func (h *EnvironmentManagerGetAccentColorMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *EnvironmentManagerGetAccentColorMethod) GetIsStatic() bool { return true }
func (h *EnvironmentManagerGetAccentColorMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *EnvironmentManagerGetAccentColorMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *EnvironmentManagerGetAccentColorMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
