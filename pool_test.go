package gopool

import (
	"testing"
)

// Helper type for testing
type testType struct {
	val int
}

func TestGoPool_Get_FromIdle(t *testing.T) {
	pool := New(func() testType { return testType{val: 42} })
	item := testType{val: 99}
	pool.idle <- item

	got := pool.Get()
	if got.val != 99 {
		t.Errorf("expected value from idle channel, got %v", got.val)
	}
}

func TestGoPool_Get_FromSyncPool(t *testing.T) {
	pool := New(func() testType { return testType{val: 42} })
	item := testType{val: 77}
	pool.pool.Put(item)

	got := pool.Get()
	if got.val != 77 {
		t.Errorf("expected value from sync.Pool, got %v", got.val)
	}
}

func TestGoPool_Get_NewFunc(t *testing.T) {
	pool := New(func() testType { return testType{val: 123} })

	got := pool.Get()
	if got.val != 123 {
		t.Errorf("expected value from new func, got %v", got.val)
	}
}
