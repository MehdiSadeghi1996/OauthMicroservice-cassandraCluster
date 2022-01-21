
# Oauth Microservice Go,Cassandra

implement microservice of oauth using golang and cassandra to store user tokens







![Logo](https://cilium.io/static/04756b327d0e0b7bd9c44efa00e2839a/f275c/cassandra_go2.webp)


## Optimizations

using one session of cassandra
in cassandra doc :it is safe to use one session of it or cluster of it.



## Features

- using respository pattern and service pattern
- using one interface for throw errors
- clean structure of go code


## Run Locally

ساخت کانتینر کاساندرا در هاست داکر با داکر کاکپوز

```bash
version: '3.7'

services:
  cassandra1:
    container_name: cassandra
    image: cassandra:latest
    ports:
       - 9042:9042
    volumes:
      - cassandra_persist:/var/lib/cassandra
    restart: always

volumes:
  cassandra_persist:

```

ساخت کی اسپیس و فیلد های آن با توجه به دامین مدل پروژه

```bash
  CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':1}
  USE oauth;
  CREATE TABLE access_tokens(access_token varchar PRIMARY KEY,user_id bigint,client_id bigint,expires bigint);

```
