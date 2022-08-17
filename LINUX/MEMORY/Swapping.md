# Swapping
Swapping a process out means removing all of its pages from memory, or marking them so that they will be removed by the normal page replacement process

## Commands
`mdswap` -> create swap
<br />
`swapon` -> Enable or Disable swap
<br />
`vmstat` -> VM stat
<br />
`/proc/meminfo` -> Memory Info
<br />
`/proc/<pid>/smaps` -> Memory Info of app
=> More Info

*   RSS -> How much if this segment is present in memory

**Point: id=82 label disk automatically considered as swap**


