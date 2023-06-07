# What is arp?

As we know switches are layer 2 and do not understand IP addresses but how can we ping an IP address in switch layer 2?
<br />
the answer is in the **arp** hand..

you think that your ip address is _192.168.1.30_ and you want to ping _192.168.1.20_ what will happen?
<br />

the arp broadcast a message to all devices within switch and says _who has 192.168.1.20_ ip address? 
<br />
after the _192.168.1.20_ device say what was me and send the message to you and give you it's mac address now you can commiunicate witheach other with out knowing ip address..
<br />
**Importent note**: The ip address in this way is set manually and do not set with swich layer 2 <small>with `ip` command</small>
<br />

## More technichally

Address Resolution Protocol (ARP) is a procedure for mapping a dynamic Internet Protocol address (IP address) to a permanent physical machine address in a local area network (LAN).[techtarget](https://searchnetworking.techtarget.com/definition/Address-Resolution-Protocol-ARP)

### More resources

[wikipedia](https://en.wikipedia.org/wiki/Address_Resolution_Protocol#Operating_scope)
