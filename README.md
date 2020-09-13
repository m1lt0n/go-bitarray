# go-bitarray

Welcome to go-bitarray! go-bitarray is a go implementation of bit array that allows creating and querying a bit array.

Bit arrays are a compact, space efficient data structure that represent a group of boolean values. A very common use is in probabilistic data structures (like Bloom filter).

## Installation

```bash
go get github.com/m1lt0n/go-bitarray
```

or add github.com/m1lt0n/go-bitarray to your module's dependencies.

## Usage

In order to create a bit array, only its size is required:

```golang
# Create a bit array that can store 100 bits
arr := gobitarray.New(100)
```

Some common operations on individual bits that go-bitarray supports:

```golang
# set the 5th bit
err := arr.Set(5)

# get the 5th bit
bitValue, err := arr.Get(5)

# unset the 5th bit
err := arr.Unset(5)

# toggle the value of the 10th bit
newValue, err := arr.Toggle(10)
```

The set and unset operations are idempotent, as the final state of the bit is the same. This is not the case with toggle, as it changes the bit value from 1 to 0 or 0 to 1. The Set, Unset, Toggle and Get operations can be unsuccessful if the provided index is out of the range of the bit array. In this case, an IndexError is returned.

Some operations that apply on the bit array itself:

```golang
# resets the values of all bits to zero
arr.Reset()
```

**Note**: The bit array implementation is thread-safe.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/[USERNAME]/go-bitarray. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The module is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the go-bitarray project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/[USERNAME]/go-bitarray/blob/master/CODE_OF_CONDUCT.md).

