package services

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/models"
	"log"
	"sort"
	"sync"
)

type GroupService struct {
	// groups can have duplicate ids and names
	// so instead of map - use array
	groups []models.Group
	cfg    *config.Config
	fw     *FileWatcher
	mu     sync.RWMutex
}

func NewGroupService(cfg *config.Config) (*GroupService, error) {
	var err error
	log.Printf("NewGroupService")
	s := &GroupService{
		cfg: cfg,
	}
	file := s.cfg.GroupFile
	// watch for file changes
	// we could reload the file on every change
	// instead we just keep dirty flag and reload only when there is web request
	// to test: sed -i'.bak' -e 's/adm/admX/g' test/group
	if s.fw, err = NewFileWatcher(file); check(err) {
		return s, err
	}
	err = s.loadGroups(file)
	if check(err) {
		return nil, err
	}

	go s.fw.Watch()
	return s, nil
}

// FindGroups searches groups matching given example
// if example is nil - all groups are returned
// if id is > 0 - matching GID only, otherwise - do not check GID
func (s *GroupService) FindGroups(example *models.Group) []models.Group {
	s.loadIfDirty()
	l := len(s.groups)
	// sanity check
	if l == 0 {
		return nil
	}
	if example == nil {
		return s.groups
	}
	match := groupsByIdName(example, s.groups)
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

// reload group array if file has been modified
func (s *GroupService) loadIfDirty() {
	if s.fw.HasChanged() {
		err := s.loadGroups(s.cfg.GroupFile)
		check(err)
	}
}

func (s *GroupService) loadGroups(fileName string) error {
	log.Printf("loadGroups: %v", fileName)
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
		return list[i].Name < list[j].Name
	})

	// write lock
	s.mu.Lock()
	s.groups = list
	// clear dirty flag
	s.fw.SetDirty(false)
	s.mu.Unlock()
	return nil
}

// groupsByIdName matches groups given example group with id and/or name
// id -1 means no check
func groupsByIdName(example *models.Group, groups []models.Group) []models.Group {
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
