# event-driven-architecture
Event-driven architecture

## Installating Rabbitmq via docker:
~~~
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management
~~~

To access panel : click [here](http://localhost:15672)

uname & pwd : guest

To install without docker, click [here](https://www.rabbitmq.com/download.html)

---

## Installing mongodb
~~~
docker run -it --rm --name mongo_db  -d -p 27018:27017 mongo
~~~

Note: Make sure to disable windows firewall while connecting mongodb in docker.

---