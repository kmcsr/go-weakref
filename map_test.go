
package weakref_test

import (
	"testing"
	"time"
	"runtime"
	weakref "github.com/kmcsr/go-weakref"
)

func TestMap(t *testing.T){
	type T struct{ byte }
	m := weakref.NewMap()
	m.Set(1, &T{})
	m.GetOrSet(2, func()(interface{}){
		var t T
		return &t
	})
	t.Log("map:", m)
	runtime.GC()
	t.Log("=== After GC")
	t.Log("map:", m)
}

func TestSyncMap(t *testing.T){
	type T struct{ int }
	m := weakref.NewMap()
	for i := 0; i < 1000; i++ {
		v := &T{i}
		m.Set(i, v)
		j := i
		go func() {
			time.Sleep(time.Millisecond * 20)
			m.Pop(j)
		}()
	}
	time.Sleep(time.Millisecond * 40)
}

func TestForEachMap(t *testing.T){
	type T struct{ int }
	m := weakref.NewMap()
	for i := 0; i < 20; i++ {
		v := &T{i}
		m.Set(i, v)
	}
	for k, v := range m.AsMap() {
		_, _ = k, v.(*T)
	}
}
