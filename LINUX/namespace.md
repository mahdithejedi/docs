# What is namespace?

Namespaces are a feature of the Linux kernel that partitions kernel resources such that one set of processes sees one
set of resources while another set of processes sees a different set of resources.

Namespace functionality is the same across all kinds: each process is associated with a namespace and can only see or
use the resources associated with that namespace, and descendant namespaces where applicable.

## Namespace kinds

### Mount (mnt)

Mount namespaces control mount points. Upon creation the mounts from the current mount namespace are copied to the new
namespace, but mount points created afterwards do not propagate between namespaces

**_The clone flag used to create a new namespace of this type is CLONE_NEWNS - short for "NEW NameSpace"._**

### Process ID (pid)

The PID namespace provides processes with an independent set of process IDs (PIDs) from other namespaces.

### Network (net)

Network namespaces virtualize the network stack. On creation, a network namespace contains only a loopback interface.

Each network interface (physical or virtual) is present in exactly 1 namespace and can be moved between namespaces.

Each namespace will have a private set of IP addresses, its own routing table, socket listing, connection tracking
table, firewall, and other network-related resources.

Destroying a network namespace destroys any virtual interfaces within it and moves any physical interfaces within it
back to the initial network namespace.

### Inter-process Communication (ipc)

### User ID (user)

User namespaces are a feature to provide both privilege isolation and user identification segregation across multiple
sets of processes
<br />
With administrative assistance it is possible to build a container with seeming administrative rights without actually
giving elevated privileges to user processes. Like the PID namespace, user namespaces are nested and each new user
namespace is considered to be a child of the user namespace that created it.

### Control group (cgroup) Namespace

The cgroup namespace type hides the identity of the control group of which process is a member. A process in such a
namespace, checking which control group any process is part of, would see a path that is actually relative to the
control group set at creation time, hiding its true control group position and identity. 

### Time Namespace

# Handling namespace commands
### [`Unshare` command](./Commands/unshare.md)


# Resources
[Wikipedia namespace](https://en.wikipedia.org/wiki/Linux_namespaces)
<br />

