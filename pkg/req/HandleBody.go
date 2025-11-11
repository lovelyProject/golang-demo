package pkg

import (
	"go/adv-example/pkg/res"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		return nil, err
	}

	err = Validate(body)
	if err != nil {
		res.Json(w, http.StatusBadRequest, body)
		return nil, err
	}

	return &body, nil
}
