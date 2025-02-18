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
