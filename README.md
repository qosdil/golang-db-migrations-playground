# pressly-goose-test 
v1.3.0

Playing around with pressly/goose.

## Prerequisites
* Go v1.11
* Dep v0.5 (for non-Docker)
* MySQL v5.5

## Build

Make sure that your system has met the prerequisites above before running the following build steps.

### Docker

#### 1. Clone The Repository
#### 2. Set up Configuration

Copy `/config.example.json` file as `/config.json`, then change the values with the real ones.
#### 3. Build Image

Move to the cloned repository:

	$ cd [app_root]

Then run Docker build command. In this example, version is `1.2.0`:

	$ docker build -t pressly-goose-test-image:1.2.0 .

### Non-Docker

#### 1. Clone The Repository

Clone the repo under your `$GOPATH/src` directory.

#### 2. Set up Configuration

Copy `/config.example.json` file as `/config.json`, then change the values with the real ones.

#### 3. Install Dependencies

Move to the cloned repository, then run the following:

	$ dep ensure
	
#### 4. Build

	$ go build
	
## Migrate Schemas
	
### 1. Create a Database
*If it is not the first installation / release, skip this step.*

Create a MySQL database, for example: `pressly_goose_test`.

	CREATE DATABASE `pressly_goose_test` DEFAULT CHARACTER SET = `utf8`;

### 2. Install pressly/goose v2.3.0

	$ bash ./install-pressly-goose-2.3.0.sh
	
### 3. Move to Migrations Directory

	$ cd [app_root]/database/migrations
	
### 4. Run Migrations

To create/update the database schemas, run the `up` command like the following example:

	# Database name: pressly_goose_test
	# User: root, password: (not set)
	# Host: localhost, port: 3306
	
	$ goose mysql "root@tcp(localhost:3306)/pressly_goose_test?parseTime=true" up
	
To roll it back, run the `down` command one or more times until it reaches the previous state:

	$ goose mysql "root@/tcp(localhost:3306)pressly_goose_test?parseTime=true" down

	
## Run

After running the steps above, you can run the service:

### Docker

The following run command example binds port 3000 of the container to port 8080 on the host machine.

	$ docker run -d -p 8080:3000 -e PORT=3000 --name pressly-goose-test-container-1.2.0 pressly-goose-test-image:1.2.0

### Non-Docker

	$ cd [app_root]
	$ PORT=8080 ./pressly-goose-test

***
&copy; 2018 Kudo.