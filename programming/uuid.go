package programming

import (
	"strings"

	"github.com/google/uuid"
)

// NewUuid generates an UUID with the possibility
// to remove the hyphens
func (pf *ProgrammingFunctions) NewUuid(withoutHyphen bool) string {
	uuidWithHyphen := uuid.New()

	if withoutHyphen {
		return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	}

	return uuidWithHyphen.String()
}
