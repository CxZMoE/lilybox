# LilyBox
Lilybox is a storage microservice easy to use. Powered by rpc.  
Currently lilybox supports there kinds of data type: `string`, `int` and `bytes`.  
But everything is stored in RAM, next thing to do is adding adapter for FS/DBS and adding new query method such as http.
# How to use
### Starting a Server
``` shell
./server.exe -port 12000 # run the server on port 12000
```
### Do Client RPC Call
``` golang
type Storage struct {
	pool map[string]interface{}
}
// Storage of type string
func (ac *Storage) AddString(args [2]string) {
	ac.pool[args[0]] = args[1]
}
func (ac *Storage) GetString(key string) string {
	return ac.pool[key].(string)
}
// Storage of type int
func (ac *Storage) AddInt(args [2]string) {
	ac.pool[args[0]], _ = strconv.Atoi(args[1])
}
func (ac *Storage) GetInt(key string) int {
	return ac.pool[key].(int)
}
// Storage of type bytes
func (ac *Storage) AddBytes(args [2][]byte) {
	ac.pool[string(args[0])] = args[1]
}
func (ac *Storage) GetBytes(key string) []byte {
	return ac.pool[key].([]byte)
}

func main() {
    client := gorpc.NewTCPClient("127.0.0.1:6666")
    dispatcher := gorpc.NewDispatcher()

    storage := new(Storage)
    storage.pool = make(map[string]interface{})
    dispatcher.AddService("Storage", storage)

    serviceClient := d.NewServiceClient("Storage", c)
    client.Start()

    // Set string value
    serviceClient.Call("AddString", [2]string{"string_text", func() string {
        var str string
        for i := 0; i < 128; i++ {
            str += "haha"
        }
        return str
    }()})
    res, err := sc.Call("GetString", "string_text")
    if err != nil {
        panic(err)
    }
    log.Println("GetString:", res)

    // Set int value
    sc.Call("AddInt", [2]string{"count", "1"})
    res, err = sc.Call("GetInt", "count")
    if err != nil {
        panic(err)
    }
    log.Println("GetInt:", res)

    // set value of []byte
    sc.Call("AddBytes", [2][]byte{[]byte("bytes_text"), []byte("sadasdasd")})
    res, err = sc.Call("GetBytes", "bytes_text")
    if err != nil {
        panic(err)
    }
    log.Println("GetBytes:", string(res.([]byte)))
    defer c.Stop()
}
```

