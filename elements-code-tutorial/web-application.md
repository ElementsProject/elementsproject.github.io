---
layout: page
title: Elements Python Web app
permalink: /elements-code-tutorial/web-application
---

# Elements code tutorial

## Python web application examples for Django and Flask

Django and Flask are two of the most popular web frameworks for Python. We will show examples of how to use either framework to display the results of RPC calls to our Elements node on a web page, forming the building block for web application development.

<a id="django"></a>
### Using Django

We will use the Python code from the previous tutorial section and have it run from Django so that we can write the output to a web page. 

Assuming you have already installed the prerequisites in the [Desktop application example in Python]({{ site.url }}/elements-code-tutorial/desktop-application-python) section, we just need to install Django:

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

If you get an error saying they are already running, that is fine.

##### Note: If you get an error connecting to the elements client when you run the code below it may be because your node has been left in an altered state after quitting the tutorial code at an early stage. To refresh and reset the daemon’s blockchain and config files re-run the first section of the tutorial code up to and including the lines where the 3 config files are copied into the new directories, then run the commands above to start the required daemons.

Now create a new directory in home named **elements-django** and within that create a file named **elementstutorial.py**, into which you should paste the following code, most of which you will be familiar with from the last section. Again, we will be using AuthServiceProxy to handle the connection, authentication and data typing for us as we communicate with our node.

##### Note: Python requires that lines are indented correctly. Make sure the code below is copied correctly and is using 4 spaces for indenting lines. Also note that some of the lines below wrap when viewed in a browser.

~~~~
#!/usr/bin/env python
from bitcoinrpc.authproxy import AuthServiceProxy, JSONRPCException

from django.conf import settings 
from django.http import HttpResponse
from django.core.management import execute_from_command_line
from django.conf.urls import url, include

import sys

settings.configure(
    DEBUG=True,
    SECRET_KEY='asecretkey',
    ROOT_URLCONF=sys.modules[__name__],
)
 
def index(request):
    rpc_port = 18884
    rpc_user = 'user1'
    rpc_password = 'password1'

    try:
        rpc_connection = AuthServiceProxy("http://%s:%s@127.0.0.1:%s"%(rpc_user, rpc_password, rpc_port))
    
        result = rpc_connection.getwalletinfo("bitcoin")
    
    except JSONRPCException as json_exception:
        return HttpResponse("A JSON RPC Exception occured: " + str(json_exception))
    except Exception as general_exception:
        return HttpResponse("An Exception occured: " + str(general_exception))

    return HttpResponse(result['balance'])    

urlpatterns = [
    url(r'^elementstutorial/', index)
]
 
if __name__ == "__main__":
    execute_from_command_line(sys.argv) 
~~~~

Navigate to the right directory, start the web server and execute our code:

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

As in the previous exercise, this is intended to get you up and running. You now have a functioning setup which you can use as a building block for further development using Django.

* * * 

<a id="flask"></a>
### Using Flask

Flask is one of the most popular web frameworks for Python. This example shows how to use the framework to display the results of RPC calls to our Elements node on a web page, forming the building block for web application development.

We will use the Python code from the previous tutorial section and have it run from Flask so that we can write the output to a web page. 

Assuming you have already installed the prerequisites in the [Desktop application example in Python]({{ site.url }}/elements-code-tutorial/desktop-application-python) section, we just need to install Flask:

~~~~
sudo pip install Flask
~~~~

Before we try running any code make sure the required daemons are running:

~~~~
cd
cd elements
cd src
bitcoind -datadir=$HOME/bitcoindir
./elementsd -datadir=$HOME/elementsdir1
~~~~

If you get an error saying they are already running, that is fine.

Now create a new directory in home named **elements-flask** and within that create a file named **elementstutorial.py**, into which you should paste the following code, most of which you will be familiar with from the last section. Again, we will be using AuthServiceProxy to handle the connection, authentication and data typing for us as we communicate with our node.

##### Note: Python requires that lines are indented correctly. Make sure the code below is copied correctly and is using 4 spaces for indenting lines. Also note that some of the lines below wrap when viewed in a browser.

~~~~
from flask import Flask
from bitcoinrpc.authproxy import AuthServiceProxy, JSONRPCException

app = Flask(__name__)
 
@app.route("/")
def elements():
    rpc_port = 18884
    rpc_user = 'user1'
    rpc_password = 'password1'

    try:
        rpc_connection = AuthServiceProxy("http://%s:%s@127.0.0.1:%s"%(rpc_user, rpc_password, rpc_port))
    
        result = rpc_connection.getwalletinfo("bitcoin")
    
    except JSONRPCException as json_exception:
        return "A JSON RPC Exception occured: " + str(json_exception)
    except Exception as general_exception:
        return "An Exception occured: " + str(general_exception)

    return str(result['balance'])
 
if __name__ == "__main__":
    app.run()
~~~~

Navigate to the right directory, start the web server and execute our code:

~~~~
cd
cd elements-flask
python elementstutorial.py
~~~~

You should see the web server startup. Now all you have to do is open a browser and go to:

<div class="console-output">http://127.0.0.1:5000
</div>

Which simply writes the result of the "balance" call to the page:

<img class="" alt="" src="{{ site.url }}/images/flask.png" />

As in the previous exercise, this is intended to get you up and running. You now have a functioning setup which you can use as a building block for further development using Flask.


[Next: Desktop application example in C# using .NET Core]({{ site.url }}/elements-code-tutorial/desktop-application-dotnetcore)

