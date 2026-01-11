package req

import (
	"net/http"
	"validation/api/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JSON(*w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.JSON(*w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return nil, err
	}
	return &body, nil
}
