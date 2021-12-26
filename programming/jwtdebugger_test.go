package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugJWT(t *testing.T) {
	// arrange
	var pf ProgrammingFunctions = ProgrammingFunctions{}
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	// act
	header, payload, err := pf.DebugJWT(tokenString)

	// assert
	expectedHeader := "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
	expectedPayload := "{\"iat\":1516239022,\"name\":\"John Doe\",\"sub\":\"1234567890\"}"

	assert.Nil(t, err)
	assert.Equal(t, expectedHeader, header)
	assert.Equal(t, expectedPayload, payload)
}

func TestDebugJWTWithInvalidToken(t *testing.T) {
	// arrange
	var pf ProgrammingFunctions = ProgrammingFunctions{}
	tokenString := "xxxxx.yyyyy.zzzzz"

	// act
	header, payload, err := pf.DebugJWT(tokenString)

	// assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error parsing token")
	assert.Empty(t, header)
	assert.Empty(t, payload)
}
