
package weakref

import (
	sync "sync"
)

type Map struct{
	mux sync.RWMutex
	m map[interface{}]IPointer
}

func NewMap()(*Map){
	return &Map{
		m: make(map[interface{}]IPointer),
	}
}

func (m *Map)Len()(int){
	m.mux.RLock()
	defer m.mux.RUnlock()

	return len(m.m)
}

func (m *Map)Has(k interface{})(ok bool){
	m.mux.RLock()
	defer m.mux.RUnlock()

	_, ok = m.m[k]
	return
}

func (m *Map)Get(k interface{})(interface{}){
	m.mux.RLock()
	defer m.mux.RUnlock()

	v, ok := m.m[k]
	if !ok {
		return nil
	}
	return iPointerToPtr(v)
}

func (m *Map)Set(k interface{}, ptr interface{})(interface{}){
	m.mux.Lock()
	defer m.mux.Unlock()

	m.setLocked(k, ptr)
	return ptr
}

func (m *Map)setLocked(k interface{}, ptr interface{}){
	if ptr == nil {
		panic("ptr cannot be nil")
	}
	m.m[k] = SetFinalizer(ptr, func(interface{}){
		m.Pop(k)
	})
	return
}

func (m *Map)GetOrSet(k interface{}, s func()(interface{}))(v interface{}){
	m.mux.RLock()
	p, ok := m.m[k]
	m.mux.RUnlock()
	if ok {
		v = iPointerToPtr(p)
	}else{
		v = m.Set(k, s())
	}
	return
}

func (m *Map)Pop(k interface{})(interface{}){
	m.mux.Lock()
	defer m.mux.Unlock()

	v, ok := m.m[k]
	if ok {
		delete(m.m, k)
		return iPointerToPtr(v)
	}
	return nil
}

func (m *Map)Reset(){
	m.mux.Lock()
	defer m.mux.Unlock()

	m.m = make(map[interface{}]IPointer)
}

func (m *Map)AsMap()(p map[interface{}]interface{}){
	m.mux.RLock()
	defer m.mux.RUnlock()

	p = make(map[interface{}]interface{}, len(m.m))
	for k, v := range m.m {
		p[k] = iPointerToPtr(v)
	}
	return
}
