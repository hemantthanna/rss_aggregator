package main

import (
	"database/sql"
	// "fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/hemantthanna/rss_aggregator/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

 func main()  {
	// feed, err := urlToFeed("https://wagslane.dev/index.xml")
	// if err != nil {
	// 	log.Fatal(err)		
	// }
	// fmt.Println(feed)

	/// pull environment variables into the program
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	connection , err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
		
	}

	db := database.New(connection)
	apiCfg := apiConfig{
		DB: db,
	}


	go startScarping(
		db, 10, time.Minute,
	)

	router :=  chi.NewRouter()

	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"https://*","http://*"},
			AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
			AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{"Link"},
			AllowCredentials: false,
			MaxAge: 300,
		},
	))

	// TO CHECK IF SERVER IS HEALTH AND LIVE. 
	// URL WILL BE /v1/healthz
	// this is so that in future we can rollout another server while this one still working
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handleRediness)
	v1Router.Get("/error", handleError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	
	router.Mount("/v1",v1Router)


	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}


	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		
	}


 }