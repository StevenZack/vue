package vue

import "net/http"

func HandleJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(Bytes_VueMinJs)
}
