package barrier

import (
    "sync"
)


type Barrier struct {
    c int
    n int
    m sync.Mutex
    ch chan int
}


func New(n int) *Barrier {
    return &Barrier{
        n: n,
        ch: make(chan int, n),
    };
}


func (b *Barrier) Wait() {
    b.m.Lock()
        b.c++;
        if (b.c == b.n) {
            for i:= 0; i < b.n; i++ {
                b.ch <- 1
            }
            b.c = 0;
        }    
    b.m.Unlock()
    <-b.ch
}