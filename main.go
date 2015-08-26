package main
import (
	"github.com/aerospike/aerospike-client-go"
	"encoding/gob"
	"bytes"
	"log"
)

type TestObj struct {
	Hello string
	World string
	pvtProperty string
}

func main()  {
	client, err := aerospike.NewClient("127.0.0.1", 3000)
	panicOnError(err)

	testObj := &TestObj{
		Hello: "Arnold",
		World: "Facepalmer",
		pvtProperty: "Tralalala",
	}

	var encoded bytes.Buffer
	encoder := gob.NewEncoder(&encoded)
	encoder.Encode(testObj)

	bins := aerospike.BinMap{
		"test1": encoded.Bytes(),
		"test2": 42,
	}

	key, err := aerospike.NewKey("test", "aerospike", "key")
	panicOnError(err)

	err = client.Put(nil, key, bins)
	panicOnError(err)

	// read it back!
	rec, err := client.Get(nil, key)
	panicOnError(err)

	test1 := rec.Bins["test1"].([]byte)

	buf := bytes.NewBuffer(test1)

	var decodedObj TestObj
	decoder := gob.NewDecoder(buf)
	err = decoder.Decode(&decodedObj)

	log.Println(decodedObj)
	log.Println(decodedObj.pvtProperty)
}

func panicOnError(err error)  {
	if err != nil {
		panic(err)
	}
}
