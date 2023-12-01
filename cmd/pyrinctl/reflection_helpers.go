package main

import (
	"reflect"
	"unicode"

	"github.com/Pyrinpyi/pyipad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

// protobuf generates the command types with two types:
// 1. A concrete type that holds the fields of the command bearing the name of the command with `RequestMessage` as suffix
// 2. A wrapper that implements isPyipadMessage_Payload, having a single field pointing to the concrete command
//    bearing the name of the command with `PyipadMessage_` prefix and `Request` suffix

// unwrapCommandType converts a reflect.Type signifying a wrapper type into the concrete request type
func unwrapCommandType(requestTypeWrapped reflect.Type) reflect.Type {
	return requestTypeWrapped.Field(0).Type.Elem()
}

// unwrapCommandValue convertes a reflect.Value of a pointer to a wrapped command into a concrete command
func unwrapCommandValue(commandValueWrapped reflect.Value) reflect.Value {
	return commandValueWrapped.Elem().Field(0)
}

// isFieldExported returns true if the given field is exported.
// Currently the only way to check this is to check if the first rune in the field's name is upper case.
func isFieldExported(field reflect.StructField) bool {
	return unicode.IsUpper(rune(field.Name[0]))
}

// generatepyipadMessage generates a wrapped pyipadMessage with the given `commandValue`
func generatepyipadMessage(commandValue reflect.Value, commandDesc *commandDescription) (*protowire.PyipadMessage, error) {
	commandWrapper := reflect.New(commandDesc.typeof)
	unwrapCommandValue(commandWrapper).Set(commandValue)

	pyipadMessage := reflect.New(reflect.TypeOf(protowire.PyipadMessage{}))
	pyipadMessage.Elem().FieldByName("Payload").Set(commandWrapper)
	return pyipadMessage.Interface().(*protowire.PyipadMessage), nil
}

// pointerToValue returns a reflect.Value that represents a pointer to the given value
func pointerToValue(valuePointedTo reflect.Value) reflect.Value {
	pointer := reflect.New(valuePointedTo.Type())
	pointer.Elem().Set(valuePointedTo)
	return pointer
}
