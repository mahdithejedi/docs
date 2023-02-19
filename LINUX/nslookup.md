[geeksforgeeks](https://www.geeksforgeeks.org/nslookup-command-in-linux-with-examples/#discuss)
<br />
[What is nameserever](https://www.geeksforgeeks.org/nslookup-command-in-linux-with-examples/#discuss)
<br />

### Example

Ask for _a_ record of twitter from one of twitter namespaces

[//]: # (__NAMESPACE__ , __NSLOOKUP__)

```shell
nslookup -type=a twitter.com ns1.p34.dynect.net
```
<br />

`ns1.p34.dynect.net` come from

```shell
nslookup -type=ns twitter.com
```
which asks for _namespace_ of twitter

