## What is zombie mode?
you want to check if a host if open on specific port, but you do not want to directly send TCP handshake so you can use another host to do the handshake, who? let me explain:

as you may know _Every IP packet on the Internet has a fragment identification number (IP ID). Since many operating systems simply increment this number for each packet they send, probing for the IPID can tell an attacker how many packets have been sent since the last probe_ , so at first you send a packet to zomibe host, save it's IP ID, then send a ACK to the host you want to check, then the host send the ask-knowldge to the zombie host, then you send another ACK to zombie, if you see any increasement in IP ID then you can understand that you host has a open TCP socket!, for more information you can see [nmap TCP Idle Scan(-sI)](https://nmap.org/book/idlescan.html) 

example:  `zombie -> nmap -Pn -sI <ZOMBIE-IP> <TARGET> ( -p NUMBER or --packet-trace)`

**-Pn is necessary for stealth, otherwise ping packets would be sent to the target from Attacker's real address**

Full Open Scan (-sT) -> do all 3 hand shake steps
<br />
Stealth,syn half of Scan (-sS)
<br />
Xmas scan (-sX)
<br />
Fin scan (passfirewalls) (-sF)
<br />
Null scan (all flags are off) (-sN)
<br />
zombie -> nmap -Pn -sI <ZOMBIE-IP> <TARGET> ( -p NUMBER or --packet-trace)
zombie -> nmap -Pn -sI <ZOMBIE-IP> <TARGET> ( -p NUMBER or --packet-trace)
<br />

### Resources
[nmap official book](https://nmap.org/book/toc.html)
<br />
[man official page](https://linux.die.net/man/1/nmap)
<br />
[nmap mit](https://stuff.mit.edu/afs/athena/astaff/project/opssrc/nmap-3.00/docs/nmap_manpage.html)
<br />

