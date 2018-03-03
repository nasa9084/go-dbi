package dbi

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"time"
)

// type aliases
type (
	TxOptions      = sql.TxOptions
	DBStats        = sql.DBStats
	IsolationLevel = sql.IsolationLevel
	NamedArgs      = sql.NamedArg
	NullBool       = sql.NullBool
	NullFloat64    = sql.NullFloat64
	NullInt64      = sql.NullInt64
	NullString     = sql.NullString
	Out            = sql.Out
	RawBytes       = sql.RawBytes
	Scanner        = sql.Scanner
	Result         = sql.Result
)

// const
const (
	LevelDefault IsolationLevel = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
	LevelSerializable
	LevelLinearizable
)

type Execer interface {
	Exec(query string, args ...interface{}) (Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)
}

type Preparer interface {
	Prepare(query string) (Stmt, error)
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

type Queryer interface {
	Query(query string, args ...interface{}) (Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row
	QueryRowContext(ctx context.Context, vquery string, args ...interface{}) Row
}

type DB interface {
	Begin() (Tx, error)
	BeginTx(context.Context, *TxOptions) (Tx, error)
	Close() error
	Conn(context.Context) (Conn, error)
	Driver() driver.Driver
	Execer
	Ping() error
	PingContext(context.Context) error
	Preparer
	Queryer
	SetConnMaxLifetime(time.Duration)
	SetMaxIdleConns(int)
	SetMaxOpenConns(int)
	Stats() DBStats
}

type Tx interface {
	Commit() error
	Execer
	Preparer
	Queryer
	Rollback() error
	Stmt(Stmt) Stmt
	StmtContext(context.Context, Stmt) Stmt
}

type Stmt interface {
	Close() error
	Exec(args ...interface{}) (Result, error)
	ExecContext(ctx context.Context, args ...interface{}) (Result, error)
	Query(args ...interface{}) (Rows, error)
	QueryContext(ctx context.Context, args ...interface{}) (Rows, error)
	QueryRow(args ...interface{}) Row
	QueryRowContext(ctx context.Context, args ...interface{}) Row
}

type Rows interface {
	Close() error
	ColumnTypes() ([]ColumnType, error)
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Scan(dest ...interface{}) error
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Conn interface {
	BeginTx(context.Context, *TxOptions) (Tx, error)
	Close() error
	ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)
	PingContext(context.Context) error
	PrepareContext(ctx context.Context, query string) (Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) Row
}

type ColumnType interface {
	DatabaseTypeName() string
	DecimalSize() (precision, scale int64, ok bool)
	Length() (length int64, ok bool)
	Name() string
	Nullable() (nullable, ok bool)
	ScanType() reflect.Type
}
