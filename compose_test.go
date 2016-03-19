package mux

import (
	"github.com/nbio/st"
	"gopkg.in/vinci-proxy/vinci.v0/utils"
	"net/http"
	"net/url"
	"testing"
)

func TestMuxComposeIfMatches(t *testing.T) {
	mx := New()
	mx.Use(If(Method("GET"), Host("foo.com")).Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
		w.Header().Set("foo", "bar")
		h.ServeHTTP(w, r)
	}))

	wrt := utils.NewWriterStub()
	req := newRequest()
	req.URL.Host = "foo.com"

	mx.Middleware.Run("request", wrt, req, nil)
	st.Expect(t, wrt.Header().Get("foo"), "bar")
}

func TestMuxComposeIfUnmatch(t *testing.T) {
	mx := New()
	mx.Use(If(Method("GET"), Host("bar.com")).Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
		w.Header().Set("foo", "bar")
		h.ServeHTTP(w, r)
	}))

	wrt := utils.NewWriterStub()
	req := newRequest()
	req.URL.Host = "foo.com"

	mx.Middleware.Run("request", wrt, req, nil)
	st.Expect(t, wrt.Header().Get("foo"), "")
}

func TestMuxComposeOrMatch(t *testing.T) {
	mx := New()
	mx.Use(Or(Method("GET"), Host("bar.com")).Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
		w.Header().Set("foo", "bar")
		h.ServeHTTP(w, r)
	}))

	wrt := utils.NewWriterStub()
	req := newRequest()
	req.URL.Host = "foo.com"

	mx.Middleware.Run("request", wrt, req, nil)
	st.Expect(t, wrt.Header().Get("foo"), "bar")
}

func TestMuxComposeOrUnMatch(t *testing.T) {
	mx := New()
	mx.Use(Or(Method("GET"), Host("bar.com")).Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
		w.Header().Set("foo", "bar")
		h.ServeHTTP(w, r)
	}))

	wrt := utils.NewWriterStub()
	req := newRequest()
	req.URL.Host = "foo.com"

	mx.Middleware.Run("request", wrt, req, nil)
	st.Expect(t, wrt.Header().Get("foo"), "bar")
}

func newRequest() *http.Request {
	return &http.Request{URL: &url.URL{}, Header: make(http.Header), Method: "GET"}
}