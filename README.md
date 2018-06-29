# DAS
An open dancesport competition management system.


[Join the DAS Slack Community](https://join.slack.com/t/ballroomdev/shared_invite/enQtMzg4OTU4OTAyNjQ3LTZmYWU2ZTA1Njc2NmI2YWIyZGRlMTQ1MGYyODM0ZWVmZjAzN2ZkNTcyYzNiM2NiMjE2MjI4YjQyNjcyNTU1MGE)


[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![GoDoc](https://godoc.org/github.com/DancesportSoftware/das?status.svg)](https://godoc.org/github.com/DancesportSoftware/das)
[![Build Status](https://travis-ci.org/DancesportSoftware/das.svg?branch=development)](https://travis-ci.org/DancesportSoftware/das)
[![codecov](https://codecov.io/gh/DancesportSoftware/das/branch/development/graph/badge.svg)](https://codecov.io/gh/DancesportSoftware/das)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/76bb2da0aa0e4a2486365500d3f93e8f)](https://www.codacy.com/app/DancesportSoftware/das_2?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=DancesportSoftware/das&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/DancesportSoftware/das)](https://goreportcard.com/report/github.com/DancesportSoftware/das)
[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=das&metric=alert_status)](https://sonarcloud.io/api/project_badges/measure?project=das&metric=alert_status)

DAS is an open and free competition management system for competitive ballroom
dance. This project (along with [dasdb](https://github.com/DancesportSoftware/dasdb) and 
[dasweb](https://github.com/DancesportSoftware/dasweb)) aims to provide the dancesport community
an open and secure implementation of competition management system.

### Development Setup

A complete guide can be found [here](https://github.com/DancesportSoftware/das/wiki/Development-Setup).

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
  
## API Documents
[Postman Document](https://documenter.getpostman.com/view/2986351/RWEfMemn)

   

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