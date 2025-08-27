package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type NewMenuFunction struct{}

func NewNewMenuFunction() data.FuncStmt {
	return &NewMenuFunction{}
}

func (h *NewMenuFunction) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := applicationsrc.NewMenu()
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *NewMenuFunction) GetName() string            { return "wails\\application\\newMenu" }
func (h *NewMenuFunction) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *NewMenuFunction) GetIsStatic() bool          { return true }
func (h *NewMenuFunction) GetParams() []data.GetValue {
	return []data.GetValue{}
}
func (h *NewMenuFunction) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *NewMenuFunction) GetReturnType() data.Types { return data.NewBaseType("void") }
