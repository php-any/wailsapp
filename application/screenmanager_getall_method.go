package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ScreenManagerGetAllMethod struct {
	source *applicationsrc.ScreenManager
}

func (h *ScreenManagerGetAllMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetAll()
	return data.NewAnyValue(ret0), nil
}

func (h *ScreenManagerGetAllMethod) GetName() string            { return "getAll" }
func (h *ScreenManagerGetAllMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ScreenManagerGetAllMethod) GetIsStatic() bool          { return true }
func (h *ScreenManagerGetAllMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ScreenManagerGetAllMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ScreenManagerGetAllMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
