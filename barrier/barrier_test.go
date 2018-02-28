package barrier

import (
    "testing"
    "sync"
)


type counter struct {
    m sync.Mutex
    c int
}

func NewCounter() *counter {
    return &counter{}
}

func (c *counter) Inc() int {
    c.m.Lock();
    defer c.m.Unlock();
    c.c++;
    return c.c;
}

func (c *counter) Get() int {
    c.m.Lock();
    defer c.m.Unlock();
    return c.c;
}


func TestBarrier(t *testing.T) {  
    c := NewCounter();
    b := New(3);
    var wg sync.WaitGroup;
    
    for i := 0; i < 3; i++ {
        wg.Add(1)
        index := i;
        go worker(index, t, c, b, &wg);
    }
        
    t.Logf("Main: waiting all\n");
    wg.Wait()    
}


func worker(index int, t *testing.T, c *counter, b *Barrier, wg *sync.WaitGroup) {
    for i := 0; i < 3; i++ {
        b.Wait();
        c.Inc();
        b.Wait();
        t.Logf("Worker #%d: %d\n", index, c.Get());
    }
    
    wg.Done();
}
 