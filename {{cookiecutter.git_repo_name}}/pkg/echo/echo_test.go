package echo_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"

	"{{cookiecutter.app_go_package_path}}/pkg/echo"
)

func TestEcho_toReturnApplicationName(t *testing.T) {
	// given
	expectedName := "Thank you for using the 3 Musketeers!"
	os.Setenv("ECHO_MESSAGE", expectedName)

	// when
	name := echo.Echo()

	// then
	assert.Equal(t, expectedName, name)
}
