# Theory of Operation - Parallelization

This is my attempt to make this task more performant by parallelizing the heavy lift (IsPrime).

For this particular challenge, the order of primes matters, and the Nth prime can't be known from the beginning. If one simply parallelized the computation without synchronizing the insertion order, you introduce a race condition where the Nth + 1 prime can get recorded before the Nth prime. In that case, the incorrect Nth prime would be returned.

To account for this, I am using channels to synchronize the order in which the go routines record primes as they are found. Code flow outline:


- for each successive natural number create a new go routine, passing in a read channel and a write channel
  - the read channel is shared between the current go routine to the previous go routine
  - the write channel is shared between the current go routine to the next go routine
  - i.e. if the current go routine is evaluating 1234567, the read channel is shared between it and the go routine evaluating 1234566, and the write channel is shared between it and the go routine evaluating 1234568
- go routine immediately begins evaluating primeness of target number
- after primeness is established, go routine attempts to read from read channel, effectively forcing the routine to wait until the previous routine has finished
- the go routine only becomes unblocked when the previous go routine writes to it's read channel
- the now unblocked go routine can record it's prime value (assuming it has one), before writing to the write channel, thereby freeing up the next go routine to record it's value, and so on
