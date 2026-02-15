package req

import (
	"go/adv-example/pkg/res"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := DecodeBody[T](req.Body)
	if err != nil {
		return nil, err
	}

	err = Validate(body)
	if err != nil {
		res.Json(w, body, 402)
		return nil, err
	}

	return &body, nil
}
