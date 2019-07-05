---
layout: page
title: Elements - Other Languages
permalink: /elements-code-tutorial/other-languages
---

# Elements code tutorial

## Other language examples

Our aim for each language example is to make simple calls to elementsd using RPC. The examples are very basic but provide a way to get to get a functioning setup which you can use as a building block for further development.

### Ruby

You can check if Ruby is installed on your operating system, and install it if not, by following the steps [here](https://www.ruby-lang.org/en/documentation/installation/).

Create a new file named 'elementsrpcruby.rb' and paste the code below into it:

~~~~
require 'net/http'
require 'uri'
require 'json'

class ElementsRPC
  def initialize(service_url)
    @uri = URI.parse(service_url)
  end

  def method_missing(name, *args)
    post_body = { 'method' => name, 'params' => args, 'id' => 'jsonrpc' }.to_json
    resp = JSON.parse( http_post_request(post_body) )
    raise JSONRPCError, resp['error'] if resp['error']
    resp['result']
  end

  def http_post_request(post_body)
    http    = Net::HTTP.new(@uri.host, @uri.port)
    request = Net::HTTP::Post.new(@uri.request_uri)
    request.basic_auth @uri.user, @uri.password
    request.content_type = 'application/json'
    request.body = post_body
    http.request(request).body
  end

  class JSONRPCError < RuntimeError; end
end

if $0 == __FILE__
  elements = ElementsRPC.new('http://user1:password1@127.0.0.1:18884')
 
  p elements.getblockcount
end
~~~~

Execute the code from the command line:

~~~~
ruby elementsrpcruby.rb
~~~~

The output will show the current block count.

* * * 

### Java

The example code doesn't require any external dependencies, but you do need the Java compiler (javac) as well as the Java runtime environment itself. These can be installed on Ubuntu using the command:

~~~~
sudo apt install default-jdk
~~~~

Create a new file named 'elementsrpcjava.java' and paste the code below into it:

~~~~
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import java.net.Authenticator;
import java.net.PasswordAuthentication;

public class elementsrpcjava {
	
	public static void main (String []args) throws IOException{
		URL url = new URL ("http://127.0.0.1:18884");
		String rpcuser ="user1";
		String rpcpassword ="password1";
		
		Authenticator.setDefault(new Authenticator() {
            protected PasswordAuthentication getPasswordAuthentication() {
                return new PasswordAuthentication (rpcuser, rpcpassword.toCharArray());
            }
        });
  
		HttpURLConnection conn = (HttpURLConnection)url.openConnection();
		conn.setRequestMethod("POST");
		conn.setRequestProperty("Content-Type", "application/json; utf-8");
		conn.setRequestProperty("Accept", "application/json");
		conn.setDoOutput(true);
				
		String json = "{\"method\": \"getwalletinfo\", \"jsonrpc\": \"2.0\"}";
		
		try(OutputStream os = conn.getOutputStream()){
			byte[] input = json.getBytes("utf-8");
			os.write(input, 0, input.length);			
		}
		
		try(BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream(), "utf-8"))){
			StringBuilder sb = new StringBuilder();
			String line = null;
			
			while ((line = br.readLine()) != null) {
				sb.append(line.trim());
			}
			
			System.out.println(sb.toString());
		}
	}
}
~~~~

Compile and run the code from the command line:

~~~~
javac elementsrpcjava.java
java elementsrpcjava
~~~~

The output will show wallet information.

* * * 

### Go

To install Go: [https://golang.org](https://golang.org). Note the importance of setting the PATH environment variable.

Create a directory src/elements inside your Go workspace directory (probably %HOME/go) so that the full path is: $HOME/go/src/elements

With that directory create a file named 'elementsrpcgo.go' and past the following into it:

~~~~
package main

import (
    "log"
    "bytes"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func main() {
    url := "http://user1:password1@localhost:18884"
    
    var jsonStr = []byte(`{"jsonrpc":"1.0","method":"getblockcount","params":[]}`)

    request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    
    request.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    
    response, err := client.Do(request)
    
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    
    body, _ := ioutil.ReadAll(response.Body)
    
    log.Println(string(body))
    
    //Use as JSON...
    var dat map[string]interface{}
    
    if err := json.Unmarshal(body, &dat); err != nil {
        panic(err)
    }
    
    blocks := dat["result"].(float64)
    log.Println(blocks)
}
~~~~

Compile and run the code:

~~~~
go build
./elements
~~~~

The code prints out the current block count.

* * * 

### Perl

Perl can be installed from [https://www.perl.org](https://www.perl.org) and comes pre-installed on Ubuntu.

We will be using the [JSON::RPC::Client](https://metacpan.org/pod/release/MAKAMAKA/JSON-RPC-0.95/lib/JSON/RPC/Client.pm) Perl implementation of a JSON-RPC client which will make the rpc calls and results handling simpler.

To install JSON::RPC::Client open the CPAN tool:

~~~~
perl -MCPAN -e shell
~~~~

And from within the cpan console, install the RPC Client:

~~~~
install JSON::RPC::Client
~~~~

And then exit the cpan tool:

~~~~
quit
~~~~

You must close the terminal window so that when we run the code below it will pick up the new environment variables written.

Create a new file named 'elementsrpcperl.pl' and paste the code below into it:

~~~~
use JSON::RPC::Client;
use Data::Dumper;

my $client = new JSON::RPC::Client;

$client->ua->credentials(
    'localhost:18884', 'jsonrpc', 'user1' => 'password1'
    );

my $uri = 'http://localhost:18884/';
my $obj = {
    method  => 'getwalletinfo',
    params  => [],
};

my $res = $client->call( $uri, $obj );

if ($res){
    if ($res->is_error) { print "Error : ", $res->error_message; }
    else { print Dumper($res->result); }
} else {
    print $client->status_line;
}
~~~~

Run the code from the command line (remember that this must be a new terminal window from the one we used before to install the RPC Client):

~~~~
perl elementsrpcperl.pl
~~~~

The output will show wallet information.


[Next: An easy way to run the main tutorial code]({{ site.url }}/elements-code-tutorial/easy-run-code)
