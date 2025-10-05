This was a hobby project to get minecraft running on a distroless installation. 
I had never worked with go before so I did a side project learning that through the course of this project.

Documentation and how to use it if you actually want to:

This docker compose file enables you to install java and it also makes port configurations in server.properties less annoying because you just need the yaml file

It is also distroless (for fun, security and size)

Currently the env variables currently supported are 
Xmx="4G"
jarName="server.jar"

the values can be anything (I should really change the naming scheme for the jar variable wtf is that, but oh shit if somene actually uses this and i push a commit with something simpler like jar that will break for that person. Oh well I guess this is how we get tech debt)

regardless if you want to add anything it's pretty easy to add some code to the main.go file just use my code as a template.
glhf
