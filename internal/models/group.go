package models

type Group struct {
	GID     int32    `json:"gid"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
}


func (m *Group) Validate(i interface{}) error {

	errs := new(RequestErrors)
	if m.Name == "" {
		errs.Append(ErrNameEmpty)
	}
	if errs.Len() == 0 {
		return nil
	}
	return errs
}
