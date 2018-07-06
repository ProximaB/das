# DAS Architecture
This document explains the design of DAS from the architecture perspective in the hope
that it will help developer understand and follow the software design to make this
project maintainable in the long run.

## Design Philosophy
Many sources of thoughts have an impact on the design and implementation
of DAS. The primary ones are:
* Clean Architecture (Architecture)
* SOLID Principles (Architecture and Implementation)
* .NET MVC (Implementation)

Therefore, if you are familiar with the Clean Architecture and SOLID principles, it should 
not be difficult to navigate and understand the responsibilities of packages in DAS.

## Architectural Overview
There are three major components in DAS and their responsibilities are:
* Business Logic (__businesslogic__ package)
  * Competition management at an abstract level
* Data Access (__dataaccess__ package)
  * Store and provide data that are requested by Business Logic
* Controller (__controller__ package)
  * Handles HTTP requests and instruct Business Logic to perform actions

*businesslogic* is the *core* of DAS and has no dependency on any other packages in
DAS. Such design follows the clean architecture and the SOLID principles. *businesslogic*
is *the system* thus has no concern with lower level operations (such as I/O or interaction
with the persistence layer). *businesslogic* specifies the interface that any persistence layer
should implement.

*dataaccess* is *an implementation* of the repository interfaces that are specified in 
*businesslogic*. This implementation primarily uses PostgreSQL database
as the main storage mechanism. Ideally, if there is better storage mechanism than 
relational database, we can easily implement those interfaces using the newer technology
and DAS should still function.

*controller* is *an implementation* of any controller that can be used to manipulate the
operation of *businesslogic*. This implementation primarily focuses on the interaction
with HTTP. It is totally possible to implement a different set of controller that 
makes DAS a standalone desktop application.

There are a few minor components that either provides utility functions (such as *util* and *viewmodel*) or work as a
"glue" package that pack up the entire system for shipping (such as *config*).