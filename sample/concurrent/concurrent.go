//go:generate mockgen -destination mock/concurrent_mock.go github.com/golang/mock/sample/concurrent Math

// Package concurrent demonstrates how to use gomock with goroutines.
package concurrent

type Math interface {	//要mock的接口
	Sum(a, b int) int
}
