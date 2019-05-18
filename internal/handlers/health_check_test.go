package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	// Setup
	//todo: redis init
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/health")
	//c.SetParamNames("email")
	//c.SetParamValues("jon@labstack.com")
	err := HealthCheck(c)
	// 	e.GET("/health", handlers.HealthCheck)
	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

//func TestLoad(t *testing.T) {
//	tests := map[string]struct {
//		input string
//		want  *models.MetaMap
//		good  bool
//	}{
//		"bad": {
//			input: "foo",
//			want:  nil,
//			good:  false,
//		},
//	}
//
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			m, err := LoadMeta(tc.input)
//			if tc.good && err != nil {
//				t.Fatalf("load error: %v", err)
//			}
//			if tc.want != nil {
//				tc.want.Range(func(key, value interface{}) bool {
//					x, ok := m.Load(key)
//					assert.True(t, ok)
//					if !reflect.DeepEqual(x.(models.Meta), value.(models.Meta)) {
//						t.Fatalf("expected: %v, got: %v", tc.want, m)
//					}
//					return true
//				})
//			}
//		})
//	}
//}
