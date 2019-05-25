package services

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/models"
	"sort"
	"sync"
)

type UserService struct {
	// users can have duplicate ids and names
	// so instead of map - use array
	users []models.User
	cfg   *config.Config
	fw    *FileWatcher
	mu    sync.RWMutex
}

func NewUserService(cfg *config.Config) (*UserService, error) {
	var err error
	//log.Printf("NewUserService")
	s := &UserService{
		cfg: cfg,
	}
	file := s.cfg.UserFile
	// watch for file changes
	// we could reload the file on every change
	// instead we just keep dirty flag and reload only when there is web request
	// to test: sed -i'.bak' -e 's/adm/admX/g' test/user
	if s.fw, err = NewFileWatcher(file); check(err) {
		return s, err
	}
	err = s.loadUsers(file)
	if check(err) {
		return nil, err
	}

	go s.fw.Watch()
	return s, nil
}

// FindUsers searches users matching given example
// if example is nil - all users are returned
// if id is > 0 - matching GID only, otherwise - do not check GID
func (s *UserService) FindUsers(example *models.User) []models.User {
	s.loadIfDirty()
	l := len(s.users)
	// sanity check
	if l == 0 {
		return nil
	}
	if example == nil {
		return s.users
	}
	match := usersByIdName(example, s.users)
	res := make([]models.User, 0)
	for _, u := range match {
		if matched(*example, u) {
			res = append(res, u)
		}
	}
	return res
}

// check user vs example rec for exact match on non-empty fields
func matched(ex models.User, u models.User) bool {
	if ex.Comment != "" && u.Comment != ex.Comment {
		return false
	}
	if ex.Home != "" && u.Home != ex.Home {
		return false
	}
	if ex.Shell != "" && u.Shell != ex.Shell {
		return false
	}
	if ex.GID >= 0 && u.GID != ex.GID {
		return false
	}
	// name and UID should be compared already - double check
	if ex.UID >= 0 && u.UID != ex.UID {
		return false
	}
	if ex.Name != "" && u.Name != ex.Name {
		return false
	}
	return true
}

// reload user array if file has been modified
func (s *UserService) loadIfDirty() {
	if s.fw.HasChanged() {
		err := s.loadUsers(s.cfg.UserFile)
		check(err)
	}
}

func (s *UserService) loadUsers(fileName string) error {
	//log.Printf("loadUsers: %v", fileName)
	lines, err := readLines(fileName)
	if err != nil {
		return err
	}
	list := make([]models.User, 0)
	for _, line := range lines {
		u, err := models.NewUser(line)
		// ignore bad lines
		if err != nil {
			continue
		}
		list = append(list, *u)
	}

	//sort by name to enable binary search
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})

	// write lock
	s.mu.Lock()
	s.users = list
	// clear dirty flag
	s.fw.SetDirty(false)
	s.mu.Unlock()
	return nil
}

// usersByIdName matches users given example user with id and/or name
// id -1 means no check
func usersByIdName(example *models.User, users []models.User) []models.User {
	if example == nil {
		return users
	}
	match := make([]models.User, 0)
	id := example.UID
	// binary search by name
	if len(example.Name) > 0 {
		i := sort.Search(len(users),
			func(i int) bool { return users[i].Name >= example.Name })

		if i < len(users) && users[i].Name == example.Name {
			g := users[i]
			// it is possible to have duplicate user names and ids -
			// but we won't bother as it is border case
			if id < 0 || id == g.UID {
				match = append(match, g)
			}
		}
		return match
	}
	// no name - check uid, assume duplicates, no sort by uid
	for _, g := range users {
		if id < 0 || g.UID == id {
			match = append(match, g)
		}
	}
	return match
}
