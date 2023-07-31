package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

func TestMain(m *testing.M){
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine{
	r := gin.Default()

	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool){
	// response recorder
	w := httptest.NewRecorder()

	// serve http
	r.ServeHTTP(w, req)

	if !f(w){
		t.Fail()
	}
}

// save main lists to temp lists for testing
func saveLists(){
	tmpArticleList = articleList
}

// restore main lists from temp lists
func restoreLists(){
	articleList = tmpArticleList
}