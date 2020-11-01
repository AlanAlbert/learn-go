package objpool

import (
	"errors"
	"time"
)

// Obj any data type
type Obj struct{}

// ObjPool pool of object
type ObjPool struct {
	pool chan *Obj
}

// NewObjPool new a object pool
func NewObjPool(size int) *ObjPool {
	objPool := ObjPool{}
	objPool.pool = make(chan *Obj, size)
	for i := 0; i < size; i++ {
		objPool.pool <- &Obj{}
	}

	return &objPool
}

// GetObj get a object from pool
func (objPool *ObjPool) GetObj(timeout time.Duration) (*Obj, error) {
	select {
	case obj := <-objPool.pool:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

// ReleaseObj put object into object pool
func (objPool *ObjPool) ReleaseObj(obj *Obj) error {
	select {
	case objPool.pool <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
