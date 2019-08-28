package scripting

import (
	"github.com/d5/tengo/compiler/token"
	"github.com/d5/tengo/objects"
	"github.com/google/uuid"
)

const typeName = "google.UUID"

// UUIDScriptObject wraps a github.com/google/uuid.UUID object for use from
// within the Tengo scripting language.
type UUIDScriptObject struct {
	uuid.UUID
}

// NewUUID creates a new instance of the UUIDScriptObject which implements Tengo's
// Object interface, allowing it to be used from within the Tengo language.
func NewUUID(id uuid.UUID) *UUIDScriptObject {
	return &UUIDScriptObject{UUID: id}
}

// TypeName returns the string "google.UUID" for a UUIDScriptObject.
func (UUIDScriptObject) TypeName() string {
	return typeName
}

// String returns a friendly representation of the UUID as a string (in the
// typical UUID format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (u *UUIDScriptObject) String() string {
	return u.UUID.String()
}

// BinaryOp always raises an invalid operation error for a UUIDScriptObject
// as mathematical/binary operations don't make sense with the UUID object.
func (UUIDScriptObject) BinaryOp(op token.Token, _ objects.Object) (objects.Object, error) {
	return nil, newInvalidOperationError(op, typeName)
}

// Equals does a []byte comparison of the two UUID objects, identifying if a
// UUID object is, indeed, the same as another UUID object.
func (u *UUIDScriptObject) Equals(o objects.Object) bool {
	if ou, ok := o.(*UUIDScriptObject); ok {
		return u.UUID == ou.UUID
	}

	return false
}

// IsFalsy always reports that UUIDScriptObjects are never false, this means
// they can be used a "true" values in condition tests.
func (UUIDScriptObject) IsFalsy() bool {
	return false
}

// Copy makes a copy of the UUIDScriptObject.
func (u *UUIDScriptObject) Copy() objects.Object {
	return &UUIDScriptObject{UUID: u.UUID}
}
