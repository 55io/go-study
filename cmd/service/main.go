package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/55io/go-study/pkg/config"
	"github.com/55io/go-study/pkg/logger"
)

func h(name string, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("%s: Вы вызвали %s методом %s\n", name, r.URL.String(), r.Method)
		log.Info(message)
		fmt.Fprintf(w, message)
	}
}

// func updateHandler(id string, db *interface{}) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		db.r(id, r.PostForm)
// 		fmt.Fprintf(w, "%s: Вы вызвали %s методом %s\n", id, r.URL.String(), r.Method)
// 	}
// }

func main() {
	cf := config.Get()
	log := logger.Get(cf.Env)

	//dbConnect := db.Get(cf.Db)

	m := http.NewServeMux()
	m.Handle("GET /posts/latest", h("latest", log))
	m.Handle("GET /posts/{id}", h("id", log))
	m.Handle("GET /posts", h("posts", log))
	http.ListenAndServe(":80", m)
}
