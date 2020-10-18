First of all
<br />
**[What is exec?](../Command/exec.md)**:
The exec() family of functions ***replaces*** the current process image  with a new process image
<br />
in the other hand
<br />
**[Fork system call](../Command/fork.md)**: fork() creates a ***new*** process by duplicating the calling process
<br />
## So the diffrence is
*	fork starts a new process which is a copy of the one that calls it, while exec replaces the current process image with another (different) one.
*	Both parent and child processes are executed simultaneously in case of fork() while Control never returns to the original program unless there is an exec() error.

### What is a thread in a process call one of Fork or Exec?
**Does all the thread dulpicate in fork and all thread replace in new process of Exec?**
<br />
in UNIX systems we have two kind of forks. The one which duplicate all threads and the one which just duplicate this process
<br />
But in exec if a thread invokes the exec() system call, the program specified in the parameter to exec() will replace the entire process—***ncluding all threads.***



### Resources
[Difference between fork() and exec() (geekforgeeks)](https://www.geeksforgeeks.org/difference-fork-exe)
<br />

