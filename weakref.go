
package weakref

type WeakRef struct{
	a bool
	p IPointer
}

func NewWeakRef(ptr interface{})(w *WeakRef){
	if ptr == nil {
		panic("ptr cannot be nil")
	}
	w = &WeakRef{
		a: true,
	}
	w.p = SetFinalizer(ptr, func(interface{}){
		w.a = false
	})
	return
}

func (w *WeakRef)Ok()(bool){
	if w == nil {
		return false
	}
	return w.a
}

func (w *WeakRef)Value()(interface{}){
	if w == nil || !w.a {
		return nil
	}
	return iPointerToPtr(w.p)
}
