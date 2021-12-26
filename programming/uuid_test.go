package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUuidWithHyphen(t *testing.T) {
	var pf ProgrammingFunctions = ProgrammingFunctions{}

	uuidWithHyphen := pf.NewUuid(false)

	assert.Len(t, uuidWithHyphen, 36)
	assert.Contains(t, uuidWithHyphen, "-")
}

func TestNewUuidWithoutHyphen(t *testing.T) {
	var pf ProgrammingFunctions = ProgrammingFunctions{}

	uuidWithHyphen := pf.NewUuid(true)

	assert.Len(t, uuidWithHyphen, 32)
	assert.NotContains(t, uuidWithHyphen, "-")
}
