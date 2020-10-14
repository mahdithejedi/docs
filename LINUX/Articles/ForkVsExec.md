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

### Resources
[Difference between fork() and exec() (geekforgeeks)](https://www.geeksforgeeks.org/difference-fork-exe)
<br />

