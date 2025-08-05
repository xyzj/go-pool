# go-pool

`go-pool` 是一个基于泛型的高性能对象池，支持自定义对象的创建和最大空闲池容量限制。适用于高并发场景下对象复用，减少 GC 压力，提高性能。

## 功能介绍

- 支持任意类型的对象池化管理（基于 Go 1.18+ 泛型）
- 最大空闲池容量可配置，超出部分自动交由 `sync.Pool` 管理
- 并发安全，适合多 goroutine 环境
- 支持自定义对象的创建逻辑

## 安装

```sh
go get github.com/yourusername/go-pool
```

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/yourusername/go-pool"
)

func main() {
	// 创建一个 int 类型的对象池，最大空闲池容量为 5
	pool := gopool.New(func() int {
		return 0 // 对象的创建逻辑
	}, gopool.WithMaxPoolSize(5))

	// 从池中获取对象
	obj := pool.Get()
	fmt.Println("Get from pool:", obj)

	// 使用完毕后归还对象
	pool.Put(obj)
}
```

## API 说明

### 创建对象池

```go
pool := gopool.New(func() T { ... }, gopool.WithMaxPoolSize(n))
```
- `func() T`：对象的创建函数
- `WithMaxPoolSize(n)`：可选，设置最大空闲池容量，默认 10
- `WithWarmCount(n)`：可选，设置池内对象预热数量，默认 0

### 获取对象

```go
obj := pool.Get()
```

### 归还对象

```go
pool.Put(obj)
```

## 线程安全性

本对象池所有方法均为并发安全，可放心在多 goroutine 场景下使用。

## License

GNU 3 License

---

如需更详细的用法或扩展功能，请查阅源码或提交 issue。