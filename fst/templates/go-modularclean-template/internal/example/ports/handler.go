package ports

import "net/http"

type IExampleHandler interface {
	CreateExample(w http.ResponseWriter, r *http.Request)
}
