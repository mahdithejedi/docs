# how to detach a shell process?

## start
some shell proceess have STDIN and STDOUT and we have to find a way to save put STDIN and read STDOUT from a source so how?
<br /> <br />
`COMMAND </PATH/TO/REDIRECT &> /PATH/TO/REDIRECT &`

or we can say

`COMMAND </PATH/TO/READ &> /PATH/TO/WRITE &`

### point: you can redirect the pathes to /dev/null where nothing will save

## kill
you can see the running shell with `jobs` and kill them with `kill %1`

## links

https://superuser.com/questions/178587/how-do-i-detach-a-process-from-terminal-entirely

https://www.tecmint.com/run-linux-command-process-in-background-detach-process/

https://explainshell.com/explain?cmd=ping+%3C%2Fdev%2Fnull+%26%3E+%2Fdev%2Fnull+%26

https://unix.stackexchange.com/questions/104821/how-to-terminate-a-background-process

https://explainshell.com/explain?cmd=ps+-eaf+%7C+grep+%5Bw%5Dget





