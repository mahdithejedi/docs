## What Even is a Syscall?

A system call is a special procedure that lets a program start a transition from user space to kernel space, jumping
from the program’s code into OS code.

# Basics recap

* Processors execute instructions in an infinite fetch-execute loop and don’t have any concept of operating systems or
  programs. The processor’s mode, usually stored in a register, determines what instructions may be executed. Operating
  system code runs in kernel mode and switches to user mode to run programs.
* To run a binary, the operating system switches to user mode and points the processor to the code’s entry point in RAM.
  Because they only have the privileges of user mode, programs that want to interact with the world need to jump to OS
  code for help(for example need resources or call `read` or `write` on a file). System calls are a standardized way for
  programs to switch from user mode to kernel mode and into OS code and then return to user mode(this can happen
  by  [Interrupt](https://en.wikipedia.org/wiki/Interrupt) )
* Programs typically use these syscalls by calling shared library functions. These wrap machine code for either software
  interrupts or architecture-specific syscall instructions that transfer control to the OS kernel and switch rings. The
  kernel does its business and switches back to user mode and returns to the program code.

# slice dat time

## Preemption (computing)

In computing, preemption is the act of temporarily interrupting an executing task, with the intention of resuming it at
a later time. This interrupt is done by an external scheduler with no assistance or cooperation from the task.
<br />
In any given system design, some operations performed by the system may not be preemptable. This usually applies to
kernel functions and service interrupts which, if not permitted to run to completion, would tend to produce race
conditions resulting in deadlock. Barring the scheduler from preempting tasks while they are processing kernel functions
simplifies the kernel design at the expense of system responsiveness. The distinction between user mode and kernel mode,
which determines privilege level within the system, may also be used to distinguish whether a task is currently
preemptable.

### Preemptive multitasking

The term preemptive multitasking is used to distinguish a multitasking operating system, which permits preemption of
tasks, from a cooperative multitasking system wherein processes or tasks must be explicitly programmed to yield when
they do not need system resources.
<br />
In general, preemption means "prior seizure of". When the high-priority task at that instance seizes the currently
running task, it is known as preemptive scheduling.
<br />
Preemptive multitasking allows the computer system to more reliably guarantee each process a regular "slice" of
operating time. It also allows the system to rapidly deal with important external events like incoming data, which might
require the immediate attention of one or another process.

# How to run a program (Linux Os on X86 architecture)

## Execve or Execveat

when you run a program, linux will call a function [`execve`](https://man7.org/linux/man-pages/man2/execve.2.html). this
function will get following arguments:
<br />
`int execve(const char *filename, char *const argv[], char *const envp[]);`

* The __filename__ argument specifies a path to the program to run.
* __argv__ is a null-terminated (meaning the last item is a null pointer) list of arguments to the program. The _argc_
  argument you’ll commonly see passed to C main functions is actually calculated later by the syscall, thus the
  null-termination.
* The __envp__ argument contains another null-terminated list of environment variables used as context for the
  application. They’re… conventionally __KEY=VALUE__ pairs. Conventionally. I love computers.

## do_execveat_common

If you call `execveat` then os call another function by the name of `do_execveat_common` which
make [`linux_binprm`](https://github.com/torvalds/linux/blob/22b8cc3e78f5448b4c5df00303817a9137cd663f/include/linux/binfmts.h#L15-L65)
struct. There are some points which is important in this struct

* `mm_struct` and `vm_area_struct` are defined to prepare virtual memory management for the new program.
* _argc_ and _envc_ are calculated and stored to be passed to the program.
* _filename_ and _interp_ store the filename of the program and its interpreter, respectively. These start out equal to
  each other, but can change in some cases: one such case occurs when the binary being executed is different from the
  program name is when running interpreted programs like Python scripts with a shebang. In this example, _filename_
  points to the Python file but the _interp_ is the Python interpreter’s path.
* _buf_ is an array filled with the first 256 bytes of the file to be executed. It’s used to detect the format of the
  file and load script shebangs.

## Binfmts

The kernel’s next major job is iterating through a bunch of “binfmt” (binary format) handlers. These handlers are
defined in files like _
fs/binfmt_elf.c_ and _fs/binfmt_flat.c_. Kernel modules can also add their own binfmt handlers to the pool.

Each handler exposes a _load_binary()_ function which takes a _linux_binprm_ struct and checks if the handler
understands the program’s format.<small> Either by _buf_ struct of format of file </small>

# Becoming an Elf-Lord

the kernel will reach a final program containing machine code for it to launch. Typically, a setup process is required
before actually jumping to the code
<br />
we have standard file formats that specify how to set up a program for execution. While Linux supports many such
formats, the most common format by far is ELF
<br />
ELF binaries are handled by the _binfmt_elf_ handler and It’s responsible for parsing out certain details from the ELF
file and using them to load the process into memory and execute it.
<br />

![Struct of an ELF file](https://cpu.land/images/elf-file-structure.png)

## ELF Header

contains data in which:

* What processor it’s designed to run on. ELF files can contain machine code
* Whether the binary is meant to be run on its own as an executable, or whether it’s meant to be loaded by other
  programs as a “dynamically linked library.”
* The entry point of the executable. <small>  The entry point is a memory address pointing to where the first machine
  code instruction is in memory after the entire process has been loaded. </small>

## Program header table

![Common program header types](https://cpu.land/images/elf-program-header-types.png)

* It points to the position of its data within the ELF file.
* It can specify what virtual memory address the data should be loaded into memory at.
* Two fields specify the length of the data: one for the length of the data in the file, and one for the length of the
  memory region to be created. If the memory region length is longer than the length in the file, the extra memory will
  be filled with zeroes. This is beneficial for programs that might want a static segment of memory to use at runtime;
  these empty segments of memory are typically called BSS segments.
* Finally, a flags field specifies what operations should be permitted if it’s loaded into memory: `PF_R` -> readable, _
  PF_W_ -> writable, `PF_X` -> allowed to execute on the CPU.

[load_elf_library function in kernel where elf file is loaded](https://github.com/torvalds/linux/blob/18b44bc5a67275641fb26f2c54ba7eef80ac5950/fs/binfmt_elf.c#L1368)
<br />
[part about getting virtual memory in load_elf_libray](https://github.com/torvalds/linux/blob/18b44bc5a67275641fb26f2c54ba7eef80ac5950/fs/binfmt_elf.c#L1417C9-L1424C38)

## Data

![](https://cpu.land/images/elf-data-section.png)

The program and section header table entries all point to blocks of data within the ELF file, whether to load them into
memory, to specify where program code is, or just to name sections. All of these different pieces of data are contained
in the data section of the ELF file.

## static vs dynamic linking

![static vs dynamic linking](https://cpu.land/images/static-vs-dynamic-linking.png)
If a **statically linked** program needs a function _foo_ from a library called _bar_, the program would include a copy
of the entirety of _foo_. However, if it’s **dynamically linked** it would only include a reference saying “I need _foo_
from library
_bar_.” When the program is run, _bar_ is hopefully installed on the computer and the _foo_ function’s machine code can
be loaded into memory on-demand
<br />
An interesting distinction between the two types of linking is that with static linking, only the portions of the
library that are used are included in the executable and thus loaded into memory. With dynamic linking, the entire
library is loaded into memory. This might initially sound less efficient, but it actually allows modern operating
systems to save more space by loading a library into memory once and then sharing that code between processes. Only code
can be shared as the library needs different state for different programs, but the savings can still be on the order of
tens to hundreds of megabytes of RAM.

## Execution

Let’s hop on back to the kernel running ELF files: if the binary it’s executing is dynamically linked, the OS can’t just
jump to the binary’s code right away because there would be missing code — remember, dynamically linked programs only
have references to the library functions they need!
<br />
To run the binary, the OS needs to figure out what libraries are needed, load them, replace all the named pointers with
actual jump instructions, and then start the actual program code.
<br />
After reading the ELF header and scanning through the program header table, the kernel can set up the memory structure
for the new program. It starts by loading all _PT_LOAD_ segments into memory, populating the program’s static data, BSS
space, and machine code. If the program is dynamically linked, the kernel will have to execute the ELF interpreter (
PT_INTERP), so it also loads the interpreter’s data, BSS, and code into memory.
<br />
The kernel is almost ready to return from the syscall (remember, we’re still in _execve_). It pushes the _argc_, _argv_,
and environment variables to the stack for the program to read when it begins.
<br />
Finally, the syscall is over and the kernel returns to userland. It restores the registers, which are now zeroed, and
jumps to the stored instruction pointer. That instruction pointer is now the starting point of the new program (or the
ELF interpreter) and the current process has been replaced!

# Fork

Anyways, Unix programs launch new programs by calling __fork__ and then immediately running __execve__ in the child
process. This is called the _fork-exec_ pattern.
<br />
You might’ve noticed that duplicating a process’s memory only to immediately discard all of it when loading a different
program sounds a bit inefficient. Luckily, we have an MMU. Duplicating data in physical memory is the slow part, not
duplicating page tables, so we simply don’t duplicate any RAM: we create a copy of the old process’s page table for the
new process and keep the mapping pointing to the same underlying physical memory.
<br />
Introducing COW (copy on write) pages. With **COW pages**, both processes read from the same physical addresses as long
as they don’t attempt to write to the memory
<br />
COW is implemented, like many fun things, with paging hacks and hardware interrupt handling. After fork clones the
parent, it flags all of the pages of both processes as read-only. When a program writes to memory, the write fails
because the memory is read-only. This triggers a segfault (the hardware interrupt kind) which is handled by the kernel.
The kernel which duplicates the memory, updates the page to allow writing, and returns from the interrupt to reattempt
the write.

## Sources

[Wikipedia Preemption](https://en.wikipedia.org/wiki/Preemption_(computing))
<br />
