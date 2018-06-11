# Dancesport Application System (DAS)

[Join the DAS Slack Community](https://ballroomdev.slack.com)

[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![GoDoc](https://godoc.org/github.com/DancesportSoftware/das?status.svg)](https://godoc.org/github.com/DancesportSoftware/das)
[![Build Status](https://travis-ci.org/DancesportSoftware/das.svg?branch=development)](https://travis-ci.org/DancesportSoftware/das)
[![codecov](https://codecov.io/gh/DancesportSoftware/das/branch/development/graph/badge.svg)](https://codecov.io/gh/DancesportSoftware/das)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/76bb2da0aa0e4a2486365500d3f93e8f)](https://www.codacy.com/app/DancesportSoftware/das_2?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=DancesportSoftware/das&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/DancesportSoftware/das)](https://goreportcard.com/report/github.com/DancesportSoftware/das)
DAS is an open and free competition management system for competitive ballroom
dance. This project (along with [dasdb](https://github.com/DancesportSoftware/dasdb) and 
[dasweb](https://github.com/DancesportSoftware/dasweb)) aims to provide the dancesport community
an open and secure implementation of competition management system.

### Goals of DAS ###
* To provide a secure and robust solution to competition organizers and competitors. Currently, the most popular systems do not have the necessary security setup to protect itself from being compromised. Data security is the top priority of this project.

* To provide a free competition management solution to most amateur and collegiate competition organizers. A competition management system should not require a specially-trained technician to be useful for the organizer. Competition organizer’s limited budget should be spent on renting a great venue, inviting unbiased and professional adjudicators, and promote dancesport in society.

* To provide insightful information to competitors, including but not limited to:
  * Dancer and Couple profiles and statistics
  * Dancer and Couple rating, ranking, and progression
  * Adjudicator preference
  
* To help organizers manage competitions efficiently:
  * A more intuitive user interface for all users
  * Small functions that improve quality of life:
    * Set up typical collegiate and amateur competitions quickly
    * Competition finance management: managing entries, tickets, and lessons sales
    * Compatible with major federations’ requirements: WDSF, WDC, USA Dance, NDCA, and Collegiate (North America)
  * TBA partner search and matchmaking

* To provide opportunities for software developers and data enthusiasts:
  * API for developing custom client applications (competition results, statistics, etc.)
  * API for competition data and statistics

## Local Development Setup
DAS is mostly developed for running on Linux platform. Though
it is totally possible to develop and run on Windows, it's not tested
as thoroughly as Linux. We assumes that you are using Ubuntu. Most of
the setup works for Mac as well.

1. Install Go

   You can install Golang SDK through `apt-get`: `$ sudo apt-get install golang-go`
   
   Detailed documentation can be found [here](https://github.com/golang/go/wiki/Ubuntu).
   
2. Go environment

   You need to add `$GOPATH` to your environment (`~/.profile`): 
   
   `export GOPATH=$HOME/go`
   
   Sometimes we will need to use binaries built from other packages:
   
   `export PATH=$GOPATH/bin:$PATH`
   
   Make sure you define `$GOPATH` before adding `$GOPATH/bin` to your `PATH`. Double check
   if they are defined by running `$ echo $GOPATH`.
   
3. Check out the repository

   First, we need to create necessary directory for DAS:
   
   `$ mkdir -p $GOPATH/src/github.com/DancesportSoftware`
   
   Change directory:

   `$ cd $GOPATH/src/github.com/DancesportSoftware`
   
   Check out the latest build:
   
   `$ git clone https://github.com/DancesportSoftware/das.git`
   
4. Get dependencies

   Most of the dependencies can be get by `go get`:
   1. Change directory to project root: `$ cd $GOPATH/src/DancesportSoftware/das`
   2. Get dependencies `$ go get ./...`
   
   This will download all the dependencies from online repositories.
   
5. Run DAS

   You will need to have the database set up in order to run DAS locally. You can
   visit the [dasdb](https://github.com/DancesportSoftware/dasdb) to build the database schema
   for local development.
   
   To run DAS services locally, run `$ go run das.go` in the root of DAS repository.
   
   If you want to build a binary and run separately:
   
   `$ cd $GOPATH/src/github.com/DancesportSoftware/das`
   
   `$ go install`
   
   If you have added `$GOPATH/bin` to your `PATH`, then there should be a command available
   to you: `$ das`, which is installed under `$GOPATH/bin`.
   
   DAS will run on port 8080: `localhost:8080`.
   
## Necessary Development Tools
* IDE: IntelliJ (with Golang plugin), Goland (requires subscription), and VS Code
* Web Service: Postman, SoapUI

## Design and Development Principle.(Technical)

* **MVC-Based**: The architecture of DAS is inspired by design philosophies of ASP.NET MVC and .NET MVC Core.
If you are familiar with ASP.NET, it should be very easy to get familiar with the architecture of DAS.

* **Clean Architecture**: SOLID principles and Uncle Bob's *Clean Architecture* are heavily applied in DAS. We want the code easy to understand
for developers and easy to maintain. For example, `businesslogic` has no dependency on `dataaccess` or `controller` package. 
`controller` does not directly depend on `dataaccess`. This allows `businesslogic` to be developed
relatively independently without worrying about how it is going to be accessed by users or 
data sources. Similarly, code in `controller` can be changed more freely without concerning with
`dataaccess`. ISP allows the independent changes to different modules. Finally, `config` package
glue everything together and `das.go` gets the entire system started.

* **Test**, but not always driven by it: critical code should be tested as thoroughly as possible. There is
no hard requirement for test coverage, but we do our best to make sure the code executes correctly most of 
the time.