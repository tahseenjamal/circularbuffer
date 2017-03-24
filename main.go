package main

import "fmt"

type RingBuffer struct {
	buffer []interface{}

	length int
}

func (this *RingBuffer) Init(length int) {

	this.buffer = make([]interface{}, 0)
	this.length = length

}

func (this *RingBuffer) Add(uuid interface{}) bool {

	for loop := 0; loop < len(this.buffer); loop++ {

		if this.buffer[loop] == uuid {
			return false

		}

	}

	this.buffer = append(this.buffer, uuid)

	if len(this.buffer) > this.length {

		this.buffer = this.buffer[1 : this.length+1]

	}

	return false

}

func (this *RingBuffer) Print() {

	fmt.Println(this.buffer)

}

func main() {

	var intBuffer RingBuffer

	intBuffer.Init(3)

	intBuffer.Add(1)
	intBuffer.Print()
	intBuffer.Add(2)
	intBuffer.Print()
	intBuffer.Add(3)
	intBuffer.Print()
	intBuffer.Add(4)
	intBuffer.Print()
	intBuffer.Add(4)
	intBuffer.Print()
	intBuffer.Add(5)
	intBuffer.Print()
	intBuffer.Add(6)
	intBuffer.Print()

	var stringBuffer RingBuffer

	stringBuffer.Init(3)

	stringBuffer.Add("a")
	stringBuffer.Print()
	stringBuffer.Add("b")
	stringBuffer.Print()
	stringBuffer.Add("a")
	stringBuffer.Print()
	stringBuffer.Add("c")
	stringBuffer.Print()
	stringBuffer.Add("d")
	stringBuffer.Print()
	stringBuffer.Add("e")
	stringBuffer.Print()
	stringBuffer.Add("f")
	stringBuffer.Print()

}
