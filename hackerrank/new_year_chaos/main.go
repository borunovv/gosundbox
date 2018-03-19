// https://www.hackerrank.com/challenges/new-year-chaos/problem

package main

import (
    "fmt"
)

func main() {
    testCount := readLineInt()
    for t := 0; t < testCount; t++ {
        count := readLineInt()
        ints := readLineInts(count)
        result := process(ints)
        if (result >= 0) {
            fmt.Printf("%d\n", result)
        } else {
            fmt.Println("Too chaotic")
        }
    }
}

func readLineInt() int {
    return readLineInts(1)[0] 
}

func readLineInts(count int) []int {
    line := make([]int, count)
    for i, _ := range line {
        fmt.Scanf("%d", &line[i])
    }
    fmt.Scanf("\n")
    return line
}

func process(list []int) int {   
    var usage []int = make([]int, len(list))
    var positions []int = make([]int, len(list))
    
    for i, val := range list {
        list[i]--
        positions[val - 1] = i
    }
        
    m := 0
    for i, val := range list {
        if (i == val) {continue}
        to := positions[i]        
        for j := to; j > i; j-- {
            list[j] = list[j - 1]
            usage[list[j]]++
            if (usage[list[j]] > 2) {return -1}
            positions[list[j]]++
        }
        list[i] = i
        m += to - i        
    }
        
    return m
}

