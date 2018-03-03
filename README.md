go-dbi
======

interface wrapper package of database/sql package.

this package only define interfaces which has same functions with structs in database/sql and implements wrapper structs.

## motivation

most of database/sql package implemented with structs, not interfaces.
so it's difficult to mocking database connection for tests.
