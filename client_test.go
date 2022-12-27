package main

import (
	"testing"

	"github.com/valyala/gorpc"
)

func TestOne(t *testing.T) {
	{
		c := gorpc.NewTCPClient("127.0.0.1:6666")
		d := gorpc.NewDispatcher()
		st := new(Storage)
		st.pool = make(map[string]interface{})
		d.AddService("Storage", st)
		sc := d.NewServiceClient("Storage", c)
		c.Start()

		sc.Call("AddString", [2]string{"haha", func() string {
			var str string
			for i := 0; i < 128; i++ {
				str += "haha"
			}
			return str
		}()})
		res, err := sc.Call("GetString", "haha")
		if err != nil {
			panic(err)
		}
		t.Log("GetString:", res)

		sc.Call("AddInt", [2]string{"count", "1"})
		res, err = sc.Call("GetInt", "count")
		if err != nil {
			panic(err)
		}
		t.Log("GetInt:", res)

		sc.Call("AddBytes", [2][]byte{[]byte("ack"), []byte("sadasdasd")})
		res, err = sc.Call("GetBytes", "ack")
		if err != nil {
			panic(err)
		}
		t.Log("GetBytes:", string(res.([]byte)))
		defer c.Stop()
	}
}
