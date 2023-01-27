import "fmt"

func countPrimes(limit int) int {
    if limit <= 2 {
        return 0
    }

    primes := make([]bool, limit+1)
    for i := 2; i < limit; i++ {
        primes[i] = true
    }

    for i := 2; i*i < limit; i++ {
        if primes[i] {
            for j := i * i; j <= limit; j += i {
                primes[j] = false
            }
        }
    }

    count := 0
    for i := 2; i < limit; i++ {
        if primes[i] {
            count++
        }
    }

    return count
}
