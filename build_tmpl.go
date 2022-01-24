package gcg

var buildSrcTemplate = `package {{.Package}}
// generated by github.com/emicklei/graphql-client-gen/cmd/gcg on {{.Created.Format "Mon, 02 Jan 2006 15:04:05 MST" }} 
// DO NOT EDIT

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func BuildQuery(q interface{}) string {
	b := new(bytes.Buffer)
	writeQuery(q, b, 0, false)
	return b.String()
}

func writeQuery(q interface{}, w io.Writer, indent int, inline bool) {
	fmt.Printf("%T\n", q)
	rt := reflect.TypeOf(q)
	rv := reflect.ValueOf(q)
	if rt.Kind() == reflect.Pointer {
		writeQuery(rv.Elem().Interface(), w, indent, inline)
		return
	}
	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		if !fv.IsZero() {
			sf := rt.Field(i)
			fmt.Println("into field:", sf.Name)
			tag, ok := sf.Tag.Lookup("graphql")
			inlineField := sf.Anonymous && !ok
			if inlineField {
				fmt.Println("inline")
				// is struct or pointer to struct
				k := sf.Type
				if k.Kind() == reflect.Pointer {
					k = k.Elem()
					fv = fv.Elem()
				}
				if k.Kind() == reflect.Struct {
					writeStruct(fv.Interface(), w, indent, inlineField)
				}				
			} else {
				// no inline
				if ok {
					fmt.Fprintf(w, "\t%s", tag)
					// is struct or pointer to struct
					k := sf.Type
					if k.Kind() == reflect.Pointer {
						k = k.Elem()
						fv = fv.Elem()
					}
					if k.Kind() == reflect.Struct {
						writeStruct(fv.Interface(), w, indent, inlineField)
					} else if k.Kind() == reflect.Slice {
						// take first if avail
						if fv.Len() > 0 {
							// always struct? TODO
							writeStruct(fv.Index(0).Interface(),w,indent,inline)
						}
					} else {
						io.WriteString(w, "\n")
					}
					io.WriteString(w, strings.Repeat("\t", indent))
				}
			}
		}
	}
}

func writeStruct(structValue interface{}, w io.Writer, indent int, inline bool) {
	io.WriteString(w, " {\n")
	io.WriteString(w, strings.Repeat("\t", indent+1))
	writeQuery(structValue, w, indent+1, inline)
	io.WriteString(w, "}\n")
}
`