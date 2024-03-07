package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/impzero/roosh/roosh"
	"github.com/pushmarket/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

type API struct {
	Router *chi.Mux
	server *http.Server

	Config
}

type Config struct {
	Addr string

	Stack      *roosh.Stack
	Calculator roosh.Calculator
}

func New(c Config) *API {
	r := chi.NewRouter()
	api := &API{
		Router: r,
		Config: c,
	}

	render.Decode = api.Decoder

	r.Route("/", func(r chi.Router) {
		r.Route("/calc", func(r chi.Router) {
			r.Post("/add", api.handleCalcAdd)
			r.Post("/div", api.handleCalcDiv)
			r.Post("/multip", api.handleCalcMultip)
			r.Post("/subs", api.handleCalcSubs)
		})
		r.Route("/stack", func(r chi.Router) {
			r.Get("/", api.handleStackFetch)
			r.Post("/", api.handleStackAdd)
			r.Delete("/", api.handleStackPop)
		})
		r.Mount("/swagger", httpSwagger.WrapHandler)
	})

	api.server = &http.Server{
		Addr:    c.Addr,
		Handler: api.Router,
	}
	return api
}

func (a *API) Start() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen, addr: %s\n", a.Config.Addr)
		}
	}()

	log.Printf("🤖 Roosh Server is listening, address: %s\n", a.Config.Addr)
	a.gracefulShutdown()
}

func (a *API) gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Printf("Received signal: %s, shutting down\n", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	a.server.SetKeepAlivesEnabled(false)
	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatalln("Failed to gracefully shutdown")
	}
	log.Println("Roosh Server stopped")
}
