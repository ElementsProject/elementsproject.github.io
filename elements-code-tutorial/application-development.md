---
layout: page
title: Elements application development
permalink: /elements-code-tutorial/application-development
---

# Elements code tutorial

## Application development on Elements

In this part of the tutorial we will learn how to develop applications on top of Elements using RPC.

Examples are provided in [Python]({{ site.url }}/elements-code-tutorial/desktop-application-python), [C#]({{ site.url }}/elements-code-tutorial/desktop-application-dotnetcore), [Ruby]({{ site.url }}/elements-code-tutorial/other-languages#ruby), [Node.js]({{ site.url }}/elements-code-tutorial/other-languages#nodejs), [Go]({{ site.url }}/elements-code-tutorial/other-languages#go), [Perl]({{ site.url }}/elements-code-tutorial/other-languages#perl), and [Java]({{ site.url }}/elements-code-tutorial/other-languages#java). The Python and C# examples are expanded to cover web applications using [Flask]({{ site.url }}/elements-code-tutorial/web-application) and the [.NET MVC]({{ site.url }}/elements-code-tutorial/web-application-dotnetcore) frameworks. The examples can easily be amended to work against Bitcoin or a codebase based upon Elements, such as [Liquid](https://blockstream.com/liquid/), by amending the rpc settings in your config file.

In the main sections of the tutorial we used a terminal to send commands to the Elements client (elements-cli), which in turn would issue RPC (Remote Procedure Call) commands to the Elements daemon (elementsd). This configuration will be familiar to those who develop for Bitcoin, which the Elements code is based upon. 

The communication between daemon and client is possible because, when started in server mode, elementsd initiates an http server and listens for requests being made to it on a port specified in the associated elements.conf file. 

Requests to the daemon are made by posting [JSON](https://www.json.org/) formatted data to the http server port that the daemon is listening on. The request is processed and the results are returned as JSON formatted data.

Verification details of the credentials needed to make these calls is also stored in the config file. This is how elements-cli and elementsd were able to communicate in the previous tutorial code, they both shared the same config file and therefore the client could satisfy the authentication checks of the daemon. 

Take a look in the elements.conf file in $HOME/elementsdir1 that we created during the tutorial and notice the rpc related settings.

We can use the same authentication details and port number to send requests to the Elements daemon ourselves by using a programming language to make the RPC calls and process the returned results. 

We'll use Python for our first example, but any language that can make and receive http requests can be used. Basic examples are also provided for the following languages: [C#]({{ site.url }}/elements-code-tutorial/desktop-application-dotnetcore), [Ruby]({{ site.url }}/elements-code-tutorial/other-languages#ruby), [Node.js]({{ site.url }}/elements-code-tutorial/other-languages#nodejs), [Go]({{ site.url }}/elements-code-tutorial/other-languages#go), [Perl]({{ site.url }}/elements-code-tutorial/other-languages#perl), and [Java]({{ site.url }}/elements-code-tutorial/other-languages#java).


[Next: Desktop application example in Python]({{ site.url }}/elements-code-tutorial/desktop-application-python)

