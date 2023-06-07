# What is VirtualMemory in Linux
Virtual Memory is a storage allocation scheme in which secondary memory can be addressed as though it were part of the main memory. The addresses a program may use to reference memory are distinguished from the addresses the memory system uses to identify physical storage sites, and program-generated addresses are translated automatically to the corresponding machine addresses. 
<br />
Why we even need that?
*   All memory references within a process are logical addresses that are dynamically translated into physical addresses at run time. This means that a process can be swapped in and out of the main memory such that it occupies different places in the main memory at different times during the course of execution.
*   A process may be broken into a number of pieces and these pieces need not be continuously located in the main memory during execution. The combination of dynamic run-time address translation and use of page or segment table permits this.


## Sources
[Geeksforgeeks](https://www.geeksforgeeks.org/virtual-memory-in-operating-system/)
<br />
[Video Course (Greybeard Qualification conference)](https://youtu.be/CPNdBCdmvHs)



