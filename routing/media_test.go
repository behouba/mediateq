package routing

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestQueryParamToInt(t *testing.T) {
	tests := []struct {
		query      string
		queryValue string
		expected   int
	}{
		{
			query:      "offset",
			queryValue: "100",
			expected:   100,
		},
		{
			query:      "limit",
			queryValue: "50",
			expected:   50,
		},

		{
			query:      "offset",
			queryValue: "invalid",
			expected:   0,
		},
	}

	for _, test := range tests {
		// Set up the request context
		ctx, _ := gin.CreateTestContext(nil)
		req, _ := http.NewRequest("GET", "/test?"+test.query+"="+test.queryValue, nil)
		ctx.Request = req

		// Call the queryParamToInt function
		actual := queryParamToInt(ctx, test.query)

		assert.Equal(t, test.expected, actual)
	}
}
