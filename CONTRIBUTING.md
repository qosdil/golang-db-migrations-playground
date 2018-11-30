# Contributing to pressly-goose-test

For installation on servers, see README.md.

## Development with Docker and codegangsta/gin
	
### 1. Clone The Repository
### 2. Set up Configuration

Copy `/config.compose.json` file as `/config.json`.


### 3. Set up Compose Environment Variables

Copy `/.env.example` file as `/.env`, then change the values with the real ones.

### 4. Run Docker Compose

	$ docker-compose up
	
Done, now you can test the service from host machine:

	$ curl http://localhost:[DEV_HOST_PORT]
	$ curl http://localhost:[DEV_HOST_PORT]/movies

You can access the database from your host machine with the following credentials:

* Host: `localhost`
* Port: `[MYSQL_HOST_PORT]`
* User: `root`
* Password: `(null)`