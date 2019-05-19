package services

import (
	"github.com/ivost/nix_users/internal/config"
	"github.com/ivost/nix_users/internal/models"
)


type GroupService struct {
	//metadata *models.Metadata
	cfg      *config.Config
}

func NewGroupService(cfg *config.Config) (*GroupService, error) {
	//mc := &GroupService{logger: logger}
	mc := &GroupService{cfg: cfg}
	//log.Printf("=== New Metadata Service")
	//mc.cfg, _ = config.InitConfig()
	//err := mc.LoadMeta()
	return mc, nil
}

func (ms *GroupService) GetGroup(id string) *models.Group {
	//log.Printf("=== GetMeta %v", id)
	//val, ok := ms.metadata.M.Load(id)
	//if !ok {
	//	return nil
	//}
	//res := val.(models.Group)
	//return &res
	return nil
}

func (ms *GroupService) GetGroupsAll() []models.Group {
	res := make([]models.Group, 0)
	//ms.metadata.M.Range(func(k, v interface{}) bool {
	//	//fmt.Printf("GetMetaAll %+v=%+v\n", k, v)
	//	res = append(res, v.(models.Group))
	//	return true
	//})
	return res
}




