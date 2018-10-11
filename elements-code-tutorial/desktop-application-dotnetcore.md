---
layout: page
title: Installing Bitcoin
permalink: /elements-code-tutorial/desktop-application-dotnetcore
---

# Elements code tutorial

## Desktop application example in C# using .NET Core

Microsoft have created an open source version of their .NET platform called .NET Core that runs on a variety of operating systems. We will use C# (the most popular .NET language) for this example.

Before installing .NET Core, you'll need to register the Microsoft key, register the product repository, and install required dependencies.

Open a command prompt and run the following two commands.

##### NOTE: There are only two lines below so note the text wrap. The first starts with "wget", the second with "sudo".

~~~~
wget -q packages-microsoft-prod.deb https://packages.microsoft.com/config/ubuntu/16.04/packages-microsoft-prod.deb
sudo dpkg -i packages-microsoft-prod.deb
~~~~

Now we can install the .NET Core SDK:

~~~~
sudo apt-get install apt-transport-https
sudo apt-get update
sudo apt-get install dotnet-sdk-2.1.4
~~~~

Check it is set up correctly and create a simple console application:

~~~~
dotnet new console -o dotnetelements
cd dotnetelements
dotnet run
~~~~

Which will output:

<div class="console-output">Hello World!
</div>

Before we try running any code we'll make sure the required daemons are running:

~~~~
cd
cd elements
cd src
bitcoind -datadir=$HOME/bitcoindir
./elementsd -datadir=$HOME/elementsdir1
~~~~

If you get an error saying they are already running that's fine.

##### NOTE: If you get an error connecting to the elements client when you run the code below it may be because your node has been left in an altered state after quitting the tutorial code at an early stage. To refresh and reset the daemon's blockchain and config files re-run the first section of the tutorial code up to and including the lines where the 3 config files are copied into the new directories then run the commands above to start the required daemons.

Edit the "Program.cs" file in $HOME/dotnetelements using a text editor and change the contents to the following and save the file:

* * * 

##### NOTE: Some of the lines below wrap when viewed in a browser.
~~~~
using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace dotnetelements
{
    class Program
    {
        private static HttpClient client = new HttpClient();
        
        static void Main(string[] args)
        {
            MainAsync().GetAwaiter().GetResult();
        }

        private static async Task MainAsync()
        {
            string url = "http://user1:password1@localhost:18884";

            var contentData = new StringContent(@"{""method"": ""getwalletinfo"", ""params"": [""bitcoin""], ""jsonrpc"": ""2.0""}", System.Text.Encoding.UTF8, "application/json");
            using (HttpResponseMessage response = await client.PostAsync(url, contentData))
            using (HttpContent content = response.Content)
            {
                string data = await content.ReadAsStringAsync();
                
                if (null != data)
                {
                    Console.WriteLine(data);
                }
            }
        }
    }
}
~~~~

* * * 

Then move to the the directory with the file in, compile and run the application again:

~~~~
cd
cd dotnetelements
dotnet run
~~~~

Which outputs:

<img class="" alt="" src="{{ site.url }}/images/dotnet.png" />

As an application would be making multiple calls to the elementsd daemon via RPC you will probably want to move the code that actually does the request and response work into its own function. Again, the code above is a starting point to get you up and running.


[Next: An easy way to run the main tutorial code]({{ site.url }}/elements-code-tutorial/easy-run-code)

