---
layout: page
title: Elements Python app
permalink: /elements-code-tutorial/desktop-application-python
---

# Elements code tutorial

## Desktop application example in Python

In this example we will be using Python to make an RPC (Remote Procedure Call) to the Elements daemon (elementsd). This was what the Elements client (elements-cli) application was doing when we executed commands in the main section of the tutorial. Any language that can make and receive http requests could be used. 

Our aim is to simply make a call to elementsd using RPC by executing some basic Python code.

First we will need to install a few prerequisites. From the terminal run the following commands one after another:

~~~~
sudo apt-get install python-pip python-dev
sudo pip install --upgrade pip 
sudo pip install requests
~~~~

That will have set up all we need to run our Python tutorial code.

Create a new file named **elementstutorial.py** in your home directory and paste the code below into it.

* * *

~~~~
from __future__ import print_function
import requests, json

rpcPort = 18884
rpcUser = 'user1'
rpcPassword = 'password1'

serverURL = 'http://' + rpcUser + ':' + rpcPassword + '@localhost:' + str(rpcPort)

headers = {'content-type': 'application/json'}
payload = json.dumps({"method": 'getwalletinfo', "params": ["bitcoin"], "jsonrpc": "2.0"})

response = requests.post(serverURL, headers=headers, data=payload)

responseJSON = response.json()
responseResult = responseJSON['result']

print(responseResult['balance'])
~~~~

* * * 

The code defines the details needed to connect to the elementsd node using RPC commands, sets up the method we want to execute and the parameter we want to pass in, executes the call and prints out the "balance" value from the results.

Before we try running the code make sure the required daemons are running:

~~~~
cd
cd elements
cd src
bitcoind -datadir=$HOME/bitcoindir
./elementsd -datadir=$HOME/elementsdir1
~~~~

If you get an error saying they are already running that is fine.

##### Note: If you get an error connecting to the elements client when you run the code below it may be because your node has been left in an altered state after quitting the tutorial code at an early stage. To refresh and reset the daemon’s blockchain and config files re-run the first section of the tutorial code up to and including the lines where the 3 config files are copied into the new directories, then run the commands above to start the required daemons.

To run our Python code execute the following command:

~~~~
cd
python elementstutorial.py
~~~~

The result of which should be:

<img class="" alt="" src="{{ site.url }}/images/python.png" />

Obviously that's a very basic example but you now have a functioning setup which you can use as a building block for further development.

As an application would be making multiple calls to the elementsd daemon via RPC, you will probably want to move the code that actually does the request and response work into its own function, or use one of the existing Python interfaces to the Bitcoin JSON-RPC API and adapt it for your own project. 

An example of this approach that uses the **AuthServiceProxy** Python interface to Bitcoin's JSON-RPC API can be found in the elements/contrib/assets_tutorial/assets_tutorial.py file on [Github](https://github.com/ElementsProject/elements).


[Next: Web application example]({{ site.url }}/elements-code-tutorial/web-application)

