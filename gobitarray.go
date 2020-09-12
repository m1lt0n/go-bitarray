package gobitarray

import (
	"errors"
	"fmt"
	"sync"
)

// BitArray : Holds the bit array data and metadata
type BitArray struct {
	size int
	data []byte
	mux  sync.Mutex
}

// New : Create a new bit array of a specified size
func New(size int) BitArray {
	return BitArray{size: size, data: make([]byte, size/8+1)}
}

// Set : set a specific position's bit. Returns the position of the bit set and an error
// if the position provided is invalid
func (arr *BitArray) Set(position int) (int, error) {
	arr.mux.Lock()
	defer arr.mux.Unlock()

	err := ensurePosition(position, arr.size)

	if err != nil {
		return -1, err
	}

	arr.data[position/8] = arr.data[position/8] | (1 << (position % 8))

	return position, nil
}

// Unset : unset a specific position's bit. Returns the position of the bit unset and an error
// if the position provided is invalid
func (arr *BitArray) Unset(position int) (int, error) {
	arr.mux.Lock()
	defer arr.mux.Unlock()

	err := ensurePosition(position, arr.size)

	if err != nil {
		return -1, err
	}

	arr.data[position/8] = arr.data[position/8] & (255 ^ (1 << (position % 8)))

	return position, nil
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
		return errors.New(fmt.Sprintf("Provided position %d exceeds it array's max index %d", position, size-1))
	} else {
		return nil
	}
}
