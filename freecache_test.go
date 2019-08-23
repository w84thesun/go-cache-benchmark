package cachebench

import (
	"fmt"
	"math/rand"
	"sync"
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
		encoded, _ := m.Marshal()
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
		wg := sync.WaitGroup{}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func() {

				k := rand.Intn(1023)
				res, err := cache.Get([]byte(keys[k]))
				if err != nil {
					b.Error(err)
				}

				m := new(SomeProtoStruct)
				err = m.Unmarshal(res)
				if err != nil {
					b.Error(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkFreeCacheSet(b *testing.B) {
	cache := freecache.NewCache(10 * 1024 * 1024)

	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func() {
				id := uuid.New().String()
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
				encoded, _ := m.Marshal()

				err := cache.Set([]byte(id), encoded, 0)
				if err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
