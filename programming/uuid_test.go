package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pf ProgrammingFunctions = ProgrammingFunctions{}

func TestNewUuidWithHyphen(t *testing.T) {
	uuidWithHyphen := pf.NewUuid(false)

	assert.Len(t, uuidWithHyphen, 36)
	assert.Contains(t, uuidWithHyphen, "-")
}

func TestNewUuidWithoutHyphen(t *testing.T) {
	uuidWithHyphen := pf.NewUuid(true)

	assert.Len(t, uuidWithHyphen, 32)
	assert.NotContains(t, uuidWithHyphen, "-")
}
