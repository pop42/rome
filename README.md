# Rome

A sample system using micro and go-micro with NATS to show a series of microservices working in tandem.

This example involves a heavy use of docker bringing together 10 seperate instances. 

The sources are not using real or even real-ish data as the purpose of this demo is to demonstrate multiple microservices working together, not a valid election data display system.

** Why is this named **Rome**?  They sort of connected a lot of things together with all their roads so it made sense at first... but I wasn't very good at keeping up with the analogy so you can see my naming convention quickly dropped off from there.   
_________
#### How this works

The three election sources are simulating election results coming in by updating the voteCount a random amount for a random candidate at random-ish intervals (every 2 seconds, but rolls to see if it really should update or just go back to sleep).  Everytime a source updates the voteCount for that candidate, it publishes it to the go.micro.srv.Notekeeper.Race channel in NATS.  

Notekeeper is listening on that channel for any updates.  When it gets an update it will pass it on to the go.db.Race.post channel (without doing an update becuase the data was already normalized, but the point is it COULD have normalized the data here).

Caesar is listening on the go.db.Race.post channel.  When it sees an update, it immediately updates the database with the appropriate data.


#### Containers
* elections_ap - Associated Press Source for election data.
* elections_mc - Maricopa County Source for election data.
* elections_sos - AZ Secretary of State Source for election data.
* notekeeper - Intermediate process (could represent normalization process).
* caesar - updates DB with publications from Notekeeper.
* nero - simple web service to display the results.
* sidecar - language agnostic RPC proxy that could be used to add other microservices to the mix.  I'm mostly using it here for service discovery and registering services.
* web - Web dashboard and reverse proxy for micro web applications.
* postgresdb - holds the database.


##### sidecar
* **localhost:8081/registry** shows JSON representation of services
* **localhost:8081/registery?service=<service.name>** shows detailed information on a specific service

##### web
* **localhost:8082** shows a web-based representation of what sidecar

##### nero
* **localhost:8083/race** shows race results
* **localhost:8083/race/wipeout** wipes out all race data



#### Installation

Prerequisites:
* docker


1. First you clone the repo.  
2. Then you CD into it and type `docker-compose build` to build the containers.  
3. Then type `docker-compose up` to start it.





