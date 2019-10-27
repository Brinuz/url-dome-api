package router

import "net/http"

// Router structure
type Router struct {
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Handler is the router main handler
func (r Router) Handler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", healthCheck)

	return mux
}
