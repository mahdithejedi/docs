# What is Thread?

In computer science, a thread of execution is the smallest sequence of programmed instructions that can be managed independently by a scheduler, which is typically a part of the operating system.
 The implementation of threads and processes differs between operating systems, but in most cases **a thread is a component of a process**
<br />

## Thread VS Process

*	rocesses are typically independent, while threads exist as subsets of a process
*	processes carry considerably more state information than threads, whereas multiple threads within a process share process state as well as memory and other resources
*	processes have separate address spaces, whereas threads share their address space
*	processes interact only through system-provided inter-process communication mechanisms
*	context switching between threads in the same process typically occurs faster than context switching between processes



### Sources
[wikipedia Thread_](https://en.wikipedia.org/wiki/Thread_\(computing\))
<br />
[An Intro to Threading in Python (realpython)](https://realpython.com/intro-to-python-threading)
<br />
[Multithread in Python | set1 (geekforgeeks)](https://www.geeksforgeeks.org/multithreading-python-set-1/)
<br />
[threading — Manage Concurrent Operations Within a Process (pymotw)](https://pymotw.com/3/threading/)
<br />

