package main

import (
	"fmt"
	"net/http"
)

func h(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s: Вы вызвали %s методом %s\n", name, r.URL.String(), r.Method)
	}
}

func main() {
	cf := config.NewConfig()
	dbConnect := db.NewDb(cf.Db)
	rabbitConnect := rabbitmq.New(cf.RabbitMQ)
	redisConnect := redis.New(cf.Redis)

	m := http.NewServeMux()
	m.Handle("GET /posts/latest", h("latest"))
	m.Handle("GET /posts/{id}", h("id"))
	m.Handle("GET /posts", h("posts"))
	http.ListenAndServe(":80", m)
}
