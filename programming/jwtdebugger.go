package programming

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// DebugJWT parses a JWT and returns the header and payload contents
// WARNING: this function does not validate the token, only inspects the content
func (pf *ProgrammingFunctions) DebugJWT(tokenString string) (string, string, error) {

	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", "", fmt.Errorf("error parsing token: %s", err.Error())
	}

	header, _ := json.Marshal(token.Header)
	payload, _ := json.Marshal(token.Claims)

	return string(header), string(payload), nil
}
