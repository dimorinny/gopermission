package gopermission

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AlwaysFalse struct{}

func (af AlwaysFalse) HasPermission(request *http.Request) bool {
	return false
}

type AlwaysTrue struct{}

func (at AlwaysTrue) HasPermission(request *http.Request) bool {
	return true
}

func TestPermission(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://someurl.com", nil)

	permission := New(AlwaysTrue{}, AlwaysFalse{})
	permission2 := New(AlwaysTrue{})
	permission3 := New(AlwaysTrue{})
	permission3.AddChecker(AlwaysFalse{})

	assert.False(t, permission.IsPermitted(req))
	assert.True(t, permission2.IsPermitted(req))
	assert.False(t, permission3.IsPermitted(req))
}
