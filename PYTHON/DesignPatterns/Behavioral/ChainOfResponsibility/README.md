# What is chain of responsibitilty

Imagine a POST request, you want to check some parts before analizy data

like chech if HEADERS are set, or use is autherized or not, or if the request is post request

Insread of wrinte a big dirty conde you can seprate every operation to a handler, so every handler will do specific (unique) task like authrization ot checking HEADRS


