// Generated from testenv/topology/topology.yaml by testenv/topology/generate.js. Do not edit by hand.
package main

import "sync"

// Store keeps user records in memory.
type Store struct {
	mu      sync.RWMutex
	records map[string]map[string]any
}

// NewStore returns an empty store.
func NewStore() *Store {
	return &Store{records: map[string]map[string]any{}}
}

func (s *Store) GetUser(key string) map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if r, ok := s.records[key]; ok {
		return r
	}
	return map[string]any{"id": key}
}

func (s *Store) SaveUser(key string, record map[string]any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.records[key] = record
}

func (s *Store) ListUserRecords() []map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]map[string]any, 0, len(s.records))
	for _, r := range s.records {
		out = append(out, r)
	}
	return out
}

func (s *Store) DeleteUser(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.records, key)
}
