package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"learning-go/examples/ch12/ex04-context-uuid/identity"
	"learning-go/examples/ch12/ex04-context-uuid/tracker"
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	bl := BusinessLogic {
		RequestDecorator: tracker.Request,
		Logger: tracker.Logger{},
		Remote: "http://www.example.com/query",
	}

	c := Controller {
		Logic: bl,
	}

	r := chi.NewRouter()
	r.Use(tracker.Middleware, identity.Middleware)
	r.Get("/", c.handleRequest)
	http.ListenAndServe(":3000", r)
}

type RequestDecorator func(*http.Request) *http.Request

type Logger interface {
	Log(context.Context, string)
}

type BusinessLogic struct {
	RequestDecorator RequestDecorator
	Logger Logger
	Remote string
}

func (bl BusinessLogic) businessLogic(ctx context.Context, user string, data string) (string, error) {
	bl.Logger.Log(ctx, fmt.Sprint("starting businessLogic for %s with %s", user, data))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, bl.Remote+"?query="+data, nil)
	if err != nil {
		bl.Logger.Log(ctx, fmt.Sprintf("error building remote request %v", err))
		return "", err
	}

	req = bl.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		bl.Logger.Log(ctx, fmt.Sprintf("error making remote request %v", err))
		return "", err
	}

	body := resp.Body
	defer body.Close()

	out, err := ioutil.ReadAll(body)
	if err != nil {
		bl.Logger.Log(ctx, fmt.Sprintf("error reading response: %v", err))
		return "", err
	}

	return string(out), nil
}

type Logic interface {
	businessLogic(ctx context.Context, user string, data string) (string, error) 
}

type Controller struct {
	Logic Logic
}

func (c Controller) handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := identity.UserFromContext(ctx)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := r.URL.Query().Get("data")
	result, err := c.Logic.businessLogic(ctx, user, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(result))
}
