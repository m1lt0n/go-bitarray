package gobitarray

import (
	"fmt"
	"sync"
)

// BitArray : Holds the bit array data and metadata
type BitArray struct {
	size int
	data []byte
	mux  sync.Mutex
}

// IndexError : Represents an index out of range error
type IndexError struct {
  position int
  arraySize int
}

func (err *IndexError) Error() string {
  return fmt.Sprintf("Provided position %d is out of range. Maximum allowed index is %d.", err.position, err.arraySize - 1)
}

// New : Create a new bit array of a specified size
func New(size int) BitArray {
	return BitArray{size: size, data: make([]byte, size/8+1)}
}

// Set : set a specific position's bit. Returns an error if unsuccessful
// if the position provided is invalid
func (arr *BitArray) Set(position int) error {
	arr.mux.Lock()
	defer arr.mux.Unlock()

	err := ensurePosition(position, arr.size)

	if err != nil {
		return err
	}

	arr.data[position/8] = arr.data[position/8] | (1 << (position % 8))

  return nil
}

// Unset : unset a specific position's bit. Returns an error if unsuccessful
// if the position provided is invalid
func (arr *BitArray) Unset(position int) error {
	arr.mux.Lock()
	defer arr.mux.Unlock()

	err := ensurePosition(position, arr.size)

	if err != nil {
		return err
	}

	arr.data[position/8] = arr.data[position/8] & (255 ^ (1 << (position % 8)))

	return nil
}

// Get : Get the value of a bit. Returns the value of the bit and an error
// if the position provided is invalid
func (arr *BitArray) Get(position int) (int, error) {
	arr.mux.Lock()
	defer arr.mux.Unlock()

	err := ensurePosition(position, arr.size)

	if err != nil {
		return -1, err
	}

	return arr.getBit(position), nil
}

// Toggle : Toggles the value of a bit. Returns the new value of the bit and an error
// if the position provided is invalid
func (arr *BitArray) Toggle(position int) (int, error) {
	arr.mux.Lock()
	defer arr.mux.Unlock()

	err := ensurePosition(position, arr.size)

	if err != nil {
		return -1, err
	}

	arr.data[position/8] = arr.data[position/8] ^ (1 << (position % 8))

	return arr.getBit(position), nil
}

// Reset : Unsets all the bits of the array
func (arr *BitArray) Reset() {
	arr.mux.Lock()
	defer arr.mux.Unlock()

	arr.data = make([]byte, arr.size/8+1)
}

func (arr *BitArray) getBit(position int) int {
	if arr.data[position/8]&(1<<(position%8)) > 0 {
		return 1
	} else {
		return 0
	}
}

func ensurePosition(position int, size int) error {
	if position >= size {
		return &IndexError{position, size}
	} else {
		return nil
	}
}
