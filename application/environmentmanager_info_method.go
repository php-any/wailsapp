package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EnvironmentManagerInfoMethod struct {
	source *applicationsrc.EnvironmentManager
}

func (h *EnvironmentManagerInfoMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Info()
	return data.NewClassValue(NewEnvironmentInfoClassFrom(&ret0), ctx), nil
}

func (h *EnvironmentManagerInfoMethod) GetName() string            { return "info" }
func (h *EnvironmentManagerInfoMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *EnvironmentManagerInfoMethod) GetIsStatic() bool          { return true }
func (h *EnvironmentManagerInfoMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *EnvironmentManagerInfoMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *EnvironmentManagerInfoMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
