package system

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/pathvar"
)

func GetRouteAddr(r *http.Request) string {
	uri := r.URL.Path

	if !strings.HasSuffix(uri, "/") {
		uri += "/"
	}

	for key, val := range pathvar.Vars(r) {
		uri = strings.Replace(uri, "/"+val+"/", "/:"+key+"/", 1)
	}

	return strings.TrimRight(uri, "/")
}
