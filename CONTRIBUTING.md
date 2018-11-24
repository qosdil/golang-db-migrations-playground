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
	
### 5. Migrate Schemas

If it is the first time, we need to create `pressly_goose_test` database inside the `mysql` container:

	$ docker exec -it pressly-goose-test-dev /bin/bash
	
From inside the container:
	
	$ mysql -h mysql -e "create database pressly_goose_test"
	
Stay inside the container, then continue with step 2 of **Migrate Schemas** section in README.md file.

After that, you can test the service from host machine:

	$ curl http://localhost:[DEV_HOST_PORT]
	$ curl http://localhost:[DEV_HOST_PORT]/movies
