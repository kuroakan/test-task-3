package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type FirstLevel struct {
	dataType byte
	data     interface{}
}

type dataWithIntegers struct {
	intData []int32
}

type dataWithFloats struct {
	floatData []float32
}

func main() {
	rand.Seed(time.Now().UnixNano())

	x := make([]byte, 129)

	x[0] = byte(rand.Intn(2))

	for i := 1; i < len(x); i++ {
		x[i] = byte(rand.Intn(255))
	}

	fl, err := FillUpLevel(x)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fl)
}

func FillUpLevel(data []byte) (result FirstLevel, err error) {
	if len(data[1:])%4 != 0 {
		return FirstLevel{}, errors.New("incorrect data length")
	}

	buf := bytes.NewReader(data[1:])

	switch data[0] {
	case 0:
		result.dataType = data[0]
		d := dataWithIntegers{make([]int32, len(data[1:])/4)}

		err := binary.Read(buf, binary.BigEndian, &d.intData)
		if err != nil {
			return FirstLevel{}, err
		}

		result.data = d
	case 1:
		result.dataType = data[0]
		d := dataWithFloats{make([]float32, len(data[1:])/4)}

		err := binary.Read(buf, binary.BigEndian, &d.floatData)
		if err != nil {
			return FirstLevel{}, err
		}

		result.data = d
	default:
		return FirstLevel{}, errors.New("unknown dataType")
	}

	return result, nil
}
