package main

import "fmt"

// Created a structure by the name of RingBuffer. Observe array and map both used. Arrays are slow to find a key if stored instead them but are required as we have to folllow FIFO but at the same time use map to find key from the stack of keys are map is fast in finding
type RingBuffer struct {

	//this is the buffer variable of type interface so that any variable type can be used and not limited to string or int or just one type. Power of interface
	buffer []interface{}

	//using map to find uuid key instead of looping through array as that would be very slow for a large stack. Again here also using interface instead of a variable
	keys map[interface{}]bool

	//this is size of stack, more elements added would push the earlier ones behind, draining out olds ones out of stack to maintain stack size "length"
	length int
}

func (this *RingBuffer) Init(length int) {

	//initialising stack, observe no type used. only interface type used

	this.buffer = make([]interface{}, 0)

	//intiailising map stack for stocking uuid
	this.keys = make(map[interface{}]bool)

	//length defined for the length of the stack
	this.length = length

}

func (this *RingBuffer) Add(uuid interface{}) bool {

	//Receive tuples, ignoring key value in case present as I already have uuid there is why _
	//if uuid present then key_present boolean value turns true
	_, key_present := this.keys[uuid]

	//exit without adding uuid to ring buffer as key already present. Return false
	if key_present {

		return false
	}

	// UUID not found in stack so add it to the array stack
	this.buffer = append(this.buffer, uuid)

	//like wise add it in the map
	this.keys[uuid] = true

	//check if the stack after append has increased beyond the length defined, if so then trim it by pushing out of stack the head element
	if len(this.buffer) > this.length {

		//before trimming the array, delete that key from map so that map also has only those keys that are in array
		delete(this.keys, this.buffer[0])

		//trimming the array by removing head element
		this.buffer = this.buffer[1 : this.length+1]

	}

	return false

}

// print the stack
func (this *RingBuffer) Print() {

	fmt.Println(this.buffer, this.keys)

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
