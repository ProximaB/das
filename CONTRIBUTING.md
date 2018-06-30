# Contributing to DAS (Draft)

First off, thank you very much for taking the time to contribute!

#### Table of Contents

## How to Participate

## What should I know to get started?

### Competitive Ballroom Dance
DAS is a system designed for competitive ballroom dance, though it may be utilized for non-dancesport events as well.

### The Go Programming Language

### .NET MVC
**MVC-Based**: The architecture of DAS is inspired by design philosophies of ASP.NET MVC and .NET MVC Core.
If you are familiar with ASP.NET, it should be very easy to get familiar with the architecture of DAS.


### Design Patterns and Clean Architecture
* **Clean Architecture**: SOLID principles and Uncle Bob's *Clean Architecture* are heavily applied in DAS. We want the code easy to understand
for developers and easy to maintain. For example, `businesslogic` has no dependency on `dataaccess` or `controller` package. 
`controller` does not directly depend on `dataaccess`. This allows `businesslogic` to be developed
relatively independently without worrying about how it is going to be accessed by users or 
data sources. Similarly, code in `controller` can be changed more freely without concerning with
`dataaccess`. ISP allows the independent changes to different modules. Finally, `config` package
glue everything together and `das.go` gets the entire system started.

### Test
**Test**, but not always driven by it: critical code should be tested as thoroughly as possible. There is
no hard requirement for test coverage, but we do our best to make sure the code executes correctly most of 
the time.