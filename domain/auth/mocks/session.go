package mocks

import (
	"context"
	"net/http"

	"github.com/go-session/session"
)

type MockSessionManager struct {
	DeleteMock  func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	RefreshMock func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error)
	StartMock   func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error)
	NeedsRefreshMock func(session session.Store) bool
}

func (m *MockSessionManager) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return m.DeleteMock(ctx, w, r)
}

func (m *MockSessionManager) Refresh(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
	return m.RefreshMock(ctx, w, r)
}

func (m *MockSessionManager) Start(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
	return m.StartMock(ctx, w, r)
}

func (m *MockSessionManager) NeedsRefresh(session session.Store) bool {
	return m.NeedsRefreshMock(session)
}

type MockSession struct {
	ContextMock   func() context.Context
	SessionIDMock func() string
	SetMock       func(key string, value interface{})
	GetMock       func(key string) (interface{}, bool)
	DeleteMock    func(key string) interface{}
	SaveMock      func() error
	FlushMock     func() error
}

func (m *MockSession) Context() context.Context {
	return m.ContextMock()
}

func (m *MockSession) SessionID() string {
	return m.SessionIDMock()
}

func (m *MockSession) Set(key string, value interface{}) {
	m.SetMock(key, value)
}

func (m *MockSession) Get(key string) (interface{}, bool) {
	return m.GetMock(key)
}

func (m *MockSession) Delete(key string) interface{} {
	return m.DeleteMock(key)
}

func (m *MockSession) Save() error {
	return m.SaveMock()
}

func (m *MockSession) Flush() error {
	return m.FlushMock()
}
