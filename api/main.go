package main

import (
	"context"
	_ "ifood_case/api/docs"
	"ifood_case/api/handlers"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var ctx = context.Background()
var rdb *redis.Client

// @title Prompt API
// @version 1.0
// @description Uma api que processar prompts, chama um script python que processa o prompt com o modelo treinado e retorna a resposta.
// @host localhost:8080
// @BasePath /
func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	rdb = redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Printf("Connected to Redis: %s", pong)
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/prompt", handlers.PromptHandler).Methods("POST")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func cacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cacheKey := "prompt:" + r.URL.String()

		cachedResponse, err := rdb.Get(ctx, cacheKey).Result()
		if err == redis.Nil {

			next.ServeHTTP(w, r)

		} else if err != nil {
			http.Error(w, "Redis error", http.StatusInternalServerError)
			return
		} else {

			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cachedResponse))
		}
	}
}
