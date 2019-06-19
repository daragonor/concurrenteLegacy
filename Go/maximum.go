package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

const N = 1000000
const TH = 2

func max(s []float64) {
    n := len(s)
    b := int(math.Ceil(float64(n)) / float64(TH))
    m := make([]float64, TH)
    for i := 0; i < TH; i++ {
        if s[i] > maxi{
            maxi = s[i[]]
    
    }
    return maxi
}

func main() {
    rand.Seed(1981)
    s :=make([]float64, N)
    for i := range s{
        s[i] = float64(rand.Intn(1))
    }
    fmt.Println([s:10])
    
}