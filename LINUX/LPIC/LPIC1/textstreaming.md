##  cat

this can joint two files

**`-`** this will stay for input ex:
<br />

`cat myfile - output.txt` first show myfiles then will add STDIN until you press CTRL+D and then print output.txt


## split

will split the input to several output

`split -l10 `:This will split every 10 line to a new file\
<br />

`split -b40 ./txt.txt splited` will split 40 byte 40 byte and set the name of outputs splitedaa, splitedab ... and if we want to show splited1, splited2 instead of splitedaa,... we have to use `split -b40 -d ./txt.txt splited`
<br />

## tail -f 

:)

## pr
will print :)

## nl == cat -n

put line number in the first of every line

## fmt 

this will fomat input like we want to join the lines together

## sort

`-n` will sort numeric


## uniq

this tool work better with `sort`

`-c`:count
<br />

`-d`:show the items which themself are not uniq 
<br />

`u`: show real uniq items(which themself are uniq)


## cut

`-f`:which coloumn to show(f=field)
<br />

`-d`: what make the cut to space (default is TAB )eg:

`-c`

`cut -f2 -d,` seprate with comoa
<br />

`cut -f2 -d" "` sperate with space


## paste

## join

## sed 

## tr