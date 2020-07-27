# What is SOAP?

[SOAP (abbreviation for Simple Object Access Protocol) is a messaging protocol specification for exchanging structured information in the implementation of web services in computer networks. Its purpose is to provide extensibility, neutrality, verbosity and independence.It uses XML Information Set for its message format, and relies on application layer protocols, most often Hypertext Transfer Protocol (HTTP), although some legacy systems communicate over Simple Mail Transfer Protocol (SMTP), for message negotiation and transmission](https://en.wikipedia.org/wiki/SOAP)



## [SOAP example over HTTP] (https://www.w3schools.com/xml/xml_soap.asp)

```
<?xml version="1.0"?>

<soap:Envelope
xmlns:soap="http://www.w3.org/2003/05/soap-envelope/"
soap:encodingStyle="http://www.w3.org/2003/05/soap-encoding">

<soap:Body>
  <m:GetPrice xmlns:m="https://www.w3schools.com/prices">
    <m:Item>Apples</m:Item>
  </m:GetPrice>
</soap:Body>

</soap:Envelope>
```


### Refrences

#### Official w3 attrs

[soap env](https://www.w3.org/2003/05/soap-envelope/)
[soap encing](https://www.w3.org/2003/05/soap-encoding)

[what is soap?](https://www.guru99.com/soap-simple-object-access-protocol.html)
