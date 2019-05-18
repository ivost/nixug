package services

import (
	"github.com/ivost/nix_users/internal/config"
	"github.com/ivost/nix_users/internal/models"
)
//"github.com/labstack/gommon/log"


type GroupService struct {
	metadata *models.Metadata
	cfg      *config.Config
}

func NewGroupService(cfg *config.Config) (*GroupService, error) {
	//mc := &GroupService{logger: logger}
	mc := &GroupService{cfg: cfg}
	//log.Printf("=== New Metadata Service")
	//mc.cfg, _ = config.InitConfig()
	err := mc.LoadMeta()
	return mc, err
}

func (ms *GroupService) GetMeta(id string) *models.Group {
	//log.Printf("=== GetMeta %v", id)
	val, ok := ms.metadata.M.Load(id)
	if !ok {
		return nil
	}
	res := val.(models.Group)
	return &res
}

func (ms *GroupService) GetMetaAll() []models.Group {
	res := make([]models.Group, 0)

	ms.metadata.M.Range(func(k, v interface{}) bool {
		//fmt.Printf("GetMetaAll %+v=%+v\n", k, v)
		res = append(res, v.(models.Group))
		return true
	})
	return res
}

// embed in LoadMeta
//func (ms *GroupService) LoadDefaultMeta(path string) (*models.Metadata, error) {
//	return ms.LoadMeta(MetadataFile)
//}
func (ms *GroupService) LoadMeta() error {
	//log.Printf("** LoadMeta")
	//f := ms.cfg.GroupFile
	//d, err := ioutil.ReadFile(f)
	//if err != nil {
	//	log.Printf("file %v not found", f)
	//	return err
	//}

	//ms.metadata = new(models.Metadata)
	//err = json.Unmarshal(d, &ms.metadata.A)
	//if err != nil {
	//	log.Printf("invalid json - file %v", f)
	//	return err
	//}
	//
	//for _, v := range ms.metadata.A {
	//	//v.Path = ms.cfg.GetPath(v.Type, v.Id)
	//	if len(v.Persist) == 0 {
	//		v.Persist = "tmp"
	//	}
	//	//log.Printf("v %+v", v)
	//	ms.metadata.M.Store(v.Id, v)
	//}
	//log.Printf("** %+v", ms.metadata.M)
	return nil
}

func (ms *GroupService) Create(m models.Group) error {
	////m.Path = ms.cfg.GetPath(m.Type, m.Id)
	//if len(m.Persist) == 0 {
	//	m.Persist = "tmp"
	//}
	////fmt.Printf("******* Create %+v\n", m)
	//ms.metadata.M.Store(m.Id, m)
	////ms.metadata.A = append(ms.metadata.A, m)
	////x, ok := ms.metadata.M.Load(m.Id)
	////fmt.Printf("++++ After Create ok %v %+v\n", ok, x)
	//todo: save
	return nil
}



