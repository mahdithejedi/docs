## redirect output
with `>` symboul you can redirect a result of a script or cammand to another place
wg:
`ls > text.tx` <br />
this commend will insert ls commend into **text.tx** 

the `>` will override tx and will remove all the text.tx data for adding data into endof text.tx you can use *`>>`*
eg:
`ls >> text.tx`

## redirect output
with `<` you can redirect output 
eg:
`echo < ./users` <br />
this will redirect any data from ./users to echo 

another eg:
`tr ' ' ',' < pidof Telegram` <br />

*what is `<<`?*
this is like ''' in python eg:
`echo << END`

this will echo any inputed data **untill the `END` is intered**
like
```
echo << END
>This
>is
>a
>test
>END
```
output will be 
This<br />
is <br />
a<br />
test<br /><br />

