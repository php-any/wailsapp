package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMenuBarThemeClass() data.ClassStmt {
	return &MenuBarThemeClass{
		source:    nil,
		construct: &MenuBarThemeConstructMethod{source: nil},
	}
}

func NewMenuBarThemeClassFrom(source *applicationsrc.MenuBarTheme) data.ClassStmt {
	return &MenuBarThemeClass{
		source:    source,
		construct: &MenuBarThemeConstructMethod{source: source},
	}
}

type MenuBarThemeClass struct {
	node.Node
	source    *applicationsrc.MenuBarTheme
	construct data.Method
}

func (s *MenuBarThemeClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMenuBarThemeClassFrom(&applicationsrc.MenuBarTheme{}), ctx.CreateBaseContext()), nil
}
func (s *MenuBarThemeClass) GetName() string         { return "wails\\application\\MenuBarTheme" }
func (s *MenuBarThemeClass) GetExtend() *string      { return nil }
func (s *MenuBarThemeClass) GetImplements() []string { return nil }
func (s *MenuBarThemeClass) AsString() string        { return "MenuBarTheme{}" }
func (s *MenuBarThemeClass) GetSource() any          { return s.source }
func (s *MenuBarThemeClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Default":
		return node.NewProperty(nil, "Default", "public", true, data.NewClassValue(NewTextThemeClassFrom(s.source.Default), runtime.NewContextToDo())), true
	case "Hover":
		return node.NewProperty(nil, "Hover", "public", true, data.NewClassValue(NewTextThemeClassFrom(s.source.Hover), runtime.NewContextToDo())), true
	case "Selected":
		return node.NewProperty(nil, "Selected", "public", true, data.NewClassValue(NewTextThemeClassFrom(s.source.Selected), runtime.NewContextToDo())), true
	}
	return nil, false
}

func (s *MenuBarThemeClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Default":  node.NewProperty(nil, "Default", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Hover":    node.NewProperty(nil, "Hover", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Selected": node.NewProperty(nil, "Selected", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
	}
}

func (s *MenuBarThemeClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Default":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.TextTheme); ok {
					s.source.Default = ptr
				}
			}
		case *data.AnyValue:
			s.source.Default = v.Value.(*applicationsrc.TextTheme)
		}
	case "Hover":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.TextTheme); ok {
					s.source.Hover = ptr
				}
			}
		case *data.AnyValue:
			s.source.Hover = v.Value.(*applicationsrc.TextTheme)
		}
	case "Selected":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.TextTheme); ok {
					s.source.Selected = ptr
				}
			}
		case *data.AnyValue:
			s.source.Selected = v.Value.(*applicationsrc.TextTheme)
		}
	}
}

func (s *MenuBarThemeClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *MenuBarThemeClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *MenuBarThemeClass) GetConstruct() data.Method { return s.construct }

type MenuBarThemeConstructMethod struct {
	source *applicationsrc.MenuBarTheme
}

func (h *MenuBarThemeConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MenuBarThemeConstructMethod) GetName() string               { return "construct" }
func (h *MenuBarThemeConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MenuBarThemeConstructMethod) GetIsStatic() bool             { return false }
func (h *MenuBarThemeConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MenuBarThemeConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MenuBarThemeConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
