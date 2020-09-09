# how to face with : Could not get lock /var/lib/dpkg/lock

## first

find the process which locked the apt the stop it

`sudo pidof apt`
`sudo ps -aux | grep -i apt`

then you will get the process id which you can kill it with

`sudo kill <PROCCESS_ID>`
if this can not kill the proccess use

`sudo kill -9 <PROCCESS_ID>`


## second

you have to see which proccess is using lock files with **lsof**

check with command bellow

`sudo lsof /var/lib/dpkg/lock`
<br />
`sudo lsof /var/lib/apt/lists/lock`
<br />
`sudo lsof /var/cache/apt/archives/lock`

then you can get proccess id and kill it and finally you can safely remove 
files

`sudo rm /var/lib/dpkg/lock`
<br />
`sudo rm /var/lib/apt/lists/lock`
<br />
`sudo rm /var/cache/apt/archives/lock`


and then reconfig the package manager

`sudo dpkg --configure -a`



## other ways

* **1**
```
sudo lsof /var/lib/dpkg/lock-frontend
sudo kill -9 PID
sudo rm /var/lib/dpkg/lock-frontend
sudo apt update
```

*   **2**
```
sudo lsof /var/lib/dpkg/lock-frontend
sudo kill -9 PID
sudo rm /var/lib/dpkg/lock-frontend
sudo dpkg --configure -a
```