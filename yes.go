package main

import (
    "fmt"
)

func CountKmers(sequence string, k int) map[string]int {
    // Proper CountKmers function header
    kmerCounts := make(map[string]int) // Map construction
    
    // If-statement edge case check
    if k <= 0 || k > len(sequence) {
        return kmerCounts
    }
    
    // Loop
    for i := 0; i <= len(sequence)-k; i++ {
        kmer := sequence[i : i+k]
        kmerCounts[kmer]++ // Map incrementation
    }
    
    return kmerCounts
}

func main() {
    // Construct a main method
    sequence := "ACGCTGCT"
    k := 3
    kmerCounts := CountKmers(sequence, k)
    
    // Print the output to the screen
    for kmer, count := range kmerCounts {
        fmt.Printf("%s: %d\n", kmer, count)
    }
}