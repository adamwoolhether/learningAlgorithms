## Key-Value Association

Collections of key-value pairs are a _symbolic table_ data type. Keys are not stored in a specific order, but it brings optimal performance.

Finding a value for a given key in a array of N (k, v) pairs would give a performance of O(N).
A hash will give O(1), independent of the number of k,v pairs.

## Hash Function and Hash Codes
Our hash function `base26` maps keys to a fixed-size _hash code_ value. 

Hash funcitons have one necessary property: two objects that are equal to each other _must compute to the same hash code_.
The computed hash code is usually stored to reduce the overall computation.

Hash functions *don't* need a unique hash for each key. `hash(key) % M` uses the modulo operation to compute a value that
returns an integer from 0 to M -1.

### Detecting and Resolving Collisions with Linear Probing
_Open addressing_ resolves hash collisions by _probing_ through other locations in the table when put encounters a
collision. There are a few approaches. _Linear probing_ incrementally checks higher index positions for the next available empty bucket, and starts at
position 0 if the end is reached.

This only succeeds when we ensure that there is always one empty bucket.

