package TSmap

import (
	"testing"
	"fmt"
)

func testAssert(t *testing.T, b bool, objs ...interface{}) {
	if !b {
		t.Fatal(objs...)
	}
}

func TestBasic(t *testing.T)  {
	tsmap := NewTSmap{
		ConMap:  make(map[interface{}]interface{}),
	}
	go func() {
		for i := 0; i <= 1000; i++ {
			go tsmap.Set(fmt.Sprintf("%v",i),fmt.Sprintf("counter: '%v'",i))
		}
	}()
	go func() {
		for i := 0; i <= 1000; i++ {
			tsmap.Get(fmt.Sprintf("%v",i))
		}
	}()
}