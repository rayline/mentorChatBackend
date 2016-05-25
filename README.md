# mentorChatBackend
backEnd for mentorChat, a project for homework in BUAA mentor project

## INSTALLATION
The program relys on Redis and upsetly, Redis has no windows binaries so you can only install it on OS like linux or Mac  
Since it is just a pre-release, you may have to install it manually, including fixing dependencies yourself, here is what we need
  
+ Redis Server (go download at <code>redis.io</code>)
+ Go tools (please go to <code>golang.org/dl</code> to download and install one since the one on apt-get repo is ridiculously old)
+ beego (<code>go get "github.com/astaxie/beego"</code>)
+ redigo (<code>go get "github.com/garyburd/redigo/redis"</code>)

And then with <code>go build</code> you can "easily" build and run the server
With <code>app.conf</code> you will be able to adjust some settings
