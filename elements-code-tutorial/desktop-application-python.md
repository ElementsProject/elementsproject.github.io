---
layout: page
title: Installing Bitcoin
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

That will have set up all we need to run our tutorial Python code.

Create a new file in your home directory, name it elementstutorial.py and paste the code below into it.

* * *

##### Note: Python requires that lines are indented correctly - make sure the code below is copied correctly with 4 spaces as indentations. Also note that some of the lines below wrap when viewed in a browser.

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

The code defines the details needed to connect to the elementsd node using RPC commands, sets up the method we want to execute as well as the parameter we want to pass in, executes the call and prints out the "balance" value from the results.

Before we try running the code make sure the required daemons are running:

~~~~
cd
cd elements
cd src
bitcoind -datadir=$HOME/bitcoindir
./elementsd -datadir=$HOME/elementsdir1
~~~~

If you get an error saying they are already running that's fine.

##### Note: If you get an error connecting to the elements client when you run the code below it may be because your node has been left in an altered state after quitting the tutorial code at an early stage. To refresh and reset the daemon’s blockchain and config files re-run the first section of the tutorial code up to and including the lines where the 3 config files are copied into the new directories then run the commands above to start the required daemons.

To run our Python code execute the following command:

~~~~
cd
python elementstutorial.py
~~~~

The result of which should be:

<img class="" alt="" src="{{ site.url }}/images/python.png" />

Obviously that's a very basic example but we now have the correct set up and you can use it as a building block for your future development work.

As an application would be making multiple calls to the elementsd daemon via RPC you will probably want to move the code that actually does the request and response work into its own function or use one of the existing Python interfaces to the Bitcoin JSON-RPC API and adapt it for your elements project. 

An example of this approach using AuthServiceProxy to execute rpc calls can be found on [Github](https://github.com/ElementsProject/elements/blob/elements-0.14.1/contrib/assets_tutorial/assets_tutorial.py).


[Next: Web application example]({{ site.url }}/elements-code-tutorial/web-application)

