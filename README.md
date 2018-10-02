# pressly-goose-test 
v1.1.0

Playing around with pressly/goose.

## Prerequisites
* Go v1.11
* MySQL v5.5

## Build

Make sure that your system has met the prerequisites above before running the following build steps.

### 1. Install Dependency Packages

	$ go get github.com/gorilla/mux
	$ go get github.com/go-sql-driver/mysql

### 2. Clone The Repository

Clone the repo under your `$GOPATH/src` directory.


### 3. Build

	$ go build
	
## Migrate Schemas
	
### 1. Create a Database
*If it is not the first installation / release, skip this step.*

Create a MySQL database, for example: `pressly_goose_test`.

### 2. Install pressly/goose

	$ go get -u github.com/pressly/goose/cmd/goose
	
### 3. Move to Migrations Directory

	$ cd [app_root]/database/migrations
	
### 4. Run Migrations

To create/update the database schemas, run the `up` command like the following example:

	# Database name: pressly_goose_test
	# User: root, password: (not set)
	# Host: localhost, port: 3306
	
	$ goose mysql "root@/pressly_goose_test?parseTime=true" up
	
To roll it back, run the `down` command one or more times until it reach the previous state:

	$ goose mysql "root@/pressly_goose_test?parseTime=true" down

	
## Run

After running the steps above, you can run the service:

	$ ./pressly-goose-test

***
&copy; 2018 Kudo.