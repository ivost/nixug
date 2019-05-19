package services

import (
	"testing"
)

var (
	testFile = ""
)

//var map1 = test.NewConfig()

//func TestSave(t *testing.T) {
//	tests := map[string]struct {
//		path string
//		mm   models.Metadata
//		want error
//		good bool
//	}{
//		"create1": {
//			path: testMetaFile,
//			mm:   map1,
//			good: true,
//		},
//	}
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

//func TestLoad(t *testing.T) {
//	tests := map[string]struct {
//		input string
//		want  *models.Metadata
//		good  bool
//	}{
//		//"bad": {
//		//	input: "foo",
//		//	want:  nil,
//		//	good:  false,
//		//},
//		"m1": {
//			input: testMetaFile,
//			want:  &map1,
//			good:  true,
//		},
//	}
//
//	cfg, _ := config.ReadConfig("")
//
//	// ensure files are written
//	TestSave(t)
//
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			cfg.MetaFile = tc.input
//			ms, err := NewGroupService(cfg)
//			assert.Nil(t, err)
//			err = ms.LoadMeta()
//			assert.Nil(t, err)
//
//			t.Logf("metadata %+v", ms.metadata)
//
//			if tc.good && err != nil {
//				t.Fatalf("**** load error: %v", err)
//			}
//			if tc.want != nil {
//				tc.want.M.Range(func(key, value interface{}) bool {
//					t.Logf("key %v", key)
//					x := ms.GetMeta(key.(string))
//					assert.NotNil(t, x)
//					if x == nil {
//						t.Fail()
//						return false
//					}
//					y := value.(models.Group)
//					// path may be different
//					//y.Path = x.Path
//					assert.EqualValues(t, *x, y)
//					return true
//				})
//			}
//		})
//	}
//}

//func TestGetMeta(t *testing.T) {
//	cfg, err := config.ReadConfig("")
//	assert.Nil(t, err)
//	ms, err := NewGroupService(cfg)
//	assert.Nil(t, err)
//	v1 := ms.GetMeta(test.MetaId1)
//	assert.NotNil(t, v1)
//	m1, _ := map1.M.Load(test.MetaId1)
//	if m1 == nil {
//		t.Fail()
//	}
//
//	// path may be different
//	mm := m1.(models.Group)
//	//mm.Path = v1.Path
//	assert.EqualValues(t, mm, *v1)
//
//	v1 = ms.GetMeta("foo")
//	assert.Nil(t, v1)
//
//	lst := ms.GetMetaAll()
//	assert.True(t, len(lst) > 0)
//}

func TestGetPath(t *testing.T) {
	//p1 := "http://0.0.0.0:8484/v1/stream/cam1"
	//p2 := "http://0.0.0.0:8484/v1/string/cam1"
	//
	//cfg, err := config.InitConfig()
	//assert.Nil(t, err)
	//ms, err := NewGroupService(cfg)
	//assert.Nil(t, err)
	//
	//a1 := ms.GetMeta(test.MetaId1)
	//assert.NotNil(t, a1)
	//if a1 != nil {
	//	p := cfg.GetPath(a1.Type, a1.Id)
	//	assert.Equal(t, p1, p)
	//}
	//
	//a2 := ms.GetMeta(test.MetaId2)
	//assert.NotNil(t, a2)
	//if a2 != nil {
	//	p := cfg.GetPath(a2.Type, a2.Id)
	//	assert.Equal(t, p2, p)
	//}
}

//func TestCopy(t *testing.T) {
//	src := models.Group{Id: test.MetaId1, Name: "bar",  TTL: "1001m"}
//	dst := models.Group{Id: test.MetaId1, Name: "foo",  Cap: 101}
//	// reflection copy, omit "" and 0 from source
//	copyFields(&src, &dst)
//	//log.Printf("src %+v", src)
//	//log.Printf("dst %+v", dst)
//
//	assert.Equal(t, test.MetaId1, dst.Id)
//	assert.Equal(t, src.Name, dst.Name)
//	assert.EqualValues(t, 101, dst.Cap)
//	assert.EqualValues(t, "1001m", dst.TTL)
//}
//
