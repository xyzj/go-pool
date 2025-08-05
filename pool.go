package gopool

import (
	"sync"
)

type GoPool[T any] struct {
	mu   sync.Mutex
	opt  *PoolOpt
	idle chan T
	pool sync.Pool
	new  func() T
}

// Get retrieves an item of type T from the pool. It first attempts to obtain an idle item from the pool's idle channel.
// If no idle item is available, it tries to get an item from the underlying sync.Pool. If the pool returns nil,
// a new item is created using the pool's new function. The retrieved or newly created item is then returned.
func (g *GoPool[T]) Get() T {
	select {
	case x := <-g.idle:
		return x
	default:
		v := g.pool.Get()
		if v == nil {
			return g.new()
		}
		return v.(T)
	}
}

// Put adds an item of type T to the pool. If the number of idle items is less than the maximum pool size,
// the item is added to the idle channel. Otherwise, it is added to the underlying sync.Pool.
// This method is safe for concurrent use.
func (g *GoPool[T]) Put(i T) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if len(g.idle) < g.opt.maxIdleSize {
		g.idle <- i
		return
	}
	g.pool.Put(i)
}

// New creates and returns a new instance of GoPool for the specified type T.
// The 'new' parameter is a function that generates new instances of T.
// Optional PoolOpts can be provided to configure the pool's behavior, such as maximum pool size.
// Returns a pointer to the created GoPool[T].
func New[T any](new func() T, opts ...PoolOpts) *GoPool[T] {
	opt := &PoolOpt{
		maxIdleSize: 10,
		warmCount:   0,
	}
	for _, o := range opts {
		o(opt)
	}
	opt.warmCount = min(opt.maxIdleSize, opt.warmCount)
	ch := make(chan T, opt.maxIdleSize)
	if opt.warmCount > 0 {
		for range opt.warmCount {
			ch <- new()
		}
	}
	return &GoPool[T]{
		new: new,
		opt: opt,
		mu:  sync.Mutex{},
		pool: sync.Pool{
			New: func() any { return nil },
		},
		idle: ch,
	}
}
