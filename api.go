package go_testing_baru

import (
	"encoding/json"
	"net/http"
)

type UserDB interface {
	FindByID(ID string) (string, error)
}

func getUser(db UserDB) http.HandlerFunc {
	//injection
	return func(w http.ResponseWriter, r *http.Request) {
		type user struct {
			Name string `json:'name'`
		}
		name, err := db.FindByID("1")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		u := user{
			Name: name,
		}

		b, _ := json.Marshal(u)
		w.Write(b)
		return
	}
}
