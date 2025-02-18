package bloomfilter

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type bloomFilter struct {
	bitArray []uint64
	m        int
	k        int
}

func New(n int) *bloomFilter {
	p := 0.01 // false positive rate of 1%, suitable for most general purpose case

	// formula to calculate m, k comes from chatGPT, which based it on
	// https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
	m := math.Ceil(-float64(n) * math.Log(p) / (math.Ln2 * math.Ln2))
	k := int(math.Round(m * math.Ln2 / float64(n)))

	m_int := int(m)
	return &bloomFilter{
		bitArray: make([]uint64, (m_int+63)/64),
		m:        m_int,
		k:        k,
	}
}

func (b *bloomFilter) Add(s string) {
	hashes := b.generateHashes([]byte(s)) // including % m for each hash
	for _, hash := range hashes {
		idx := hash >> 6
		pos := hash & 63
		b.bitArray[idx] = (b.bitArray[idx] | (uint64(1) << pos))
	}
}

func (b *bloomFilter) Has(s string) bool {
	hashes := b.generateHashes([]byte(s))
	for _, hash := range hashes {
		idx := hash >> 6
		pos := hash & 63
		if b.bitArray[idx]&(uint64(1)<<pos) == 0 {
			return false
		}
	}
	return true
}

func (b *bloomFilter) generateHashes(data []byte) []uint64 {
	h1, h2 := murmur3.Sum128(data) // Get two independent hashes
	hashes := make([]uint64, b.k)

	for i := 0; i < b.k; i++ {
		hashes[i] = (h1 + uint64(i+1)*h2) % uint64(b.m)
	}

	return hashes
}
