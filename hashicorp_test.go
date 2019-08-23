package cachebench

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	lru "github.com/hashicorp/golang-lru"
)

func initHashicorpGet(size int) (cache *lru.Cache, keys []string) {
	cache, _ = lru.New(size)
	keys = make([]string, size)
	for i := 0; i < size; i++ {
		id := uuid.New().String()
		keys[i] = id
		m := &SomeProtoStruct{
			I:          rand.Int63(),
			D:          rand.Int63(),
			B:          rand.Int63n(2),
			ID:         rand.Int63(),
			Time:       time.Now().Unix(),
			T1:         rand.Int63(),
			T2:         rand.Int63(),
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

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			k := rand.Intn(1023)
			_, found := cache.Get(keys[k])
			if !found {
				b.Errorf("key %s not found", keys[k])
			}
		}
	})
}

func BenchmarkHashicorpSet(b *testing.B) {
	cache, err := lru.New(10000)
	if err != nil {
		b.Error(err)
		b.FailNow()
	}

	m := &SomeProtoStruct{
		I:          rand.Int63(),
		D:          rand.Int63(),
		B:          rand.Int63n(2),
		ID:         rand.Int63(),
		Time:       time.Now().Unix(),
		T1:         rand.Int63(),
		T2:         rand.Int63(),
		Name1:      "123",
		Name2:      "321",
		StringTime: "123123",
		Type:       "123123",
		Status:     "123123",
		S:          "123123",
	}

	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func() {
				id := uuid.New().String()

				evicted := cache.Add(id, m)
				if evicted {
					//b.Logf("key %s was evicted", id)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
