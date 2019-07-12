// implements Knuth or Fisher-Yates shuffle
package knuth

import (
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func Shuffle(slc []interface{}) {
    N := len(slc)
    for i := 0; i < N; i++ {
        // choose index uniformly in [i, N-1]
        r := i + rand.Intn(N-i)
        slc[r], slc[i] = slc[i], slc[r]
    }
}

func ShuffleInts(slc []int) {
    N := len(slc)
    for i := 0; i < N; i++ {
        // choose index uniformly in [i, N-1]
        r := i + rand.Intn(N-i)
        slc[r], slc[i] = slc[i], slc[r]
    }
}

func ShuffleStrings(slc []string) {
    N := len(slc)
    for i := 0; i < N; i++ {
        // choose index uniformly in [i, N-1]
        r := i + rand.Intn(N-i)
        slc[r], slc[i] = slc[i], slc[r]
    }
}
