package TSmap

import "sync"

type TSmap interface {
	Set(k,v interface{})
	Get(k interface{}) (interface{},bool)
	Delete(k interface{})
}

type NewTSmap struct{
	ConMap map[interface{}]interface{}
	lock sync.RWMutex
}

func (this *NewTSmap)Set(k,v interface{})  {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.ConMap[k] = v
}

func (this *NewTSmap)Get(k interface{}) (interface{},bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	v,ok := this.ConMap[k]
	return v,ok
}

func (this *NewTSmap)Delete(k interface{})  {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.ConMap,k)
}