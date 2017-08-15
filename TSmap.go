package TSmap

import "sync"

type TSmap interface {
	Set(k,v interface{})
	Get(k interface{}) (interface{},bool)
	Delete(k interface{})
	ForEach(f func(k,v interface{}))
	GoForEach(f func(k,v interface{}))
	Len() int
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

func (this *NewTSmap)ForEach(f func(k,v interface{})) {
	for key := range this.ConMap {
		_k,_ := this.Get(key)
		f(key,_k)
	}
}

func (this *NewTSmap)GoForEach(f func(k,v interface{})) {
	for key := range this.ConMap {
		go func() {
			_k,_ := this.Get(key)
			f(key,_k)
		}()
	}
}

func (this *NewTSmap)Len() int {
	return len(this.ConMap)
}