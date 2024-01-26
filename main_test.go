package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//var done = make(chan struct{})
//var logged = make(chan struct{})
//
//func TestBadness(t *testing.T) {
//	t.Parallel()
//	go func() {
//		defer close(logged)
//		<-done
//		t.Log("BWA ha ha ha ha!")
//	}()
//	t.Log("Hmm?")
//}
//
//func TestMain(m *testing.M) {
//	code := m.Run()
//	close(done)
//	<-logged
//	os.Exit(code)
//}
//
//func TestPanic(t *testing.T) {
//	var inp *int
//	assert.Equal(t, 0, *inp)
//}

var x = 1
var y = 2

func Test1(t *testing.T) {
	time.Sleep(time.Second)
	x = 2
	assert.Equal(t, 2, x)
}

func Test2(t *testing.T) {
	y = 3
	assert.Equal(t, 1, x)
}

func Test3(t *testing.T) {
	assert.Equal(t, 3, y)
}
