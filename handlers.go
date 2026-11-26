// Generated from testenv/topology/topology.yaml by testenv/topology/generate.js. Do not edit by hand.
package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handler serves the HTTP routes.
type Handler struct {
	svc *Service
}

func (h *Handler) register(r chi.Router) {
	r.Post("/oauth/token", h.issueToken)
	r.Post("/oauth/introspect", h.introspect)
	r.Get("/v1/users/{user_id}", h.getUser)
	r.Get("/v1/users/{user_id}/risk", h.userRisk)
	r.Get("/.well-known/jwks.json", h.jwks)
}

func readJSON(r *http.Request) map[string]any {
	out := map[string]any{}
	_ = json.NewDecoder(r.Body).Decode(&out)
	return out
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) healthz(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{"status": "ok"})
}

func (h *Handler) issueToken(w http.ResponseWriter, r *http.Request) {
	body := readJSON(r)
	result := h.svc.IssueToken(body)
	writeJSON(w, result)
}

func (h *Handler) introspect(w http.ResponseWriter, r *http.Request) {
	body := readJSON(r)
	result := h.svc.Introspect(body)
	writeJSON(w, result)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	result := h.svc.GetUser(userID)
	writeJSON(w, result)
}

func (h *Handler) userRisk(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	result := h.svc.UserRisk(userID)
	writeJSON(w, result)
}

func (h *Handler) jwks(w http.ResponseWriter, r *http.Request) {
	result := map[string]any{}
	writeJSON(w, result)
}
