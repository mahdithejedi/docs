# What is Weakrefrence?

***Some prerequisites***:
<br />
**What is Unreachble memory?**:
In computer science, unreachable memory is a block of memory allocated dynamically where the program that allocated the memory no longer has any reachable pointer that refers to it
[1](https://en.wikipedia.org/wiki/Unreachable_memory)
<br />
**And what is Refrenc counting**
In computer science, reference counting is a programming technique of storing the number of references, pointers, or handles to a resource, such as an object, a block of memory, disk space, and others.
[2](https://en.wikipedia.org/wiki/Reference_counting)
<br />


So finally what is weakrefrence?

**In computer programming, a weak reference is a reference that does not protect the referenced object from collection by a garbage collector, unlike a strong reference.. An object referenced only by weak references  meaning "every chain of references that reaches the object includes at least one weak reference as a link" – is considered weakly reachable, and can be treated as unreachable and so may be collected at any time.**
[3](https://en.wikipedia.org/wiki/Weak_reference)


# Weakrefrence in PYTHON(3.8)

The weakref module allows the Python programmer to create weak references to objects.
<br />
In the following, the term referent means the object which is referred to by a weak reference.
<br />
A weak reference to an object is not enough to keep the object alive: when the only remaining references to a referent are weak references, garbage collection is free to destroy the referent and reuse its memory for something else. However, until the object is actually destroyed the weak reference may return the object even if there are no strong references to it.
<br />
A primary use for weak references is to implement caches or mappings holding large objects, where it’s desired that a large object not be kept alive solely because it appears in a cache or mapping.
<br />
[4](https://docs.python.org/3/library/weakref.html#module-weakref)

<br /><br />

## More Information
[pymotw 3](https://pymotw.com/3/weakref/)
<br />

## Notes

1. `weakref.fanlize` is used  when we want to clean up a obj!

