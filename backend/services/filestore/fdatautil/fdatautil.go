package fdatautil

import (
	"net/http"
	"strings"

	"github.com/bornjre/turnix/backend/xtypes/services/xfilestore"
)

func ReadAndClose(data xfilestore.FData) ([]byte, error) {
	defer data.Close()

	out, err := data.AsBytes()
	return out, err
}

const (
	CtypeJSON = "application/json"
	CtypeJS   = "application/javascript"
	CtypeBin  = "application/octet-stream"
	CtypeCSS  = "text/css"
)

func WriteAndClose(rw http.ResponseWriter, file string, data xfilestore.FData) {
	defer data.Close()

	out, err := data.AsBytes()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ffiles := strings.Split(file, ".")

	ctype := ""
	switch ffiles[1] {
	case "js":
		ctype = CtypeJS
	case "css":
		ctype = CtypeCSS
	default:
		ctype = http.DetectContentType(out)
	}

	rw.Header().Set("Context-Type", ctype)
	rw.Write(out)

}
