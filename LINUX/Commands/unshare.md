# `unshare` command
The unshare command creates new namespaces (as specified by the
command-line options described below) and then executes the
specified program. If program is not given, then "${SHELL}" is
run (default: /bin/sh).

One key feature or quirk of unshare is that, by default, it runs a program with the new namespaces created, but it does not associate that program with these namespaces. Instead, the new namespaces go to any children that program creates. You can add a –fork option to make it work more as you’d expect.

For example, let’s start a new shell in its own private:
`sudo unshare --pid --fork --mount-proc /bin/bash`

**Point1: In you run this command without `--fork` option you will face an error. [more detail](https://stackoverflow.com/a/45973522)**
<br />
**Point2: If you run this command without `--mount-proc` again you will face error. [more detail](https://stackoverflow.com/a/68706102)**


### Sources
[LINUX FU: DON’T SHARE WELL WITH OTHERS](https://hackaday.com/2021/12/28/linux-fu-dont-share-well-with-others/)
<br />
[man7 page](https://man7.org/linux/man-pages/man1/unshare.1.html)
<br />
