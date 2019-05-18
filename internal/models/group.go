package models

import "sync"

type Metadata struct {
	// map Id -> Group
	M sync.Map `json:"-"`
	// array of Group
	A []Group `json:""`
}

// Group structure represents metadata about each hash key/stream/timeseries
type Group struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`    			// string, map(hash), stream
	//// optional
	//Content   string `json:"content,omitempty"` // text (default), json, bin
	//Persist   string `json:"pers,omitempty"`    // tmp (default), none, disk
	//TTL       string `json:"ttl,omitempty"`     // TTL duration - e.g. 1h, "" = no expiration (not applicable for stream)
	//Cap       int64  `json:"cap,omitempty"`     // cap - only for streams - approx.number of elements to keep
	//// for TS
	//Labels	  []string `json:"labels,omitempty"`
	//Rules	  []string `json:"rules,omitempty"`
	//Source    string  `json:"source,omitempty"`
	//Dest      string  `json:"dest,omitempty"`
	//Agg       string  `json:"agg,omitempty"`
	//BucketSize *int64  `json:"bucket_size,omitempty"`
}

//"labels": ["vni", "fast"],
//"rules": ["ts/rule_vni_speed_avg_1h"]
//"source": "ts/speed",
//"dest": "ts/avg_vni_speed_1hr",
//"agg": "avg",
//"bucket_size": 3600
//Path      string `json:"path"`

func (md *Metadata) Add(m Group) {
	md.A = append(md.A, m)
	md.M.Store(m.Id, m)
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
