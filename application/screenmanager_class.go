package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewScreenManagerClass() data.ClassStmt {
	return &ScreenManagerClass{
		source:                     nil,
		construct:                  &ScreenManagerConstructMethod{source: nil},
		dipToPhysicalPoint:         &ScreenManagerDipToPhysicalPointMethod{source: nil},
		dipToPhysicalRect:          &ScreenManagerDipToPhysicalRectMethod{source: nil},
		getAll:                     &ScreenManagerGetAllMethod{source: nil},
		getPrimary:                 &ScreenManagerGetPrimaryMethod{source: nil},
		layoutScreens:              &ScreenManagerLayoutScreensMethod{source: nil},
		physicalToDipPoint:         &ScreenManagerPhysicalToDipPointMethod{source: nil},
		physicalToDipRect:          &ScreenManagerPhysicalToDipRectMethod{source: nil},
		screenNearestDipPoint:      &ScreenManagerScreenNearestDipPointMethod{source: nil},
		screenNearestDipRect:       &ScreenManagerScreenNearestDipRectMethod{source: nil},
		screenNearestPhysicalPoint: &ScreenManagerScreenNearestPhysicalPointMethod{source: nil},
		screenNearestPhysicalRect:  &ScreenManagerScreenNearestPhysicalRectMethod{source: nil},
	}
}

func NewScreenManagerClassFrom(source *applicationsrc.ScreenManager) data.ClassStmt {
	return &ScreenManagerClass{
		source:                     source,
		construct:                  &ScreenManagerConstructMethod{source: source},
		dipToPhysicalPoint:         &ScreenManagerDipToPhysicalPointMethod{source: source},
		dipToPhysicalRect:          &ScreenManagerDipToPhysicalRectMethod{source: source},
		getAll:                     &ScreenManagerGetAllMethod{source: source},
		getPrimary:                 &ScreenManagerGetPrimaryMethod{source: source},
		layoutScreens:              &ScreenManagerLayoutScreensMethod{source: source},
		physicalToDipPoint:         &ScreenManagerPhysicalToDipPointMethod{source: source},
		physicalToDipRect:          &ScreenManagerPhysicalToDipRectMethod{source: source},
		screenNearestDipPoint:      &ScreenManagerScreenNearestDipPointMethod{source: source},
		screenNearestDipRect:       &ScreenManagerScreenNearestDipRectMethod{source: source},
		screenNearestPhysicalPoint: &ScreenManagerScreenNearestPhysicalPointMethod{source: source},
		screenNearestPhysicalRect:  &ScreenManagerScreenNearestPhysicalRectMethod{source: source},
	}
}

type ScreenManagerClass struct {
	node.Node
	source                     *applicationsrc.ScreenManager
	construct                  data.Method
	dipToPhysicalPoint         data.Method
	dipToPhysicalRect          data.Method
	getAll                     data.Method
	getPrimary                 data.Method
	layoutScreens              data.Method
	physicalToDipPoint         data.Method
	physicalToDipRect          data.Method
	screenNearestDipPoint      data.Method
	screenNearestDipRect       data.Method
	screenNearestPhysicalPoint data.Method
	screenNearestPhysicalRect  data.Method
}

func (s *ScreenManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewScreenManagerClassFrom(&applicationsrc.ScreenManager{}), ctx.CreateBaseContext()), nil
}
func (s *ScreenManagerClass) GetName() string                            { return "wails\\application\\ScreenManager" }
func (s *ScreenManagerClass) GetExtend() *string                         { return nil }
func (s *ScreenManagerClass) GetImplements() []string                    { return nil }
func (s *ScreenManagerClass) AsString() string                           { return "ScreenManager{}" }
func (s *ScreenManagerClass) GetSource() any                             { return s.source }
func (s *ScreenManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *ScreenManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *ScreenManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "dipToPhysicalPoint":
		return s.dipToPhysicalPoint, true
	case "dipToPhysicalRect":
		return s.dipToPhysicalRect, true
	case "getAll":
		return s.getAll, true
	case "getPrimary":
		return s.getPrimary, true
	case "layoutScreens":
		return s.layoutScreens, true
	case "physicalToDipPoint":
		return s.physicalToDipPoint, true
	case "physicalToDipRect":
		return s.physicalToDipRect, true
	case "screenNearestDipPoint":
		return s.screenNearestDipPoint, true
	case "screenNearestDipRect":
		return s.screenNearestDipRect, true
	case "screenNearestPhysicalPoint":
		return s.screenNearestPhysicalPoint, true
	case "screenNearestPhysicalRect":
		return s.screenNearestPhysicalRect, true
	}
	return nil, false
}

func (s *ScreenManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.dipToPhysicalPoint,
		s.dipToPhysicalRect,
		s.getAll,
		s.getPrimary,
		s.layoutScreens,
		s.physicalToDipPoint,
		s.physicalToDipRect,
		s.screenNearestDipPoint,
		s.screenNearestDipRect,
		s.screenNearestPhysicalPoint,
		s.screenNearestPhysicalRect,
	}
}

func (s *ScreenManagerClass) GetConstruct() data.Method { return s.construct }

type ScreenManagerConstructMethod struct {
	source *applicationsrc.ScreenManager
}

func (h *ScreenManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *ScreenManagerConstructMethod) GetName() string               { return "construct" }
func (h *ScreenManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *ScreenManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *ScreenManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *ScreenManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *ScreenManagerConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
