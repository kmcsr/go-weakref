
# Weakref

Feature
=======

- **Weakref** for ***golang***
- Support weakref **map** _(weak value)_
- Check if value available

Usage
=====

type WeakRef
------------

A weakref instance

### func WeakRef.Ok

```golang
func (*WeakRef)Ok()(bool)
```

Check if this weakref is available

### func WeakRef.Value

```golang
func (*WeakRef)Value()(interface{})
```

Get the value of the weakref


func NewWeakRef
---------------

```golang
func NewWeakRef(ptr interface{})(w *WeakRef)
```

Create a instance of `weakref.WeakRef`, `ptr` must be a type of pointer or `unsafe.Pointer`

type Map
--------

A weak value map, if the value is not available, the map will auto remove the key  
The map is ***theard safty***

### func Map.Len

```golang
func (*Map)Len()()
```

### func Map.Has

```golang
func (*Map)Has(k interface{})(bool)
```

Check if key-value exists

### func Map.Get

```golang
func (*Map)Get(k interface{})(interface{})
```

Get value of key

### func Map.Set

```golang
func (*Map)Set(k interface{}, v interface{})(interface{})
```

Set value of key, and return the value

### func Map.GetOrSet

```golang
func (*Map)GetOrSet(k interface{}, setter func()(interface{}))(interface{})
```

Get value of key, if key not exists, set the value returned by `setter`, and return the value

### func Map.Pop

```golang
func (*Map)Pop(k interface{})(interface{})
```

Remove value of the key, and return the value

### func Map.Reset

```golang
func (*Map)Reset()
```

Reset the map _(clear all of the key-value)_

### func Map.AsMap

```golang
func (*Map)AsMap()(map[interface{}]interface{})
```

Create a new `builtin.map`, and copy all the key-value from `weakref.Map`

func NewMap
-----------

```golang
func NewMap()(*Map)
```

Create a new empty `weakref.Map`


Example for `weakref.WeakRef`:
------------------------------
```golang

import (
	weakref "github.com/zyxkad/weakref"
)

type T struct{
	id int
	name string
}

var global_cache_t *weakref.WeakRef

func GetT()(*T){
	if global_cache_t.Ok(){ // If this weakref is available
		return global_cache_t.Value().(*T) // type assertion
	}
	t := &T{
		id: 1,
		name: "name",
	}

	global_cache_t = weakref.NewWeakRef(t) // create a new weakref

	return t
}

// ...

func main(){
	for {
		// ...
		if true {
			t := GetT()
			do_something_with_t(t)
		}
		// ...
	}
}

```

Example for `weakref.Map`:
--------------------------
```golang

import (
	weakref "github.com/zyxkad/weakref"
)

type User struct{
	name string
	score int
}

var weak_cache *weakref.Map = weakref.NewMap()

func ReadUserFromFile(name string)(u *User){
	// ...
	return
}

func GetUser(name string)(*User){
	return weak_cache.GetOrSet(name, func()(interface{}){
		return ReadUserFromFile(name)
	}).(*User)
}


func main(){
	for {
		// ...
		{
			a := GetUser("a")
			a.score++
			SaveUser(a)
		}
		{
			b := GetUser("b")
			b.score--
			SaveUser(b)
		}
		// ...
}

```
