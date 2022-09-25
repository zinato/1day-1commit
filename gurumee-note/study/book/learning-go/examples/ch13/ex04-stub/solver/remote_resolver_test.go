package solver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResolve(t *testing.T) {
	type info struct {
		expr string
		code int
		body string
	}

	var io info
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expr := r.URL.Query().Get("expression")
		if expr != io.expr {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid expression: " + io.expr))
			return
		}

		w.WriteHeader(io.code)
		w.Write([]byte(io.body))
	}))
	defer server.Close()

	rs := RemoteResolver{
		MathServerURL: server.URL,
		Client:        server.Client(),
	}

	data := []struct {
		name   string
		io     info
		result float64
		errMsg string
	}{
		{"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22, ""},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := rs.Resolve(context.Background(), d.io.expr)
			if result != d.result {
				t.Errorf("io %f, got %f", d.result, result)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if errMsg != d.errMsg {
				t.Errorf("io error %s, got %s", d.errMsg, errMsg)
			}
		})
	}
}
