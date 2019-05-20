package services

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/models"
	"sort"
	"sync"
)

type GroupService struct {
	// groups can have duplicate ids and names
	groups  []models.Group
	mu      sync.RWMutex
	changed bool
	cfg 	*config.Config
}

func NewGroupService() (*GroupService, error) {
	c, err := config.NewConfig("")
	if err != nil {
		return nil, err
	}
	s := &GroupService{
		cfg: c,
		changed: true,
	}
	err = s.loadGroups(s.cfg.GroupFile)
	return s, err
}

func (s *GroupService) loadGroups(fileName string) error {
	lines, err := readLines(fileName)
	if err != nil {
		return err
	}
	list := make([]models.Group, 0)
	for _, line := range lines {
		g, err := models.NewGroup(line)
		// ignore bad lines
		if err != nil {
			continue
		}
		list = append(list, *g)
	}
	//sort by name to enable binary search
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name > list[j].Name
	})
	// write lock
	s.mu.Lock()
	s.groups = list
	// clear changed flag
	s.changed = false
	s.mu.Unlock()
	return nil
}

/*
GET /groups/query[?name=<nq>][&gid=<gq>][&member=<mq1>[&member=<mq2>][&. ..]]
Return a list of groups matching all of the specified query fields.
The bracket notation indicates that any of the following query parameters may be supplied:
- name
- gid
- member (repeated)
Any group containing all the specified members should be returned,
i.e. when query members are a subset of group members.
Example Query: ​GET /groups/query?member=_analyticsd&member=_networkd
Example Response:
[
{“name”: “_analyticsusers”, “gid”: 250, “members”: [“_analyticsd’,”_networkd”,”_timed”]}
]
*/
// FindGroups searches groups matching given example
// if example is nil - all groups are returned
// if id is > 0 - matching GID only, otherwise - do not check GID
func (s *GroupService) FindGroups(example *models.Group) []models.Group {
	// read lock
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.changed {
		s.loadGroups(s.cfg.GroupFile)
	}
	l := len(s.groups)
	// sanity check
	if l == 0 {
		return nil
	}
	if example == nil {
		return s.groups
	}
	match := groupByIdName(example, s.groups)
	// members?
	if len(example.Members) == 0 {
		return match
	}
	// lets check members
	// we don't expect large number of values to match -
	// no sorting/bin.search at this time
	// we could sort target list
	res := make([]models.Group, 0)
	for _, g := range match {
		if containsAll(g.Members, example.Members) {
			res = append(res, g)
		}
	}
	return res
}

// groupByIdName matches groups given example group with id and/or name
// id -1 means no check
func groupByIdName(example *models.Group, groups []models.Group ) []models.Group {
	if example == nil {
		return groups
	}
	match := make([]models.Group, 0)
	id := example.GID
	// binary search by name
	if len(example.Name) > 0 {
		i := sort.Search(len(groups),
			func(i int) bool { return groups[i].Name >= example.Name })
		if i < len(groups) && groups[i].Name == example.Name {
			g := groups[i]
			//log.Printf("found %v at index %d\n", g, i)
			// it is possible to have duplicate group names and ids -
			// but we won't bother as it is border case
			if id < 0 || id == g.GID {
				match = append(match, g)
			}
		}
		return match
	}
	// no name - check gid, assume duplicates, no sort by gid
	for _, g := range groups {
		if id < 0 || g.GID == id {
			match = append(match, g)
		}
	}
	return match
}