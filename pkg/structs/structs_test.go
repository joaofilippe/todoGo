package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetTags(t *testing.T) {
	type Config struct {
		Host     string `yaml:"host" env:"DB_HOST"`
		Port     int    `yaml:"port" env:"DB_PORT"`
		User     string `yaml:"user" env:"DB_USER"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		DBName   string `yaml:"dbname" env:"DB_NAME"`
		Dsn      string
	}

	c := Config{}

	values, err := GetTagsValues(c, "env")
	assert.NotEmpty(t, len(values), "no values returned")
	assert.Nil(t, err, "error must be nil")
}
