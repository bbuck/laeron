package scripting

import (
	"fmt"

	"github.com/d5/tengo/compiler/token"
)

type InvalidOperationError struct {
	op       token.Token
	typeName string
}

func newInvalidOperationError(op token.Token, typeName string) *InvalidOperationError {
	return &InvalidOperationError{op: op, typeName: typeName}
}

func (ioe *InvalidOperationError) Error() string {
	return fmt.Sprintf("invalid operation %s on %s", ioe.op, ioe.typeName)
}

type ArgumentCountError struct {
	expected int
	received int
}

func newArgumentCountError(expected, received int) *ArgumentCountError {
	return &ArgumentCountError{expected: expected, received: received}
}

func (ace *ArgumentCountError) Error() string {
	return fmt.Sprintf("expected %d argument(s), received %d", ace.expected, ace.received)
}
