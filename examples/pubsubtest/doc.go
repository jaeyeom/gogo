// Package pubsubtest demonstrates testing pubsub Publish method using
// interfaces. There are 2 ways to test it.
//
// 1. Define an interface that returns an interface and wrap topic with a
// wrapper that implements the interface. (interface.go and interface_test.go)
// 2. Implement synchronous version and use it for asynchronous call.
// (gochan.go and gochan_test.go)
package pubsubtest
