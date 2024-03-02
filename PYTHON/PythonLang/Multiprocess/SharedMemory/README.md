# python SharedMemory

## Cpython SharedMemory In Unix
Python behind the scenes when you
call [_SharedMemoery_](https://docs.python.org/3.11/library/multiprocessing.shared_memory.html#multiprocessing.shared_memory.SharedMemory) ,
[cpython first will call](https://github.com/python/cpython/blob/7e30821b17b56bb5ed9799f62eb45e448cb52c8e/Lib/multiprocessing/shared_memory.py#L104)
[`shm_open`](https://man7.org/linux/man-pages/man3/shm_open.3.html) to open a shared memory which Linux will return a fd. Then if you set a size then cpython will call
[`ftruncate`](https://man7.org/linux/man-pages/man2/ftruncate.2.html) to set the size of shared memory then cpython will call [`mmap`](https://man7.org/linux/man-pages/man2/mmap.2.html) to
attach that process to given process


### Sources
[python SharedMemory explanation](https://python.plainenglish.io/a-simple-guide-to-shared-memory-in-python-3c2e946ece0)
<br />
[python shared memory documentation](https://docs.python.org/3.11/library/multiprocessing.shared_memory.html)