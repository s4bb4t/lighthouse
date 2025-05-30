package sp

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"github.com/s4bb4t/lighthouse/core/levels"
)

func (e *Error) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	e.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

func (e *Error) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')
	w.RawString(`"messages":`)
	if e.messages == nil {
		w.RawString(`null`)
	} else {
		w.RawByte('{')
		v1First := true
		for v1Name, v1Value := range e.messages {
			if !v1First {
				w.RawByte(',')
			}
			v1First = false
			w.String(v1Name)
			w.RawByte(':')
			w.String(v1Value)
		}
		w.RawByte('}')
	}
	w.RawByte(',')
	w.RawString(`"description":`)
	w.String(e.desc)
	w.RawByte(',')
	w.RawString(`"hint":`)
	w.String(e.hint)
	w.RawByte(',')
	w.RawString(`"source":`)
	w.String(e.source)

	if e.httpCode != 0 {
		w.RawByte(',')
		w.RawString(`"http_code":`)
		w.Int(e.httpCode)
	}

	w.RawByte(',')
	w.RawString(`"level":`)
	w.Int(int(e.level))

	if len(e.meta) != 0 {
		w.RawByte(',')
		w.RawString(`"meta":`)
		if e.meta == nil {
			w.RawString(`null`)
		} else {
			w.RawByte('{')
			v2First := true
			for v2Name, v2Value := range e.meta {
				if !v2First {
					w.RawByte(',')
				}
				v2First = false
				w.String(v2Name)
				w.RawByte(':')
				w.Raw(json.Marshal(v2Value))
			}
			w.RawByte('}')
		}
	}

	w.RawByte('}')
}

func (e *Error) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	e.UnmarshalEasyJSON(&r)
	return r.Error()
}

func (e *Error) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "messages":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if e.messages == nil {
					e.messages = make(map[string]string)
				}
				for !in.IsDelim('}') {
					key := in.String()
					in.WantColon()
					var v1 string
					v1 = in.String()
					(e.messages)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "description":
			e.desc = in.String()
		case "hint":
			e.hint = in.String()
		case "source":
			e.source = in.String()
		case "http_code":
			e.httpCode = in.Int()
		case "level":
			e.level = levels.Level(in.Int())
		case "meta":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if e.meta == nil {
					e.meta = make(map[string]interface{})
				}
				for !in.IsDelim('}') {
					key := in.String()
					in.WantColon()
					var v2 interface{}
					if m, ok := v2.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v2.(json.Unmarshaler); ok {
						in.AddError(m.UnmarshalJSON(in.Raw()))
					} else {
						v2 = in.Interface()
					}
					(e.meta)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
