package test

import (
	"ewallet/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	id := utils.GenerateID()

	// Проверяем, что id не пустой
	assert.NotEmpty(t, id, "Generated ID should not be empty")

}
