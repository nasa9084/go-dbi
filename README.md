go-dbi
======

interface wrapper package of database/sql package.

this package only define interfaces which has same functions with structs in database/sql and implements wrapper structs.

there's no relation to perl's DBI.

## motivation

most of database/sql package implemented with structs, not interfaces in spite of there's no exported members.
so it's difficult to mocking database connection for tests.

## how to use

you can replace database/sql into this package like below.

``` go
import sql "github.com/nasa9084/go-dbi"
```
