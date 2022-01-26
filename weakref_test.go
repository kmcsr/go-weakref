
package weakref_test

import (
	"testing"
	"runtime"
	weakref "github.com/kmcsr/go-weakref"
)

func TestWeakRefNil(t *testing.T){
	var ref *weakref.WeakRef
	if ref.Ok() {
		t.Fatal("weakref.WeakRef(nil).Ok() must be false")
	}
	if ref.Value() != nil {
		t.Fatal("weakref.WeakRef(nil).Value() must be nil")
	}
}

func TestWeakRef(t *testing.T){
	type T struct{ byte }
	v := &T{}
	ref := weakref.NewWeakRef((interface{})(v))
	t.Log("ref:", ref)
	t.Log("Alive:", ref.Ok())
	if ref.Ok() {
		t.Log("Value:", (ref.Value()).(*T))
	}
	runtime.GC()
	t.Log("=== After GC")
	t.Log("Alive:", ref.Ok())
	if ref.Ok() {
		t.Log("Value:", (ref.Value()).(*T))
	}
}
