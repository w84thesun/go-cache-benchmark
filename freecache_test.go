package benchCaches

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/coocood/freecache"
	"github.com/google/uuid"
)

func initFreeCacheGet(size int) (cache *freecache.Cache, keys []string) {
	cache = freecache.NewCache(size * 1024 * 1024)
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
		encoded, _ := m.Encode()
		err := cache.Set([]byte(id), encoded, 0)
		if err != nil {
			fmt.Println(err)
		}
	}

	return cache, keys
}

func BenchmarkFreeCacheGet(b *testing.B) {
	cache, keys := initFreeCacheGet(1024)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		k := rand.Intn(1023)
		res, err := cache.Get([]byte(keys[k]))
		if err != nil {
			b.Error(err)
		}

		m := new(SomeStruct)
		err = m.Decode(res)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFreeCacheSet(b *testing.B) {
	cache := freecache.NewCache(10 * 1024 * 1024)

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
		encoded, _ := m.Encode()
		err := cache.Set([]byte(id), encoded, 0)
		if err != nil {
			fmt.Println(err)
		}
	}
}
