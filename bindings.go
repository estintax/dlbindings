//go:build (linux && cgo) || (darwin && cgo) || (freebsd && cgo)

package dlbindings

import (
	"errors"
	"plugin"
)

var dinolangReady bool = false

var AddClass func(name string, usedByDefault bool, caller func(args []string, segmentName string) bool, loader func() bool) bool
var RemoveClass func(name string) bool
var SetClassUsage func(className string, inUse bool, isGlobal bool) bool
var CheckOnVariable func(str string) bool
var CleanUp func(cleanClasses bool) bool
var GetType func(str string) string
var GetTypeOfVar func(str string) string
var GetTypeEx func(str string) string
var IfVariableReplaceIt func(str string) interface{}
var GetVariableValue func(v string) interface{}
var SetVariable func(name string, value interface{}) bool
var SetReturned func(varType string, value interface{}, segmentName string)
var StringToText func(str string) string
var TextToString func(str string) string
var PrintString func(str string, newline bool)
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

	setReturned, err := plugin.Lookup("SetReturned")
	if err != nil {
		return err
	}

	addClass, err := plugin.Lookup("AddClass")
	if err != nil {
		return err
	}

	removeClass, err := plugin.Lookup("RemoveClass")
	if err != nil {
		return err
	}

	checkOnVariable, err := plugin.Lookup("CheckOnVariable")
	if err != nil {
		return err
	}

	cleanUp, err := plugin.Lookup("CleanUp")
	if err != nil {
		return err
	}

	getType, err := plugin.Lookup("GetType")
	if err != nil {
		return err
	}

	getTypeOfVar, err := plugin.Lookup("GetTypeOfVar")
	if err != nil {
		return err
	}

	getTypeEx, err := plugin.Lookup("GetTypeEx")
	if err != nil {
		return err
	}

	ifVariableReplaceIt, err := plugin.Lookup("IfVariableReplaceIt")
	if err != nil {
		return err
	}

	stringToText, err := plugin.Lookup("StringToText")
	if err != nil {
		return err
	}

	textToString, err := plugin.Lookup("TextToString")
	if err != nil {
		return err
	}

	getVariableValue, err := plugin.Lookup("GetVariableValue")
	if err != nil {
		return err
	}

	setVariable, err := plugin.Lookup("SetVariable")
	if err != nil {
		return err
	}

	printError, err := plugin.Lookup("PrintError")
	if err != nil {
		return err
	}

	runCode, err := plugin.Lookup("RunCode")
	if err != nil {
		return err
	}

	parseFile, err := plugin.Lookup("ParseFile")
	if err != nil {
		return err
	}

	execute, err := plugin.Lookup("Execute")
	if err != nil {
		return err
	}

	setClassUsage, err := plugin.Lookup("SetClassUsage")
	if err != nil {
		return err
	}

	printString, err := plugin.Lookup("PrintString")
	if err != nil {
		return err
	}

	PiniginShell = piniginShell.(func())
	SetReturned = setReturned.(func(varType string, value interface{}, segmentName string))
	AddClass = addClass.(func(name string, usedByDefault bool, caller func(args []string, segmentName string) bool, loader func() bool) bool)
	RemoveClass = removeClass.(func(name string) bool)
	CheckOnVariable = checkOnVariable.(func(str string) bool)
	CleanUp = cleanUp.(func(checkClasses bool) bool)
	GetType = getType.(func(str string) string)
	GetTypeOfVar = getTypeOfVar.(func(str string) string)
	GetTypeEx = getTypeEx.(func(str string) string)
	IfVariableReplaceIt = ifVariableReplaceIt.(func(str string) interface{})
	StringToText = stringToText.(func(str string) string)
	TextToString = textToString.(func(str string) string)
	GetVariableValue = getVariableValue.(func(v string) interface{})
	SetVariable = setVariable.(func(name string, value interface{}) bool)
	PrintError = printError.(func(str string))
	RunCode = runCode.(func(code string) bool)
	ParseFile = parseFile.(func(path string) bool)
	Execute = execute.(func(code string, segmentName string, line int, isInIfElse bool) bool)
	SetClassUsage = setClassUsage.(func(className string, inUse bool, isGlobal bool) bool)
	PrintString = printString.(func(str string, newline bool))

	return nil
}
