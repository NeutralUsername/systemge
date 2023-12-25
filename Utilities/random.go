package Utilities

import (
	"math/rand"
	"sync"
	"time"
)

type Random struct {
	RandSeed      rand.Source
	RandSeedMutex *sync.Mutex
}

func (rand *Random) GenerateRandomString(length int, validChars string) string {
	str := ""
	for i := 0; i < length; i++ {
		rand.RandSeedMutex.Lock()
		str += string(validChars[rand.RandSeed.Int63()%int64(len(validChars))])
		rand.RandSeedMutex.Unlock()
	}
	return str
}

func CreateRandom() *Random {
	return &Random{
		RandSeed:      rand.NewSource(time.Now().UnixNano()),
		RandSeedMutex: &sync.Mutex{},
	}
}

func (rand *Random) GenerateRandomNumber(min, max int) int {
	rand.RandSeedMutex.Lock()
	number := int(rand.RandSeed.Int63()%int64(max-min+1)) + min
	rand.RandSeedMutex.Unlock()
	return number
}
