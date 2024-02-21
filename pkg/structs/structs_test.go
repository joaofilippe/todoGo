package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Config struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     int    `yaml:"port" env:"DB_PORT"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"dbname" env:"DB_NAME"`
	Dsn      string
}

type TSt struct{
	Name string
	Age int
	Active bool
}

func Test_GetTags(t *testing.T) {


	c := Config{}

	values, err := GetTagsValues(c, "env")
	assert.NotEmpty(t, len(values), "no values returned")
	assert.Nil(t, err, "error must be nil")
}

func Test_SetStringFieldValueByName(t *testing.T) {
	tS := new(TSt)

	err := SetStringFieldValueByName(tS, "Nme", "João Filippe")
	assert.Nil(t, err, "não deveria retornar erro")
}