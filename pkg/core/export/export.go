package export

import (
	"encoding/json"
	"encoding/xml"
	"github.com/jszwec/csvutil"
	sp2 "github.com/s4bb4t/lighthouse/pkg/core/sp"
)

const (
	Xml  = "application/xml"
	Json = "application/json"
	Csv  = "text/csv"
)

type Error struct {
	Msg   string `csv:"msg,omitempty" xml:"msg,omitempty"`
	Desc  string `csv:"desc,omitempty" xml:"desc,omitempty"`
	Hint  string `csv:"hint,omitempty" xml:"hint,omitempty"`
	Time  string `csv:"time,omitempty" xml:"time,omitempty"`
	Path  string `csv:"path,omitempty" xml:"path,omitempty"`
	Level string `csv:"level,omitempty" xml:"level,omitempty"`
	Cause string `csv:"cause,omitempty" xml:"cause,omitempty"`
}

func JSON(e *sp2.Error) ([]byte, error) {
	return json.Marshal(e)
}

func CSV(errs ...*sp2.Error) ([]byte, error) {
	return exp(Csv, errs...)
}

func XML(errs ...*sp2.Error) ([]byte, error) {
	return exp(Xml, errs...)
}

func exp(_type string, errs ...*sp2.Error) ([]byte, error) {
	var arr []Error
	for _, e := range errs {
		err := Error{}
		err.Path = e.ReadSource()
		err.Msg = e.ReadMsg(sp2.En)
		err.Desc = e.ReadDesc()
		err.Hint = e.ReadHint()
		err.Time = e.ReadTime().String()
		if e.ReadCaused() != nil {
			err.Cause = e.ReadCaused().Error()
		}
		arr = append(arr, err)
	}

	switch _type {
	case Xml:
		return xml.Marshal(arr)
	case Csv:
		return csvutil.Marshal(arr)
	default:
		return JSON(errs[0])
	}
}
