package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/graph/generated"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/graph/resolver"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/repository"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/usecase"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // here
)

const defaultPort = "8080"

// EnvLoad ..
func EnvLoad() {
	if os.Getenv("GO_ENV") == "production" {
		return
	}
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	EnvLoad()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo, cleanup := repository.NewRepository()
	defer cleanup()

	uc := usecase.NewUsecase(repo)
	rs := resolver.NewResolver(uc)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: rs}))

	r := chi.NewRouter()
	// addCorsMiddleware(r)
	// addAuthMiddleware(r, repo.User())
	// addDataloaderMiddleiare(r, repo)

	r.Route("/graphql", func(r chi.Router) {
		r.Handle("/", srv)
	})

	gqlPlayground := playground.Handler("api-gateway", "/graphql")
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		msg := []byte("Hellow World")
		rw.Write(msg)
	})
	r.Get("/play", gqlPlayground)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
