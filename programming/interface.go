package programming

type Interface interface {
	NewUuid(withoutHyphen bool) string
	DebugJWT(tokenString string) (string, string, error)
}

type ProgrammingFunctions struct {
}
