# The process

Informally, as mentioned earlier, a process is a program in execution. The status of the current activity of a process
is represented by the value of the program counter and the contents of the processor’s registers. The memory layout of a
process is typically divided into multiple sections
<br />

* **Heap section**: memory that is dynamically allocated during program run time
  <br />
* **Stack section**: temporary data storage when invoking functions (such as function parameters, return addresses, and
  local variables)
  <br />

Notice that the sizes of the text and data sections are fixed, as their sizes do not change during program run time.
However, the stack and heap sections can shrink and grow dynamically during program execution. Each time a function is
called, an [activation record](../../../PYTHON/PythonInDeep/ActivationRecord.md) containing function parameters, local
variables, and the return address is pushed onto the stack; when control is returned from the function, the activation
record is popped from the stack. Similarly, the heap will grow as memory is dynamically allocated, and will shrink when
memory is returned to the system. Although the stack and heap sections grow toward one another, the operating system
must ensure they do not overlap one another.
<br />
**A program becomes a process when an executable file is loaded into memory**
<br />

## The process memory layout

•   **Text section**:  the executable code

•   **Data section** : global variables

• **Heap section** : memory that is dynamically allocated during program run time

• **Stack section** : temporary data storage when invoking functions (such as function parameters, return addresses, and
local variables)

## States of memory

