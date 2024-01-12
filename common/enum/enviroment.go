package common

import "strings"

type Environment int

const (
	Development = iota + 1
	Staging
	Production

	Development_String = "DEVELOPMENT"
	Staging_String     = "STAGING"
	Production_String  = "PRODUCTION"
	Unknown_String     = "UNKNOWN"

	Development_Short = "DEV"
	Staging_Short     = "STG"
	Production_Short  = "PROD"
)

func (e Environment) ToString() string {
	return [...]string{"Development", "Staging", "Production"}[e-1]
}

func ParseToEnviroment(env string) Environment {
	switch strings.ToUpper(env) {
	case Development_String, Development_Short:
		return 1
	case Staging_String, Staging_Short:
		return 2
	case Production_String, Production_Short:
		return 3
	default:
		return 0
	}
}
