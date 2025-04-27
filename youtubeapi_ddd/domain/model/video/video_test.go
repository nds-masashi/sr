package video

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_正常(t *testing.T) {
	// Given
	url := "http://localhost"
	title := "テスト"

	// When
	got := New(url, title)

	// Then
	assert.Equal(t, Model{
		url:   url,
		title: title,
	}, got)
}

func Test_GetURL_正常(t *testing.T) {
	// Given
	url := "http://localhost"
	title := "テスト"
	model := New(url, title)

	// When
	got := model.GetURL()

	// Then
	assert.Equal(t, url, got)
}

func Test_GetTitle_正常(t *testing.T) {
	// Given
	url := "http://localhost"
	title := "テスト"
	model := New(url, title)

	// When
	got := model.GetTitle()

	// Then
	assert.Equal(t, title, got)
}
