# Profiling

**Every time something in the code bites you, add a test**

*Add asert checks are a simple way to add some level of validation*

## Cpu profiling

### Tools

* CPU & TIME
    *  Use python `print` and `decorator` to measure time usage
    * Python [timeit](https://docs.python.org/3/library/timeit.html) module(_The timeit module temporarily disables the garbage collector. This might impact the speewd you'll see with the realworld_)
    * Linux [time](https://man7.org/linux/man-pages/man1/time.1.html) module
    * python [cProfile](https://docs.python.org/3/library/profile.html)
        * visualize `cProfile` output with [snakeviz](https://jiffyclub.github.io/snakeviz/)
    * python [profile.profile](https://docs.python.org/3/library/profile.html)
    * python [line_profile](https://github.com/rkern/line_profiler)
* Memory
    * Python [memory-profiler](https://pypi.org/project/memory-profiler/)


## Introspecting an Existing Process with PySpy

## Warpup
Some Point:
1. Seprate the section of code you want to profile and make sure you have accurate test for that part of code
2. disable any BOIS-based acceleratos
3. Try to hypothesize the expected behavior of your code and then validate the hypothesis with the result of a profiling step
4. For web servers, investigate _dower_ and _dozer_
5. user `coverage.py` to confirm that your tests are covering all the code paths

