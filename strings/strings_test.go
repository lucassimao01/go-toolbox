package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistsInSlice(t *testing.T) {
	assert.Equal(t, true, ExistsInSlice("A", []string{"A", "B", "C"}))
	assert.Equal(t, true, ExistsInSlice("B", []string{"A", "B", "C"}))
	assert.Equal(t, true, ExistsInSlice("C", []string{"A", "B", "C"}))

	assert.Equal(t, false, ExistsInSlice("I", []string{"A", "B", "C"}))
}

func TestSnakeCase(t *testing.T) {
	assert.Equal(t, "text", SnakeCase("TEXT"))
	assert.Equal(t, "text", SnakeCase("text"))
	assert.Equal(t, "text_element", SnakeCase("textElement"))
	assert.Equal(t, "text_element", SnakeCase("TextElement"))
	assert.Equal(t, "text_element_id", SnakeCase("TextElementID"))
}

func TestCoalesce(t *testing.T) {
	assert.Equal(t, "first", Coalesce("first", ""))
	assert.Equal(t, "first", Coalesce("first", "second"))
	assert.Equal(t, "second", Coalesce("", "second"))
	assert.Equal(t, "", Coalesce("", ""))
}
