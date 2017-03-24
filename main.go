package main

import "fmt"

// Created a structure by the name of RingBuffer
type RingBuffer struct {

	//this is the buffer variable of type interface so that any variable type can be used and not limited to string or int or just one type. Power of interface
	buffer []interface{}

	//this is size of stack, more elements added would push the earlier ones behind, draining out olds ones out of stack to maintain stack size "length"
	length int
}

func (this *RingBuffer) Init(length int) {

	//initialising stack, observe no type used. only interface type used

	this.buffer = make([]interface{}, 0)

	//length defined for the length of the stack
	this.length = length

}

func (this *RingBuffer) Add(uuid interface{}) bool {

	//Looping to find if the uuid exists as only unique UUID is added
	for loop := 0; loop < len(this.buffer); loop++ {

		//if UUID is found then return false without even adding
		if this.buffer[loop] == uuid {
			return false

		}

	}

	//if UUID not found then you can add it to the stack
	this.buffer = append(this.buffer, uuid)

	//check if the stack after append has increased beyond the length defined, if so then trim it by pushing out of stack the head element
	if len(this.buffer) > this.length {

		this.buffer = this.buffer[1 : this.length+1]

	}

	return false

}

// print the stack
func (this *RingBuffer) Print() {

	fmt.Println(this.buffer)

}

func main() {

	//initialise a RingBuffer, observe nothing defined as variable type
	var intBuffer RingBuffer

	//defined the stack size as 3. Means no more than 3 unique elements. If additional pushed, the 0th element would be pushed out
	intBuffer.Init(3)

	//first element tells the struct that the elements are going to be integer
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

	//Another RingBuffer defined but nothing defined as variable
	var stringBuffer RingBuffer

	//defined the stack size as 3. Means no more than 3 unique elements. If additional pushed, the 0th element would be pushed out
	stringBuffer.Init(3)

	//first element tells the struct that the elements are going to be string
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
