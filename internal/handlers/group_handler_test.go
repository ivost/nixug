package handlers

//var (
//	n1   = "cam1"
//	n2   = "config"
//	m1   = models.Meta{Name: n1, Type: "bin", Cap: 10}
//	m2   = models.Meta{Name: n2, Type: "map", Cap: 1}
//	map1 models.MetaMap
//)
//
//func init() {
//	//map1 = new(MetaMap)
//	map1.Store(n1, m1)
//	map1.Store(n2, m2)
//}
//
//
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			err := SaveMeta(tc.path, tc.mm)
//			if tc.good && err != nil {
//				t.Fatalf("error: %v", err)
//			}
//		})
//	}
//}
//
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
//		"m1": {
//			input: "/tmp/m1",
//			want:  &map1,
//			good:  true,
//		},
//		MetadataFile: {
//			input: MetadataFile,
//			want:  &map1,
//			good:  true, // assuming file is present
//		},
//	}
//
//	// ensure files are written
//	TestSave(t)
//	_, err := os.Stat(MetadataFile)
//	if err != nil {
//		w, _ := os.Create(MetadataFile)
//		r, _ := os.Open("x")
//		io.Copy(w, r)
//		w.Close()
//		r.Close()
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
//
//
//func TestGetMeta(t *testing.T) {
//	h, err := NewMetaHandler()
//	assert.Nil(t, err)
//	v1 := h.GetMeta(n1)
//	assert.EqualValues(t, m1, v1)
//}
