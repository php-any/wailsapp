package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type NewAppMenuFunction struct{}

func NewNewAppMenuFunction() data.FuncStmt {
	return &NewAppMenuFunction{}
}

func (h *NewAppMenuFunction) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := applicationsrc.NewAppMenu()
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *NewAppMenuFunction) GetName() string            { return "wails\\application\\newAppMenu" }
func (h *NewAppMenuFunction) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *NewAppMenuFunction) GetIsStatic() bool          { return true }
func (h *NewAppMenuFunction) GetParams() []data.GetValue {
	return []data.GetValue{}
}
func (h *NewAppMenuFunction) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *NewAppMenuFunction) GetReturnType() data.Types { return data.NewBaseType("void") }
