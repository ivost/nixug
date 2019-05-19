package services

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/models"
	"io/ioutil"
	"sync"
)

type GroupService struct {
	groups  map[int]models.Group
	mu      sync.RWMutex
	changed bool
}

func NewGroupService() (*GroupService, error) {
	s := &GroupService{
		changed: true,
	}
	err := s.loadGroups()
	return s, err
}

func (s *GroupService) loadGroups() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.groups == nil {
		s.groups = make(map[int]models.Group)
	} else {
		// clear, avoid allocations
		for k := range s.groups {
			delete(s.groups, k)
		}
	}
	// read groups file
	cfg, err := config.NewConfig("")
	d, err := ioutil.ReadFile(cfg.GroupFile)
	if check(err) {
		return err
	}
	// todo
	_ = d
	g := models.Group{GID: 0, Name: "root", Members: []string{"root"}}
	s.groups[g.GID] = g
	s.changed = false
	return nil
}

/*
GET /groups/query[?name=<nq>][&gid=<gq>][&member=<mq1>[&member=<mq2>][&. ..]]
Return a list of groups matching all of the specified query fields. The bracket notation indicates that any of the following query parameters may be supplied:
- name
- gid
- member (repeated)
Any group containing all the specified members should be returned, i.e. when query members are a subset of group members.
Example Query: ​GET /groups/query?member=_analyticsd&member=_networkd Example Response:
[
{“name”: “_analyticsusers”, “gid”: 250, “members”: [“_analyticsd’,”_networkd”,”_timed”]}
]
*/
// FindGroups searches groups matching given example
func (s *GroupService) FindGroups(example *models.Group) []models.Group {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.changed {

	}
	//todo
	res := make([]models.Group, 0)

	for k := range s.groups {
		res = append(res, s.groups[k])
	}
	return res
}