![process state of memory](https://www.baeldung.com/wp-content/uploads/sites/4/2021/01/5-3.png)

•   **New** The process is being created.

• **Running**: Instructions are being executed.

• **Waiting**. The process is waiting for some event to occur (such as an I/O completion or reception of a signal).

• **Ready**. The process is waiting to be assigned to a processor.

• **Terminated**. The process has finished execution

## Process scheduling

the process scheduler selects an available process (possibly from a set of several available processes)
for program execution on a core.
<br />

### process in linux

For example, the state of a process is represented by the field long state in this structure. Within the Linux kernel,
all active processes are represented using a doubly linked list of task struct . The kernel maintains a pointer –
current – to the process currently executing on the system. For more information about linux kernel process refer
to [Kernel/process.md](../../Kernel/process.md)

### Scheduling Queues

As processes enter the system, they are put into a _ready queue_, where they are ready and waiting to execute on a CPU
’s core This queue is generally stored as a linked list; a ready-queue header contains pointers to the first PCB in the
list, and each PCB includes a pointer field that points to the next PCB in the ready queue. The system also includes
other queues. When a process is allocated a CPU core, it executes for a while and eventually terminates, is interrupted,
or waits for the occurrence of a particular event, such as the completion of an I/O request. Suppose the process makes
an I/O request to a device such as a disk. Since devices run significantly slower than processors, the process will have
to wait for the I/O to become available. Processes that are waiting for a certain event to occur — such as completion of
I/O — are placed in a wait queue

![process queues](https://www.cs.uic.edu/~jbell/CourseNotes/OperatingSystems/images/Chapter3/3_06_QueueingDiagram.jpg)

• The process could issue an I/O request and then be placed in an I/O wait queue.

• The process could create a new child process and then be placed in a wait queue while it awaits the child’s
termination.

• The process could be removed forcibly from the core, as a result of an interrupt or having its time slice expire, and
be put back in the ready queue.

## Cpu scheduling

A process migrates among the ready queue and various wait queues through- out its lifetime. The role of the CPU
scheduler is to select from among the processes that are in the ready queue and allocate a CPU core to one of them

## Context switch

de, and then a state restore to resume operations. Switching the CPU core to another process requires performing a state
save of the current process and a state restore of a different process. This task is known as a
_context switch_
<br />
Context-switch time is pure overhead, because the system does no useful work while switching

## Process Creation

In linux a new process is created by `fork()` system call, After that linux will copy the memory of that process and
then make a new section in memory table of the process and point to that.

But actually linux doesn't copy, However the pointer of memory of two processes point to the same slot of memory but
then child process want to make changes then Linux will copy the memory section to another place in memory and separate
their memory pointer

The child process inherits privileges and scheduling attributes from the parent, as well certain resources, such as open
files.

![Process creation flow using Fork](https://www.cs.csustan.edu/~john/Classes/CS3750/Notes/Chap03/3_09fork().jpg)

# Interprocess Communication

Processes executing concurrently in the operating system may be either inde- pendent processes or cooperating processes.
A process is independent if it does not share data with any other processes executing in the system. A process is
cooperating if it can affect or be affected by the other processes executing in the system. Clearly, any process that
shares data with other processes is a cooperating process. There are several reasons for providing an environment that
allows process cooperation:

* **Information sharing**. Since several applications may be interested in the same piece of information (for instance,
  copying and pasting), we must provide an environment to allow concurrent access to such information.
* **Computation speedup**. If we want a particular task to run faster, we must break it into subtasks, each of which
  will be executing in parallel with the others. Notice that such a speedup can be achieved only if the computer has
  multiple processing cores.
* **Modularity**. We may want to construct the system in a modular fashion, dividing the system functions into separate
  processes or threads.

Cooperating processes require an interprocess communication ( IPC )

## Model in interprocess communication

* **Shared Memory**
* **Message passing**

### IPC in shared memory system

Typically, a shared-memory region resides in the address space of the process creating the shared-memory segment. Other
processes that wish to communicate using this shared-memory segment must attach it to their address space. Recall that,
normally, the operating system tries to prevent one process from accessing another process’s memory. Shared memory
requires that two or more processes agree to remove this restriction. They can then exchange information by reading and
writing data in the shared areas. The form of the data and the location are determined by these processes and are not
under the operating system’s control. The processes are also responsible for ensuring that they are not writing to the
same location simultaneously.

### IPC in message passing system

Message passing provides a mechanism to allow processes to communicate and to synchronize their actions without sharing
the same address space. It is particularly useful in a distributed environment, where the communicating processes may
reside on different computers connected by a network

If processes P and Q want to communicate, they must send messages to and receive messages from each other: a
communication link must exist between them.

Type of link communication:

* Direct or indirect communication
* Synchronous or asynchronous communication
* Automatic or explicit buffering

#### Naming

Processes that want to communicate must have a way to refer to each other. They can use either direct or indirect
communication.

Under _direct communication_, each process that wants to communicate must explicitly name the recipient or sender of the
communication

* A link is associated with exactly two processes.
* Between each pair of processes, there exists exactly one link.

There is another model in which only the sender names the recipient; the recipient is not required to name the sender.

* _send(P, message)_ : Send a message to process P .
* _receive(id, message)_ : Receive a message from any process. The vari- able id is set to the name of the process with
  which communication has taken place.

With _indirect communication_, the messages are sent to and received from **mailboxes**, or ports. A mailbox can be
viewed abstractly as an object into which messages can be placed by processes and from which messages can be removed.
Each mailbox has a unique identification.

## Buffering

Whether communication is direct or indirect, messages exchanged by communicating processes reside in a temporary queue.
Basically, such queues can be implemented in three ways:

* **Zero capacity**. The queue has a maximum length of zero; thus, the link cannot have any messages waiting in it. In
  this case, the sender must block until the recipient receives the message.
* **Bounded capacity**. The queue has finite length n; thus, at most n messages can reside in it.
* **Unbounded capacity**. The queue’s length is potentially infinite; thus, any number of messages can wait in it. The
  sender never blocks.

# Threads

Process creation is time consuming and resource intensive, however. If the new process will perform the same tasks as
the existing process, why incur all that overhead? It is generally more efficient to use one process that contains
multiple threads

### Benefits of multithread

* Responsiveness
* Resource sharing
* Economy
* Scalability

### Concurrency vs Parallelism

A concurrent system supports more than one task by allowing all the tasks to make progress. In contrast, a parallel
system can perform more than one task simultaneously

## Parallelism Type

* **Data parallelism** : ~ focuses on distributing subsets of the same data across multiple computing cores and
  performing the same operation on each core
* **Task parallelism** : ~ involves distributing not data but tasks (threads) across multiple computing cores. Each
  thread is performing a unique operation. Different threads may be operating on the same data, or they may be operating
  on different data.

### User thread vs Kernel Thread

User threads are supported above the kernel and are managed without kernel support, whereas kernel threads are supported
and managed directly by the operating system.

## Multithread models

* **Many-to-One**: The many-to-one model (Figure 4.7) maps many user-level threads to one kernel thread. However, the
  entire process will block if a thread makes a blocking system
  call. ![many to one thread](https://www.cs.uic.edu/~jbell/CourseNotes/OperatingSystems/images/Chapter4/4_05_ManyToOne.jpg)

* **One-to-One**: The one-to-one model (Figure 4.8) maps each user thread to a kernel thread. It provides more
  concurrency than the many-to-one model by allowing another thread to run when a thread makes a blocking system call.
  It also allows multiple threads to run in parallel on multiprocessors. The only drawback to this model is that
  creating a user thread requires creating the corresponding kernel thread, and a large number of kernel threads may
  burden the performance of a system. _Linux, along with the family of Windows operating systems, imple- ment the
  one-to-one model._
  ![one to one thread](https://www.cs.uic.edu/~jbell/CourseNotes/OperatingSystems/images/Chapter4/4_06_OneToOne.jpg)

* **Many-to-Many**: The many-to-many model (Figure 4.9) multiplexes many user-level threads to a smaller or equal number
  of kernel
  threads ![many to many threads](https://www.cs.uic.edu/~jbell/CourseNotes/OperatingSystems/images/Chapter4/4_07_ManyToMany.jpg)

### Many-to-one VS one-to-one

many-to-one model allows the developer to create as many user threads as she wishes, it does not result in parallelism,
because the kernel can schedule only one kernel thread at a time. The one-to-one model allows greater concurrency, but
the developer has to be careful not to create too many threads within an application.

### Thread context-switch overhead

There are as many answers to this as there are operating system versions. No single answer covers everything.
<br />

Basically, the cost of context switching is the cost of saving all of the cpu state relating to the process context, and
then loading in the context of a new process.
<br />

What exactly is saved is highly dependent on not just the operating system but the cpu hardware itself. For example,
processors like the Intel cpus have lots of registers that have to be saved somewhere and then reloaded with the other
process's context, while the sparc cpu keeps most of its context including all cpu registers on the stack, so a context
switch is a matter if just moving the stack pointer to a different register window.
<br />

On the other hand, most modern cpus have some of their state in cpu cache memory, and while this is not typically
swapped during a context switch, use of memory can cause cache lines to be unloaded and reloaded as the new process
executes, so that when switching back to the previous process, its cache lines may need to be reloaded. While this isn't
a direct cost of context switching, it still is there.
<br />

There are many other resources in the cpu that need to be switched, such as page maps, permission bits, etc., and the
list is different on each cpu model and handled differently by each operating system.
<br />

You reference **processes vs. threads**. At one time, there were vast differences between them, and threads had much less
context that needed switching. Then threads as light weight processes were created, so now there is very little
difference, and processes are nearly as light as threads were.

# Linux Kernel Implementation

the [`schedule`](https://github.com/torvalds/linux/blob/1c41041124bd14dd6610da256a3da4e5b74ce6b1/kernel/sched/core.c#L6768)
pick the next task to run. During this process , linux get
the [`current`](https://github.com/torvalds/linux/blob/1c41041124bd14dd6610da256a3da4e5b74ce6b1/kernel/sched/core.c#L6770)
task to run then check if the task is not running then commit the task via `sched_submit_work`. then
the [`sched_submit_work`](https://github.com/torvalds/linux/blob/1c41041124bd14dd6610da256a3da4e5b74ce6b1/kernel/sched/core.c#L6712C20-L6712C37)
function _assign_ a task to a worker and notify corresponding worker that the task is assigned to that worker