---
layout: page
title: Application development
permalink: /elements-code-tutorial/application-development
---

# Elements code tutorial

## Application development on Elements

In the main sections of the tutorial we used a terminal to send commands to the Elements client (elements-cli), which in turn issued an RPC (Remote Procedure Call) to the Elements daemon (elementsd). This will be very familiar to those who develop for Bitcoin, which the Elements code is based upon.

The communication between daemon and client is possible because, when started in server mode, elementsd initiates an http server and listens for requests being made to it on a port specified in the associated elements.conf file. 

Requests to the daemon are made by posting JSON formatted data to the http server port the daemon is listening on. The request is processed and the results are returned as JSON formatted data.

Verification details of the credentials needed to make these calls is also stored in the config file. This is how elements-cli and elementsd were able to communicate in the previous tutorial code - they both shared the same config file and therefore the client could satisfy the authentication checks of the daemon. 

Take a look in the elements.conf file in $HOME/elementsdir1 and notice the following:

<div class="console-output">rpcuser=user1
rpcpassword=password1
rpcport=18884
daemon=1
</div>

We can use the same authentication details and port number to send requests to the elements daemon ourselves using a programming language and making RPC calls. We'll use Python for our first example, but any language that can make and receive http requests could be used. 


[Next: Desktop application example in Python]({{ site.url }}/elements-code-tutorial/desktop-application-python)

