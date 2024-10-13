# TweedeKamerVolger

This project analyzes the data publisched by the dutch government `Tweede Kamer der staten genera` and makes it insightfull through a user interface.

## Gegevends Magazijn Explained

The `Tweede Kamer der Staten-Generaal` publishes all data in the [Gegevens Magazijn](https://gegevensmagazijn.tweedekamer.nl).

The documentation on the structure and content of the data can be found [here](https://opendata.tweedekamer.nl/). In this project we will use the SyncFeed API to extract all data, index it and structure it so it can be easily searched and analyzed.

The APIs are:

* [Feed](https://gegevensmagazijn.tweedekamer.nl/SyncFeed/2.0/Feed)
* [EntityID](https://gegevensmagazijn.tweedekamer.nl/SyncFeed/2.0/Entiteiten/<id>)
* [ResourceID](https://gegevensmagazijn.tweedekamer.nl/SyncFeed/2.0/Resources/<id>)

## TweedeKamerVolger Architecture

This project uses event driven architecture and has several microservices:

* FeedAnalyzed (Analyzess the data feed for changes and creates events)
* RabbitMQ (message Queue for our events)
* Redis of MongoDB
* ElasticSearch (Index the events and attached documents)
* Postgressql (stores events structured and correlated)
* FeedFront (Frontend to visualize the data)
* FeedProcessor (Processes new events, indexes and store the event)
* FeedDocument (Processes Additional documents identified in the new events)
