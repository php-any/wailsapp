package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EnvironmentManagerIsDarkModeMethod struct {
	source *applicationsrc.EnvironmentManager
}

func (h *EnvironmentManagerIsDarkModeMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsDarkMode()
	return data.NewBoolValue(ret0), nil
}

func (h *EnvironmentManagerIsDarkModeMethod) GetName() string            { return "isDarkMode" }
func (h *EnvironmentManagerIsDarkModeMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *EnvironmentManagerIsDarkModeMethod) GetIsStatic() bool          { return true }
func (h *EnvironmentManagerIsDarkModeMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *EnvironmentManagerIsDarkModeMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *EnvironmentManagerIsDarkModeMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
