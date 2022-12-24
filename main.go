package main

import (
	"blog-chi-gorm/config"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	config.SetupConfiguration()
}

func main() {
	r := chi.NewRouter()

	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	}

	db, err := config.DatabaseConnect()

	if err != nil {
		log.Fatal(err)
	}

	r.Use(middleware.Logger)

	r.Use(middlewares.MiddlewareCors)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	routes.NewAuthRoutes("/auth", db, r)
	routes.NewUserRoutes("/user", db, r)
	routes.NewArticleRoutes("/article", db, r)
	routes.NewCategoryRoutes("/category", db, r)
	routes.NewRoleRoutes("/role", db, r)
	routes.NewPermissionRoutes("/permission", db, r)
	routes.NewPostRoutes("/posts", db, r)
	routes.NewMenuRoutes("/menu", db, r)
	routes.NewSubMenuRoutes("/submenu", db, r)
	routes.NewTagRoutes("/tag", db, r)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Config.PORT),
		WriteTimeout: config.Config.WRITETIMEOUT * 10,
		ReadTimeout:  config.Config.READTIMEOUT * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		err := serve.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Connected to port:", config.Config.PORT)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)

}
