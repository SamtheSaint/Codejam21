package moonsandumbrellas

import (
	"log"
)

// MinimumCost returns the minimum cost for the mural in s
// x is the cost of CJ
// y is the cost of JC
func MinimumCost(x, y int, s string) (cost int) {
	for i, v := range s {
		if v == '?' {
			s1 := s[:i] + "C" + s[i+1:]
			s2 := s[:i] + "J" + s[i+1:]
			cost = minInt(MinimumCost(x, y, s1), MinimumCost(x, y, s2))
			return
		}
	}
	cost = evaluateCost(x, y, s)
	return
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var cache map[string]int
var cacheSize int = 10

func ClearCache() {
	cache = map[string]int{}
}

func AddToCache(x, y int, s string) (cost int) {
	if len(s) > cacheSize {
		log.Fatalln("Incorrect entry to cache size")
	}
	if val, ok := cache[s]; ok {
		cost = val
		return
	}
	for i := 0; i < len(s) - 1; i++ {
		if s[i:i+2] == "CJ" {
			cost += x
		} else if s[i:i+2] == "JC" {
			cost += y
		}
	}
	cache[s] = cost
	return
}

func evaluateCost(x, y int, s string) (cost int) {
	if len(s) <= cacheSize {
		return AddToCache(x, y, s)
	}

	for i:= 0; i < len(s); i += cacheSize {
		if i + cacheSize < len(s) {
			cost += AddToCache(x, y, s[i:i+cacheSize])
			cost += AddToCache(x, y, s[i+cacheSize-1:i+cacheSize+1])
		} else {
			cost += AddToCache(x, y, s[i:])
		}
	}
	return cost
}