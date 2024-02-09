//go:build (linux && cgo) || (darwin && cgo) || (freebsd && cgo)

package dlbindings

import (
	"errors"
	"plugin"
)

var dinolangReady bool = false

var AddClass func(name string, usedByDefault bool, caller func(args []string, segmentName string) bool, loader func() bool) bool
var RemoveClass func(name string) bool
var CheckOnVariable func(str string) bool
var GetType func(str string) string
var GetTypeOfVar func(str string) string
var GetTypeEx func(str string) string
var IfVariableReplaceIt func(str string) interface{}
var GetVariableValue func(v string) interface{}
var SetVariable func(name string, value interface{}) bool
var SetReturned func(varType string, value interface{}, segmentName string)
var StringToText func(str string) string
var TextToString func(str string) string
var PrintError func(str string)
var RunCode func(code string) bool
var ParseFile func(path string) bool
var Execute func(code, segmentName string, line int, isInIfElse bool) bool
var PiniginShell func()

func InitDinolang(dlPath string) error {
	if dinolangReady {
		return errors.New("DinoLang already initialized")
	}

	plugin, err := plugin.Open(dlPath)
	if err != nil {
		return err
	}

	piniginShell, err := plugin.Lookup("PiniginShell")
	if err != nil {
		return err
	}
	PiniginShell = piniginShell.(func())

	setReturned, err := plugin.Lookup("SetReturned")
	if err != nil {
		return err
	}
	SetReturned = setReturned.(func(varType string, value interface{}, segmentName string))

	addClass, err := plugin.Lookup("AddClass")
	if err != nil {
		return err
	}
	AddClass = addClass.(func(name string, usedByDefault bool, caller func(args []string, segmentName string) bool, loader func() bool) bool)

	removeClass, err := plugin.Lookup("RemoveClass")
	if err != nil {
		return err
	}
	RemoveClass = removeClass.(func(name string) bool)

	checkOnVariable, err := plugin.Lookup("CheckOnVariable")
	if err != nil {
		return err
	}
	CheckOnVariable = checkOnVariable.(func(str string) bool)

	getType, err := plugin.Lookup("GetType")
	if err != nil {
		return err
	}
	GetType = getType.(func(str string) string)

	getTypeOfVar, err := plugin.Lookup("GetTypeOfVar")
	if err != nil {
		return err
	}
	GetTypeOfVar = getTypeOfVar.(func(str string) string)

	getTypeEx, err := plugin.Lookup("GetTypeEx")
	if err != nil {
		return err
	}
	GetTypeEx = getTypeEx.(func(str string) string)

	ifVariableReplaceIt, err := plugin.Lookup("IfVariableReplaceIt")
	if err != nil {
		return err
	}
	IfVariableReplaceIt = ifVariableReplaceIt.(func(str string) interface{})

	stringToText, err := plugin.Lookup("StringToText")
	if err != nil {
		return err
	}
	StringToText = stringToText.(func(str string) string)

	textToString, err := plugin.Lookup("TextToString")
	if err != nil {
		return err
	}
	TextToString = textToString.(func(str string) string)

	getVariableValue, err := plugin.Lookup("GetVariableValue")
	if err != nil {
		return err
	}
	GetVariableValue = getVariableValue.(func(v string) interface{})

	setVariable, err := plugin.Lookup("SetVariable")
	if err != nil {
		return err
	}
	SetVariable = setVariable.(func(name string, value interface{}) bool)

	printError, err := plugin.Lookup("PrintError")
	if err != nil {
		return err
	}
	PrintError = printError.(func(str string))

	runCode, err := plugin.Lookup("RunCode")
	if err != nil {
		return err
	}
	RunCode = runCode.(func(code string) bool)

	parseFile, err := plugin.Lookup("ParseFile")
	if err != nil {
		return err
	}
	ParseFile = parseFile.(func(path string) bool)

	execute, err := plugin.Lookup("Execute")
	if err != nil {
		return err
	}
	Execute = execute.(func(code string, segmentName string, line int, isInIfElse bool) bool)

	return nil
}
