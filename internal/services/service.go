package services

import (
	"github.com/ivost/nixug/internal/models"
	"log"
	"reflect"
)

//func prepareReadParam(m models.Group, rp models.ReadParam) models.ReadParam {
//	//res := models.ReadParam{Start: rp.Start, End: rp.End, Limit: rp.Limit, Stream: rp.Stream}
//	//if len(rp.Start) == 0 {
//	//	res.Start = "-"
//	//}
//	//if len(rp.End) == 0 {
//	//	res.End = "+"
//	//}
//	//if rp.Limit <= 0 {
//	//	res.Limit = m.Cap
//	//}
//	//return res
//}

// reflection copy, omit "" and 0 from source
func copyFields(src *models.Group, dst *models.Group) {
	s := reflect.ValueOf(src).Elem()
	//st := s.Type()

	d := reflect.ValueOf(dst).Elem()
	dt := d.Type()

	for i := 0; i < s.NumField(); i++ {
		sf := s.Field(i)
		sfv := sf.Interface()

		df := d.FieldByName(dt.Field(i).Name)

		if sf.IsValid() && df.IsValid() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if !df.CanSet() {
				continue
			}
			switch df.Kind() {
			case reflect.String:
				// update dst only if src is not empty
				str := sfv.(string)
				if len(str) > 0 {
					df.SetString(str)
				}
			case reflect.Int64:
				// update dst only if src is not 0
				nv := sfv.(int64)
				if nv != 0 {
					df.SetInt(nv)
				}
			case reflect.Bool:
				df.SetBool(sfv.(bool))
			default:
				log.Printf("TODO %v", df.Kind())
			}
		}

	}
}

func check(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	log.Print(s)
	return true
}
