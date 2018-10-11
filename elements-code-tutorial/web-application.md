---
layout: page
title: Installing Bitcoin
permalink: /elements-code-tutorial/web-application
---

# Elements code tutorial

## Web application example using Python with Django

Now we will shift the Python code from the previous tutorial section and have it run from Django (a Python web framework) so we can write the output to a web page. 

Assuming you have already installed the prerequisites in the [Desktop application example in Python]({{ site.url }}/elements-code-tutorial/desktop-application-python) section we just need to install Django:

~~~~
sudo pip install Django
~~~~

Before we try running any code make sure the required daemons are running:

~~~~
cd
cd elements
cd src
bitcoind -datadir=$HOME/bitcoindir
./elementsd -datadir=$HOME/elementsdir1
~~~~

If you get an error saying they are already running that's fine.

##### NOTE: If you get an error connecting to the elements client when you run the code below it may be because your node has been left in an altered state after quitting the tutorial code at an early stage. To refresh and reset the daemon’s blockchain and config files re-run the first section of the tutorial code up to and including the lines where the 3 config files are copied into the new directories then run the commands above to start the required daemons.

Now create a new directory in home named elements-django and within that a file named elementstutorial.py into which you should paste the following code, most of which you will be familiar with from the last exercise. 

* * * 

##### NOTE Python requires that lines are indented correctly - make sure the code below is copied correctly with 4 spaces as indentations. Also note that some of the lines below wrap when viewed in a browser.

~~~~
#!/usr/bin/env python
from __future__ import print_function

import sys
import requests, json

from django.conf import settings 
from django.http import HttpResponse
from django.core.management import execute_from_command_line
from django.conf.urls import url, include

settings.configure(
    DEBUG=True,
    SECRET_KEY='asecretkey',
    ROOT_URLCONF=sys.modules[__name__],
)
 
def index(request):
    rpcPort = 18884
    rpcUser = 'user1'
    rpcPassword = 'password1'

    serverURL = 'http://' + rpcUser + ':' + rpcPassword + '@localhost:' + str(rpcPort)

    headers = {'content-type': 'application/json'}
    payload = json.dumps({"method": 'getwalletinfo', "params": ["bitcoin"], "jsonrpc": "2.0"})

    response = requests.post(serverURL, headers=headers, data=payload)

    responseJSON = response.json()
    responseResult = responseJSON['result']

    return HttpResponse(responseResult['balance'])

urlpatterns = [
    url(r'^elementstutorial/', index)
]
 
if __name__ == "__main__":
    execute_from_command_line(sys.argv)
~~~~

* * * 

Navigate to the right directory and start the web server and our code:

~~~~
cd
cd elements-django
python elementstutorial.py runserver
~~~~

You should see the web server startup. Now all you have to do is open a browser and go to:

<div class="console-output">http://127.0.0.1:8000/elementstutorial
</div>

Which simply writes the result of the "balance" call to the page:

<img class="" alt="" src="{{ site.url }}/images/django.png" />

As in the previous exercise, this is intended to get you up and running. The rest is up to you!


[Next: Desktop application example in C# using .NET Core]({{ site.url }}/elements-code-tutorial/desktop-application-dotnetcore)

