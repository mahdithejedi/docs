## 2 types of OS Communication
* massage passing model:
  * useful for smaller data because no conflict should be avoided
  * easier to develop
* shared-memory model: 
  * in this way process call `shared_memory_create` and `shared_memoeyr_atach`
   to gain access. 
  * OS prevent 2 processes access each other memory shared.
  * If processes want to remove some info all of them should agree
  * better for more data exchange with max speed

## Linker and loader
* linker: the _linker_ combines __relocatable object fil__ into a single binary file
* loader: a _loader_ load the binary into memory so eligible to run on a CPU core