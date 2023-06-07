# The process
Informally, as mentioned earlier, a process is a program in execution. The status of the current activity of a process is represented by the value of the program counter and the contents of the processor’s registers. The memory layout of a process is typically divided into multiple sections
<br />

*	**Heap section**: memory that is dynamically allocated during program run time
<br />
* 	**Stack section**: temporary data storage when invoking functions (such as function parameters, return addresses, and local variables)
<br />

Notice that the sizes of the text and data sections are fixed, as their sizes do not change during program run time. However, the stack and heap sections can shrink and grow dynamically during program execution. Each time a function is called, an [activation record](../../../PYTHON/PythonInDeep/ActivationRecord.md) containing function parameters, local variables, and the return address is pushed onto the stack; when control is returned from the function, the activation record is popped from the stack. Similarly, the heap will grow as memory is dynamically allocated, and will shrink when memory is returned to the system. Although the stack and heap sections grow toward one another, the operating system must ensure they do not overlap one another.
<br />
**A program becomes a process when an executable file is loaded into memory**
<br />
