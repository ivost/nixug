package test

import (
	"github.com/ivost/nix_users/internal/config"
)

const MetaId1 = "stream/cam1"
const MetaId2 = "string/cam1"

func NewConfig() *config.Config {
	host := "0.0.0.0"
	c := new (config.Config)
	c.Listener = config.Listener {Addr: host, Port: 8484}
	return c
}

//func NewMetadata() models.Metadata {
//	//map1 := models.Metadata{}
//	//cap := int64(10)
//	//// make sure Name matches real config
//	//map1.Add(models.Group{Id: MetaId1, Name: "cam1 stream", Type: models.TypeStream,
//	//	Persist: models.PersistTemp, Content: models.ContentBin, Cap: cap})
//	//map1.Add(models.Group{Id: MetaId2, Name: "cam1 config", Type: models.TypeString,
//	//	Persist: models.PersistDisk, Content: models.ContentText})
//	//map1.Add(models.Group{Id: MetaId3, Name: "VNI speed", Type: models.TypeTS,
//	//	Persist: models.PersistDisk, TTL: "30d"})
//	return map1
//}
