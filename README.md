go-dbi
======

interface wrapper package of database/sql package.

this package only define interfaces which has same functions with structs in database/sql and implements wrapper structs.

there's no relation to perl's DBI.

## motivation

most of database/sql package implemented with structs, not interfaces in spite of there's no exported members.
so it's difficult to mocking database connection for tests.

## how to use

this package has not support some structs below in database/sql package.

* IsolationLevel
* NamedArg
* Null*
* Out
* RawBytes
* Scanner

if you're not using these structs, you can replace database/sql into this package.
