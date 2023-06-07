*	**TELNET(tcp/23)**:stands for telecommunication network

Login to devices remotely

No security


*	**SSH(tcp/22)**:

incrp hight security

*	**DNS(tcp/53)**:(~Convert name to IP addr)


*	**SMTP(tcp/25)**: Simple Mail Trasfer Protocol

Server to Server com

Client to sever com

*	**SFTP(tcp/22)**: Secure File Transport Protocol(use SSH underlayer)

*	**FTP(tcp/20)<active mode data>**:File Trasfer Protocol

*	**FTP(tcp/21)<Control>**
Unlike SFTP that use SHH and it's secure FTP is not secure


* 	**TFTP(udp/69)**:Trivial File Trasfer Protocol

Unlike FTP&UFTP it's very simple

No auth

Just read and write

*	**DHCP(udp/67)(udp/68)**:Dynamic Host Configuration Protocol

- Automated Configuration of IP addr, subnet mask and other options

- Required a DHCP server


-Ip addrs are assigned in readl-time a pool
Each system is given a least and must renew at set intervals (or there is a expiration date)

- Addrss are assigned by MAC addr in the DHCP server


*	**HTTP(tcp/80)**:Hypter Text Trasfer Protocol

*	**HTTPS(tcp/443)**:Hyper Text Traster Protocol Sevure

*	**SNMP(udp/161)**:Simple Network Management Protocol

*	**RDP(tcp/3389)**:Remote Desktop Protocol

*	**NTP(udp/123)**:Network Time Protocol

Sync clock and time 

*	**SIP(tcp/5060)(tcp/5061)**:Session Intiation Protocol

VIOP signaling and voice communication

*	**POP(tcp/110)**:Post Office Protocol

we use SMTP for **SEND** email and use POP3(stand for V3)  or IMP4(stand for V4) for receiv emails

Basic management

*	**IMAP4(tcp/143)**:Internet Message Access Protocol

Management email box from multiple clients

*	**LDAP(tcp/389)**:Lightwiehgt Directory Access Protocol

Store and retrive info in a network directory(like printer or ...)

*	**LDAPS(tcp/636)**: " " " " Secure

Use SSL

*	**H.323**: for voip
