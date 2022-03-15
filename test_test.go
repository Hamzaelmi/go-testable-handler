package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingPong(t *testing.T) {

	ctrl := PublicController{Database: "hey"}
	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctrl.pong(ctx)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	x := fmt.Sprintf("{\"msg\":\"hey\"}")
	final := string(data)
	if final != x {
		t.Errorf("expected ABC got %v", string(data))
	}
}
