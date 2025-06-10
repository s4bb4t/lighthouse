package export

import (
	"encoding/json"
	"encoding/xml"
	"github.com/jszwec/csvutil"
	"github.com/s4bb4t/lighthouse/pkg/core/sperror"
)

const (
	Xml  = "application/xml"
	Json = "application/json"
	Csv  = "text/csv"
)

type Error struct {
	Msg    string `csv:"msg,omitempty" xml:"msg,omitempty"`
	Desc   string `csv:"desc,omitempty" xml:"desc,omitempty"`
	Hint   string `csv:"hint,omitempty" xml:"hint,omitempty"`
	Source string `csv:"source,omitempty" xml:"source,omitempty"`
	Level  string `csv:"level,omitempty" xml:"level,omitempty"`
	Cause  string `csv:"cause,omitempty" xml:"cause,omitempty"`
}

func JSON(e *sperror.Error) ([]byte, error) {
	return json.Marshal(e)
}

func CSV(errs ...*sperror.Error) ([]byte, error) {
	return exp(Csv, errs...)
}

func XML(errs ...*sperror.Error) ([]byte, error) {
	return exp(Xml, errs...)
}

func exp(_type string, errs ...*sperror.Error) ([]byte, error) {
	var arr []Error
	for _, e := range errs {
		err := Error{}
		err.Source = e.Source()
		err.Msg = e.Msg(sperror.En)
		err.Desc = e.Desc()
		err.Hint = e.Hint()
		if e.Caused() != nil {
			err.Cause = e.Caused().Error()
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
