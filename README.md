# Overview
a Bloom Filter

# Properties
- false positive rate is set to 1%
- the number of bits used for the array is `m = − n*ln(p)/(ln2)^2`
- the number of hash functions is `k= m*ln2/n`
- we use murmurHash3 & double hashing to simulate k hash functions

_Note_: formulas to calculate m & k are based on
https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions

# Example
```go
b := bloomfilter.New(n) // n is number of items in the set
b.Add('something')
b.Has('something') // always true
b.Has('something else') // false, but maybe true for other that is not in the set
```

# Implementation details
## Data structure
- a struct that contains the array of bits, m & k

## Algorithm
- There is not much to the algorithm, it's just a matter of setting the bits to
  true when adding an item and checking if all the bits are true when checking
