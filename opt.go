package gopool

type PoolOpt struct {
	maxIdleSize int
	warmCount   int
}
type PoolOpts func(opt *PoolOpt)

// WithMaxPoolSize returns a PoolOpts function that sets the maximum number of idle
// resources in the pool to the specified value t, with a minimum enforced value of 2.
func WithMaxPoolSize(t int) PoolOpts {
	return func(o *PoolOpt) {
		o.maxIdleSize = max(2, t)
	}
}

// WithWarmCount returns a PoolOpts function that sets the warmCount option of the pool to the specified value t,
// ensuring that it is not less than zero. This determines the number of worker goroutines to pre-initialize in the pool.
func WithWarmCount(t int) PoolOpts {
	return func(o *PoolOpt) {
		o.warmCount = max(t, 0)
	}
}
