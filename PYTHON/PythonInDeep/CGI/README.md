# What is CGI?

Common Gateway Interface (CGI) is an interface specification for web servers to execute programs like console applications (also called command-line interface programs) running on a server that generates web pages dynamically. Such programs are known as CGI scripts or simply as CGIs. The specifics of how the script is executed by the server are determined by the server. In the common case, a CGI script executes at the time a request is made and generates HTML
<small>[1](https://en.wikipedia.org/wiki/Common_Gateway_Interface)</small>


## Web browsing and CGI

To understand the concept of CGI, let us see what happens when we click a hyper link to browse a particular web page or URL.

*   Your browser contacts the HTTP web server and demands for the URL, i.e., filename.
*   Web Server parses the URL and looks for the filename. If it finds that file then sends it back to the browser, otherwise sends an error message indicating that you requested a wrong file.
*   Web browser takes response from web server and displays either the received file or error message.

However, it is possible to set up the HTTP server so that whenever a file in a certain directory is requested that file is not sent back; instead it is executed as a program, and whatever that program outputs is sent back for your browser to display. This function is called the Common Gateway Interface or CGI and the programs are called CGI scripts. These CGI programs can be a Python Script, PERL Script, Shell Script, C or C++ program, etc.
<small>[2](https://www.tutorialspoint.com/python/python_cgi_programming.htm)</small>

## The problem with CGI

CGI scripts execute time-consuming code (like opening a database) every time user send a request


***The solution***

## FastCGI

FastCGI is a way to have CGI scripts execute time-consuming code (like opening a database) only once, rather than every time the script is loaded. In technical terms, FastCGI is a language independent, scalable, open extension to CGI that provides high performance without the limitations of server specific APIs.
<small>[3](https://help.dreamhost.com/hc/en-us/articles/217298967-FastCGI-overview)</small>

