## & :

run a process as background eg:`xeyes &`

## jobs

show all background jobs

## Ctrl + z

this command will **stop** the proccess 

## bg

this command will make a bg  process to start in background(which may be stopped with `ctrl + z` ) 
<br />
eg: `bg %3` which will start the third process in bg

## fg

move from bg to foreground 
<br />
eg: `fg %2` foreground third background process

## nohup

when you close a terminal all the process will be kill so for preventing this problem we will use **nohub <COMMAND> <XARGS>**
<br />
eg:`nohup ping google.com`
<br />
<br />
This command means do not send **hup** [signal](https://www.computerhope.com/unix/signals.htm)