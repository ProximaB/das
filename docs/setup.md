# Software Version
Unlike most projects, DAS is very responsive to the change of the version its dependencies. Unless the dependency on a particular version is 
so strong that it's impossible to refactor the code without breaking the core functions. Therefore, it is almost 100% safe to assume that DAS works 
well with the latest version of its dependencies. 

On the other hand, hosting environment does not always provide the most up-to-date versions of software. Therefore, DAS must remain compatible with older versions of its dependencies until the hosting environment is updated.

Acceptable versions are listed with each component.

# Necessary Software for Development
* Go (1.8 or later)
  * Installation
      
      To install Go SDK, please follow the instruction [here](https://golang.org/doc/install).
      Once the installation is complete, please check in your terminal to see if `$GOROOT` and `$GOPATH` are defined. 
      You can check it by running `$ echo $GOROOT` and `$ echo $GOPATH` in Linux.
* PostgreSQL (9.5 or later)
* Git (2.0 or later)
* IDE

  This is totally up to your personal preference.
  * Go
    * GoLand (JetBrains subscription required)
    * Visual Studio Code (free)
  * Database
    * Datagrip (subscription required)
    * TeamSQL (free version available)
  * Web Services
    * Postman (free)
    * SoapUI

# Environment Setup
* Go

  You must have `$GOROOT` and `$GOPATH` defined. If not, add following environment variables to your system:
  * On Linux (or macOS?)

    Add `$GOPATH` to your environment (`~/.profile`):  `export GOPATH=$HOME/go`
  
    Sometimes we will need to use binaries built from other packages:
  
    `export PATH=$GOPATH/bin:$PATH`
  
    Make sure you define `$GOPATH` before adding `$GOPATH/bin` to your `PATH`. Double check
  if they are defined by running `$ echo $GOPATH`.

  * On Windows
    Add `GOPATH` and `GOROOT` under System variables. On Windows, your `GOPATH` should be
    `C:\Users\yourUserName\go`, and `GOROOT` should be `C:\Go\`

* PostgreSQL

  You must have the administrator privileges on your computer.

  You can login as user `postgres` to create role, database, etc.

  * Login PostgreSQL
    * `$ psql -U postgres`

      Enter the password of `postgres` when prompted.

  * Create a new role `dasdev` with default password:
    * `postgres=# CREATE USER dasdev WITH PASSWORD `dAs!@#$1234`;`

    Here we use the default password that DAS uses. If you are using a different
    password, you may need to modify corresponding code in DAS as well.
  * Create a new database `das`
    * `postgres=# CREATE DATABASE das;`
  * Grant all privileges to `dasdev` on `das`:
    * `postgres=# GRANT ALL PRIVILEGES ON DATABASE das TO dasdev;`
  * Environment variable
    * In your IDE, or terminal, you need to export a connection string for DAS to use.
    If you follow the instructions above, your should export `POSTGRES_CONNECTION` to
    your environment variable:
    `$ export POSTGRES_CONNECTION=user=dasdev password=dAs\!@#\$1234 dbname=das sslmode=disable`. 
    If necessary, add this environment variable to your `~/.profile`

# Source Code Compilation and Run
* Check out the repository
  * dasdb
    * Change to the directory that you would like to keep the repository
    * Clone the repository:
      * `$ git clone https://github.com/ProximaB/dasdb`
    * Build the database schema and populate data:
      * Change directory: `$ cd dasdb`
      * Build the database: `$ psql -U dasdev -d das -f build.sql`
  * das
    * Create directory for development:
      * Windows: `%GOPATH%\src\github.com\DancesportSoftware\`
      * Linux: `$GOPATH/src/github.com/DancesportSoftware/`
    * Change directory:
      * Windows: `C:\> cd %GOPATH%\src\github.com\DancesportSoftware`
      * Linux: `$ cd $GOPATH/src/github.com/DancesportSoftware`
    * Clone the repository:
      * `$ git clone https://github.com/ProximaB/das`
    * Get all dependencies:
      * Change directory to project root: `$ cd das`
      * Change branches to development: `$ git checkout development`
      * Get all dependencies, this can take a few minutes: `$ go get ./...`
      * Run DAS: `$ go run das.go`
        * You should see that DAS can connect to the database and run on
        `localhost:8080` (it may be 404 page not found that `localhost:8080`)