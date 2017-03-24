package main

import "fmt"

type RingBuffer struct {
	buffer []string

	length int
}

func (this *RingBuffer) Add(uuid string) bool {

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

	newbuf := RingBuffer{make([]string, 0), 3}

	newbuf.Add("a")
	newbuf.Print()
	newbuf.Add("b")
	newbuf.Print()
	newbuf.Add("a")
	newbuf.Print()
	newbuf.Add("c")
	newbuf.Print()
	newbuf.Add("d")
	newbuf.Print()
	newbuf.Add("e")
	newbuf.Print()
	newbuf.Add("f")
	newbuf.Print()

}
