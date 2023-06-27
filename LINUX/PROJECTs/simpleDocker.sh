#!/bin/bash

## Source: https://akashrajpurohit.com/blog/build-your-own-docker-with-linux-namespaces-cgroups-and-chroot-handson-guide/
## Hckrnews discuess: https://news.ycombinator.com/item?id=36488356

function _run(){

unshare --uts --pid --net --mount --ipc --fork

mkdir /sys/fs/cgroup/cpu/"$1"
echo 100000 > /sys/fs/cgroup/cpu/"$1"/cpu.cfs_quota_us
echo 0 > /sys/fs/cgroup/cpu/"$1"/tasks
echo $$ > /sys/fs/cgroup/cpu/"$1"/tasks

debootstrap focal ./ubuntu-rootfs http://archive.ubuntu.com/ubuntu/


mount -t proc none ./ubuntu-rootfs/proc
mount -t sysfs none ./ubuntu-rootfs/sys
mount -o bind /dev ./ubuntu-rootfs/dev
chroot ./ubuntu-rootfs /bin/bash
}

read -p "Enter namespace name:" containerName
_run "$containerName"