package web

import "net/http"

func renderHome(w http.ResponseWriter, r *http.Request) {
	out, err := TMPLERRhome("Hydrocarbon", false, 0)
	if err != nil {
		panic(err)
	}

	_, err = w.Write([]byte(out))
	if err != nil {
		panic(err)
	}
}
