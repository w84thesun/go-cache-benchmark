package benchCaches

import (
	"fmt"
	"github.com/google/uuid"
	lru "github.com/hashicorp/golang-lru"
	"math/rand"
	"testing"
	"time"
)

var mm *SomeStruct
var ok bool

func initHashicorpGet(size int) (cache *lru.Cache, keys []string) {
	cache, _ = lru.New(size)
	keys = make([]string, size)
	for i := 0; i < size; i++ {
		id := uuid.New().String()
		keys[i] = id
		m := &SomeStruct{
			I:          rand.Int(),
			D:          rand.Int(),
			B:          rand.Intn(2),
			ID:         rand.Int(),
			Time:       time.Now().Unix(),
			T1:         rand.Int(),
			T2:         rand.Int(),
			Name1:      "123",
			Name2:      "321",
			StringTime: "123123",
			Type:       "123123",
			Status:     "123123",
			S:          "123123",
		}
		evicted := cache.Add(id, m)
		if evicted {
			fmt.Printf("[init data] key %s was evicted", id)
		}
	}

	return cache, keys
}

func BenchmarkHashicorpGet(b *testing.B) {
	cache, keys := initHashicorpGet(1024)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		k := rand.Intn(1023)
		res, found := cache.Get(keys[k])
		if !found {
			b.Errorf("key %s not found", keys[k])
		}

		mm, ok = res.(*SomeStruct)

		if !ok {
			b.Errorf("obj type is not equal SomeStruct")
		}
	}
}

func BenchmarkHashicorpSet(b *testing.B) {
	cache, err := lru.New(b.N)
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	for i := 0; i < b.N; i++ {
		id := uuid.New().String()

		m := &SomeStruct{
			I:          rand.Int(),
			D:          rand.Int(),
			B:          rand.Intn(2),
			ID:         rand.Int(),
			Time:       time.Now().Unix(),
			T1:         rand.Int(),
			T2:         rand.Int(),
			Name1:      "123",
			Name2:      "321",
			StringTime: "123123",
			Type:       "123123",
			Status:     "123123",
			S:          "123123",
		}
		evicted := cache.Add(id, m)
		if evicted {
			b.Logf("key %s was evicted", id)
		}
	}
}
