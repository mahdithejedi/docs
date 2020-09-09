## refrences

[Linux boot and startup from opensource](https://opensource.com/article/17/2/linux-boot-and-startup)

[dmesg command](../../commands/dmesg.md)


## GRUP 

config of grup is located in **/boot/grub**

## init
init process **/sbin/init**  acctually the init is a shortcut to **/lib/system/systemd**

`pstree` will show the process tree

### recetly linuxes are switch to upstart and systemd
### point: in systemd every thing is a unit like socker, device,...

## dmesg

log all the boot system 

## /var/log/syslog

log of all the programms

## runlevels

0 - halt
1 -  single user mode (recovery)
2 - Debian/Ubuntu default
3 - RHEL/Fedora/Suse text mode
4 - free
5 - R/F/S graphical mode
6 - reboot

`runlevel` show the current run level

`telinit` will change run level