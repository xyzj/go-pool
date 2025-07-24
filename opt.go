package gopool

type PoolOpt struct {
	maxIdleSize int
}
type PoolOpts func(opt *PoolOpt)

// OptMaxIdleSize returns a PoolOpts function that sets the maximum number of idle
// resources in the pool to the specified value t, with a minimum enforced value of 2.
func OptMaxIdleSize(t int) PoolOpts {
	return func(o *PoolOpt) {
		o.maxIdleSize = max(2, t)
	}
}
