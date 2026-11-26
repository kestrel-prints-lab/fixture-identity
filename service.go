// Generated from testenv/topology/topology.yaml by testenv/topology/generate.js. Do not edit by hand.
package main

// Service holds the business logic of identity.
type Service struct {
	store   *Store
	clients *Clients
}

// NewService wires the store and the outbound clients.
func NewService(store *Store, clients *Clients) *Service {
	return &Service{store: store, clients: clients}
}

func str(m map[string]any, k, def string) string {
	if v, ok := m[k].(string); ok && v != "" {
		return v
	}
	return def
}

// validateInput rejects empty references.
func validateInput(v string) string {
	if v == "" {
		return "unknown"
	}
	return v
}

// IssueToken is the business logic behind POST /oauth/token.
func (s *Service) IssueToken(body map[string]any) map[string]any {
	key := str(body, "id", "user-001")
	record := map[string]any{"id": key}
	record["ref"] = validateInput(key)
	s.store.SaveUser(key, record)
	return record
}

// Introspect is the business logic behind POST /oauth/introspect.
func (s *Service) Introspect(body map[string]any) map[string]any {
	key := str(body, "id", "user-001")
	record := s.store.GetUser(key)
	return record
}

// GetUser is the business logic behind GET /v1/users/{user_id}.
func (s *Service) GetUser(userID string) map[string]any {
	record := s.store.GetUser(userID)
	return record
}

// UserRisk is the business logic behind GET /v1/users/{user_id}/risk.
func (s *Service) UserRisk(userID string) map[string]any {
	record := s.store.GetUser(userID)
	return record
}
