package DcStore

import (
	"reflect"
	"testing"
)

func TestEs(t *testing.T) {
	ss := NewStaticStore("127.0.0.1", "9200")
	ti, tt, tm := "test-index", "test-type", "type-msgid"
	tv := map[string]interface{}{"name": "xxxxx", "age": "18"}

	var r map[string]interface{}
	ss.InsertDoc(ti, tt, tm, 30, tv)
	e := ss.GetDoc(ti, tt, tm, &r)
	t.Log("get", r, e, reflect.DeepEqual(tv, r))

	vm := map[string]interface{}{"name": "liuyi's father"}
	ss.UpdateDoc(ti, tt, tm, vm)
	tv["name"] = vm["name"]
	e = ss.GetDoc(ti, tt, tm, &r)
	t.Log("get", r, e, reflect.DeepEqual(tv, r))

	ss.DeleteDoc(ti, tt, tm)
	for k, _ := range r {
		delete(r, k)
	}
	e = ss.GetDoc(ti, tt, tm, &r)
	t.Log("get", r, e != nil)
}
