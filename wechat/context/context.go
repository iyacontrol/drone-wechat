package context

import (
	"encoding/json"
	"net/http"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

// Context struct
type Context struct {
	CorpID         string
	CorpSecret     string
	Token          string
	EncodingAESKey string

	Writer  http.ResponseWriter
	Request *http.Request
}

// Query returns the keyed url query value if it exists
func (ctx *Context) Query(key string) string {
	value, _ := ctx.GetQuery(key)
	return value
}

// GetQuery is like Query(), it returns the keyed url query value
func (ctx *Context) GetQuery(key string) (string, bool) {
	req := ctx.Request
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

//Render render from bytes
func (ctx *Context) Render(bytes []byte) {
	//debug
	//fmt.Println("response msg = ", string(bytes))
	ctx.Writer.WriteHeader(200)
	_, err := ctx.Writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}

//JSON render to json
func (ctx *Context) JSON(obj interface{}) {
	writeContextType(ctx.Writer, jsonContentType)
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	ctx.Render(bytes)
}

func writeContextType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
