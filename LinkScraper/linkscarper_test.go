package linkscraper

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func Test_checkLink(t *testing.T) {
	r, _ := regexp.Compile(`\/([A-Za-z0-9-]{1,})([0-9]{1,}$)`)
	var linkTrue = "/testest"
	var linkFalse = "balbalabla"
	var noContent = ""

	assert.True(t, checkLink(linkTrue, r), "Link should be True")
	assert.False(t, checkLink(linkFalse, r), "Link should be False")
	assert.False(t, checkLink(noContent, r), "No Content should be False")
}
