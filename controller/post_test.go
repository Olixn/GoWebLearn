package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
    "title":"999",
    "content":"ssssss！",
    "community_id":2
}`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "需要登陆")

	// 2
	res := new(ResponseDate)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Error("json.Unmarshal")
	}

}
