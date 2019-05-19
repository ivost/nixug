package handlers

import (
	"github.com/ivost/nixug/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

type Handler interface {
	DoGet() error
	//DoInfo() error
	//DoPost() error
	//DoPut() error
	//DoPatch() error
	//DoDelete() error
}

type handler struct {
	context  *Context
	lock     sync.RWMutex
}

func (h *handler) PrepGet() (*models.ReadParam, error) {
	c := h.context
	err := c.Valid()
	if err != nil {
		return nil, err
	}
	//log.Printf("client accepts type %v, stream content type is %v", at, m.Content)
	rp := parseQuery(c)
	return &rp, nil
}

//func (h *handler) PrepTSGet() (*models.TSReadParam, error) {
//	c := h.context
//	err := c.Valid()
//	if err != nil {
//		return nil, err
//	}
//	//log.Printf("client accepts type %v, stream content type is %v", at, m.Content)
//	rp := parseTSQuery(c)
//	return &rp, nil
//}



func DoGet(c echo.Context) error {
	if h, err := requestCheck(c); err == nil {
		return h.DoGet()
	} else {
		return err
	}
}


func requestCheck(c echo.Context) (Handler, error) {
	//if h := pickHandler(c); h == nil {
	//	return nil, c.JSON(http.StatusBadRequest, "no handler found")
	//} else {
	//	return nil, nil
	//}

		return nil, c.JSON(http.StatusBadRequest, "no handler found")

}

//func pickHandler(c echo.Context) Handler {
//	t := c.Param("t")
//	switch t {
//	case models.TypeStream:
//		return NewStreamHandler(c)
//	case models.TypeString:
//		return NewStringHandler(c)
//	case models.TypeMap:
//		return NewMapHandler(c)
//	case models.TypeTS:
//		return NewTSHandler(c)
//
//	default:
//		log.Printf("no handler for type %v", t)
//	}
//
//	return nil
//}


func parseQuery(c echo.Context) models.ReadParam {
	var rp models.ReadParam
	s := getStrQueryParam(c, "start")
	e := getStrQueryParam(c, "end")
	n := getIntQueryParam(c, "limit")
	st := getIntQueryParam(c, "stream")

	if s != nil {
		rp.Start = *s
	}
	if e != nil {
		rp.End = *e
	}
	if n != nil {
		rp.Limit = int64(*n)
	}
	if st != nil {
		rp.Stream = *st != 0
	}
	return rp
}

//func parseTSQuery(c echo.Context) models.TSReadParam {
//	var rp models.TSReadParam
//	s := getStrQueryParam(c, "start")
//	e := getStrQueryParam(c, "end")
//	n := getIntQueryParam(c, "limit")
//
//	if s != nil {
//		rp.Start = *s
//	}
//	if e != nil {
//		rp.End = *e
//	}
//	if n != nil {
//		rp.Limit = int64(*n)
//	}
//	return rp
//}

//func parseQueryWrite(c echo.Context) models.WriteParam {
//	var wp models.WriteParam
//
//	pt := getIntQueryParam(c, "ttl")
//
//	if pt != nil {
//		wp.TTL = int64(*pt)
//	}
//	return wp
//}

//// returns "" on errors, else mime content type accepted by client
//func GetContentType(c echo.Context) string {
//	h := c.Get(echo.HeaderContentType).(string)
//	switch h {
//	case echo.MIMEOctetStream,
//		echo.MIMETextPlain, echo.MIMETextPlainCharsetUTF8,
//		echo.MIMEApplicationJSON, echo.MIMEApplicationJSONCharsetUTF8,
//		models.HeaderSSE:
//		return h
//	default:
//		return ""
//	}
//	return ""
//}

//func IsJSON(h string) bool {
//	switch h {
//	case echo.MIMEApplicationJSON,
//		 echo.MIMEApplicationJSONCharsetUTF8:
//		return true
//	}
//	return false
//}

//// returns "" on errors, else mime content type accepted by client
//func GetAcceptType(c echo.Context) string {
//	//ct := c.Get(echo.HeaderContentType).(string)
//	h := c.Request().Header.Get(echo.HeaderAccept)
//	//log.Printf("%v: %+v", echo.HeaderAccept, h)
//	if len(h) == 0 {
//		return ""
//	}
//	if h == "*/*" {
//		return echo.MIMEApplicationJSON
//	}
//
//	switch h {
//	case echo.MIMEOctetStream,
//		echo.MIMETextPlain, echo.MIMETextPlainCharsetUTF8,
//		echo.MIMEApplicationJSON, echo.MIMEApplicationJSONCharsetUTF8,
//		models.HeaderSSE:
//		return h
//	default:
//		return ""
//	}
//	return ""
//}

func getStrQueryParam(c echo.Context, name string) *string {
	s := c.QueryParam(name)
	if len(s) == 0 {
		return nil
	}
	return &s
}

func getIntQueryParam(c echo.Context, name string) *int {
	s := getStrQueryParam(c, name)
	if s == nil {
		return nil
	}
	i, err := strconv.Atoi(*s)
	if err != nil {
		return nil
	}
	return &i
}

//func check(err error) bool {
//	if err == nil {
//		return false
//	}
//	log.Print(err.Error())
//	return true
//}
