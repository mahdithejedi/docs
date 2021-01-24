## What is lru_cache?
It is just a cache clean-up strategy.

A computer has limited memory cache. If the cache is full, some contents need to be removed from cache to provide space for new content. However, which part of the cache should be removed? We hope to remove not so useful contents, while leaving useful contents untouched for future usage. So the question is, what are the criteria to determine if the data is useful or not?

LRU (Least Recently Used) cache clean-up algorithm is a common strategy. According to the name, the latest used data should be useful. Hence, when the memory cache is full, we should prioritize removing those data that haven't been used for long are not useful.

## Resources
[How to Implement LRU Cache (labuladong gitbook)](https://labuladong.gitbook.io/algo-en/iv.-high-frequency-interview-problem/lru_algorithm)
<br />
[Caching in Python Using the LRU Cache Strategy (realpython)](https://realpython.com/lru-cache-python/)
<br />
[C implementation (geekforgeeks)](https://www.geeksforgeeks.org/lru-cache-implementation/)
<br />
[Python implementation (geekforgeeks)](https://www.geeksforgeeks.org/lru-cache-in-python-using-ordereddict/)
<br />

