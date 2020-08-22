//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package middleware_test

import (
	"github.com/moemoe89/go-mongodb-gita/config"
	"github.com/moemoe89/go-mongodb-gita/routers"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	log := config.InitLog()

	router := routers.GetRouter(log, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/", nil)
	req.Header.Add("Access-Control-Request-Headers", "*")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
