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

This only succeeds when we ensure that there is always one empty bucket, and can be inefficient if the incorrect amount is used.

#### Linked Lists
Linked lists store data in memory fragments called _nodes_ that are linked together. Allowing us to follow _links_, or 
_references_ *from the first node to the next node. This allows is to conduct the following operations:
* Prepending a value
* Appending a value
* Inserting a value
* Deleting a value

_Separate chaining_ uses linked lists.

### Evaluation
The linked list implementation can grow as large as needed, limited only by memory. For open addressing, N must be smaller
than the given max.

For storage, open addressing requires storage proportional to M and N. It is `O(M)` because N < M.
Storage for separate chaining is O(M+N).

The key metric for runtime performance is the number of times an entry is inspected. The worse case of `get()` is `O(N)`.

In hashtables, the **load factor**, `N/M` determines its performance. Load factor is also the _alpha_.
Hashtables become increasingly inefficient when the load factor is higher than 0.75

### Dynamic Hashtables
_Geometric resizing_ is when we simply double the size of the storage array. Doing so properly, to avoid losing keys, 
requires first creating a temporary hashtable with twice the original storage and rehashing all the entries into the new
table.

Geometric resizing ensures that resizing occurs much less frequently as the table grows in size.

#### Performance of Dynamic Hashtables
If the worst case, `put()` and `get()` are `O(N)`. Because N < M, N/M is a constant, `O(1)`, and is independent of N.

## Perfect Hashing
A perfect hash is an optimal hash table where each hash code is a unique index location.