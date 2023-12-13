# Concepts
for transaction, we need to remember these keywords:
 * `WATCH` : for watching a key for any changes during transaction
 * `MULTI` : start a transaction
 * `EXEC`: execute command
 * `DISCARD`: for discarding transactions

**Point:** all the command during transaction will be queues and when _`EXEC`_ 
is called all the commands will be run sequentially and if one of the commands failed
whole transaction will **NOT** discard or rollback only that specific command 
will not be executed because in redis there is no _rollback_

# Example
Consider two clients in redis both want to change `name`:

client1:
```redis
# this command will watch for any changes for `name` and if so
# then will not let EXEC run successfully and return _nil_ instead of _OK_
WATCH Name
# start transaction
MULTI
# this will cause this command to Queue for execuing when `EXEC` is called
SET Name "Mahdi"
``` 
while client2 is changing Name to "Reza"

client2:
```redis
SET Name "Reza"
```

if client1 call execute it will get error because client2 change the
_Name_ key and client1 set `WATCH` command for _Name_

```redis
EXEC
# return nil
```