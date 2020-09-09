## SOME WORDS

HOTPLUG = when the system is on 

COLDPLUG = when the system is off


## process of installing a new hardware to linux

first of all the **HAL** will abstract the name of that hardware so in every linux system you can have same name like for wireless you always have the same name for it and in every device you can have the same conf

***Then***

the system will tell to **debus** to start cominication with eachother like you mount a mouse and debus tell to GUI to comiunicate with mouse

***Then***

<a name="udev"></a>**udev** can set rule for the mounted device eg you want to say to system set the name of a specific device like external hard Adata to /dev/sab 

## some commend

lsmod, lspci, lsusb, lshw, lsmod,rmmod,modprob

## /etc/modprobe.d/ direcotry

this commend will auto load the apps(the conf of that app) which inside this folder when system starts the **older** way to do this is /etc/modules file

## UUID
every device have a UUID 

connected UUID of devices! `blkid` 