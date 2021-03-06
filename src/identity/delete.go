package identity

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

// Delete remove the identity
func (i Identity) Delete() (Identity, error) {
	if i.ID != "" && i.Email != "" && i.Phone != "" {
		getEntry, err := i.Retrieve()
		if err != nil {
			return Identity{}, err
		}
		resp, err := getEntry.DeleteEntry()
		if err != nil {
			return Identity{}, err
		}

		return resp, err
	}

	return Identity{}, errors.New("you need to have all the details to delete")
}

// Delete http remove the identity
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "identityId")

	i := Identity{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	i.ID = id
	resp, err := i.Delete()
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(Response{
		Identity: resp,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("delete response err: %v", err))
	}
}
