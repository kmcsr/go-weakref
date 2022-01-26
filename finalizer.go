
package weakref

import (
	fmt "fmt"
	unsafe "unsafe"
	reflect "reflect"
	runtime "runtime"
	sync "sync"
)

type Pointer = unsafe.Pointer
type IPointer = [2]uintptr

func ptrToIPointer(ptr interface{})(IPointer){
	return *(*IPointer)((Pointer)(&ptr))
}

func iPointerToPtr(ptr IPointer)(interface{}){
	return *(*interface{})((Pointer)(&ptr))
}

var finalizer_map_mux sync.RWMutex
var finalizer_map = make(map[IPointer][]func(interface{}))

func SetFinalizer(ptr interface{}, cb func(interface{}))(iptr IPointer){
	v := reflect.ValueOf(ptr)
	kd := v.Kind()
	if kd != reflect.Ptr && kd != reflect.UnsafePointer {
		panic(fmt.Errorf("ptr must be a type of pointer or unsafe.Pointer, not %v", kd))
	}
	iptr = ptrToIPointer(ptr)
	finalizer_map_mux.RLock()
	s, ok := finalizer_map[iptr]
	finalizer_map_mux.RUnlock()
	if !ok {
		runtime.SetFinalizer((*struct{})((Pointer)(iptr[1])), func(*struct{}){
			p := iPointerToPtr(iptr)
			finalizer_map_mux.Lock()
			s := finalizer_map[iptr]
			delete(finalizer_map, iptr)
			finalizer_map_mux.Unlock()
			for _, cb := range s {
				cb(p)
			}
		})
	}
	finalizer_map_mux.Lock()
	finalizer_map[iptr] = append(s, cb)
	finalizer_map_mux.Unlock()
	return
}
