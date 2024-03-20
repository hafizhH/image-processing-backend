package bootstrap

import (
	_ "image-processing-backend/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	app := App()

	assert.NotNil(t, app, "App object should be returned")
	assert.NotNil(t, app.Env, "Returned App object should contain Env")
}

func TestEnv(t *testing.T) {
	env, err := NewEnv()

	assert.NotNil(t, env, "Env object should be returned")
	assert.Nil(t, err, "There should be no error occured")
}
