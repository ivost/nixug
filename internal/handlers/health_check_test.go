package handlers

import (
	"testing"
)

func TestHealth(t *testing.T) {
	// Setup
	//e := echo.New()
	//req := httptest.NewRequest(http.MethodGet, "/health", nil)
	//rec := httptest.NewRecorder()
	//c := e.NewContext(req, rec)
	//c.SetPath("/health")
	//err := HealthCheck(c)
	//res := e.GET("/health", HealthCheck)
	//log.Printf("%+v", res)
	//// Assertions
	//if assert.NoError(t, err) {
	//	assert.Equal(t, http.StatusOK, rec.Code)
	//}
}

//c.SetParamNames("email")
//c.SetParamValues("jon@labstack.com")

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
