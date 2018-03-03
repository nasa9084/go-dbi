package dbi

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

// database/sql errors
var (
	ErrConnDone = sql.ErrConnDone
	ErrNoRows   = sql.ErrNoRows
	ErrTxDone   = sql.ErrTxDone
)

// database/sql functions
var (
	Drivers  = sql.Drivers
	Register = sql.Register
)

type (
	// TxOptions is type alias of sql.TxOptions.
	TxOptions = sql.TxOptions
	// DBStats is type alias of sql.DBStats.
	DBStats = sql.DBStats
)

type dbImpl struct {
	*sql.DB
}

func Open(driverName, dataSourceName string) (DB, error) {
	sqldb, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &dbImpl{DB: sqldb}, nil
}

func OpenDB(c driver.Connector) DB {
	return &dbImpl{DB: sql.OpenDB(c)}
}

func (db *dbImpl) Begin() (Tx, error) {
	sqltx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &txImpl{Tx: sqltx}, nil
}

func (db *dbImpl) BeginTx(ctx context.Context, opts *TxOptions) (Tx, error) {
	sqltx, err := db.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &txImpl{Tx: sqltx}, nil
}

func (db *dbImpl) Close() error {
	return db.DB.Close()
}

func (db *dbImpl) Conn(ctx context.Context) (Conn, error) {
	sqlconn, err := db.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	return &connImpl{Conn: sqlconn}, nil
}

func (db *dbImpl) Driver() driver.Driver {
	return db.DB.Driver()
}

func (db *dbImpl) Exec(query string, args ...interface{}) (Result, error) {
	sqlresult, err := db.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &resultImpl{Result: sqlresult}, nil
}

func (db *dbImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error) {
	sqlresult, err := db.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &resultImpl{Result: sqlresult}, nil
}

func (db *dbImpl) Ping() error {
	return db.DB.Ping()
}

func (db *dbImpl) PingContext(ctx context.Context) error {
	return db.DB.PingContext(ctx)
}

func (db *dbImpl) Prepare(query string) (Stmt, error) {
	sqlstmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &stmtImpl{Stmt: sqlstmt}, nil
}

func (db *dbImpl) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	sqlstmt, err := db.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &stmtImpl{Stmt: sqlstmt}, nil
}

func (db *dbImpl) Query(query string, args ...interface{}) (Rows, error) {
	sqlrows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return &rowsImpl{Rows: sqlrows}, nil
}

func (db *dbImpl) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	sqlrows, err := db.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &rowsImpl{Rows: sqlrows}, nil
}

func (db *dbImpl) QueryRow(query string, args ...interface{}) Row {
	return &rowImpl{Row: db.DB.QueryRow(query, args...)}
}

func (db *dbImpl) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	return &rowImpl{Row: db.DB.QueryRowContext(ctx, query, args...)}
}

func (db *dbImpl) SetConnMaxLifetime(d time.Duration) {
	db.DB.SetConnMaxLifetime(d)
}

func (db *dbImpl) SetMaxIdleConns(n int) {
	db.DB.SetMaxIdleConns(n)
}

func (db *dbImpl) SetMaxOpenConns(n int) {
	db.DB.SetMaxOpenConns(n)
}

func (db *dbImpl) Stats() DBStats {
	return db.DB.Stats()
}

type txImpl struct {
	Tx *sql.Tx
}

func (tx *txImpl) Commit() error {
	return tx.Tx.Commit()
}

func (tx *txImpl) Exec(query string, args ...interface{}) (Result, error) {
	sqlresult, err := tx.Tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &resultImpl{Result: sqlresult}, nil
}

func (tx *txImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error) {
	sqlresult, err := tx.Tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &resultImpl{Result: sqlresult}, nil
}

func (tx *txImpl) Prepare(query string) (Stmt, error) {
	sqlstmt, err := tx.Tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &stmtImpl{Stmt: sqlstmt}, nil
}

func (tx *txImpl) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	sqlstmt, err := tx.Tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &stmtImpl{Stmt: sqlstmt}, nil
}

func (tx *txImpl) Query(query string, args ...interface{}) (Rows, error) {
	sqlrows, err := tx.Tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return &rowsImpl{Rows: sqlrows}, nil
}

func (tx *txImpl) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	sqlrows, err := tx.Tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &rowsImpl{Rows: sqlrows}, nil
}

func (tx *txImpl) QueryRow(query string, args ...interface{}) Row {
	return &rowImpl{Row: tx.Tx.QueryRow(query, args...)}
}

func (tx *txImpl) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	return &rowImpl{Row: tx.Tx.QueryRow(query, args...)}
}

func (tx *txImpl) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx *txImpl) Stmt(stmt Stmt) Stmt {
	sqlstmt := tx.Tx.Stmt(stmt.(*stmtImpl).Stmt)
	return &stmtImpl{Stmt: sqlstmt}
}

func (tx *txImpl) StmtContext(ctx context.Context, stmt Stmt) Stmt {
	sqlstmt := tx.Tx.Stmt(stmt.(*stmtImpl).Stmt)
	return &stmtImpl{Stmt: sqlstmt}
}

type stmtImpl struct {
	Stmt *sql.Stmt
}

func (stmt *stmtImpl) Close() error {
	return stmt.Stmt.Close()
}

func (stmt *stmtImpl) Exec(args ...interface{}) (Result, error) {
	sqlresult, err := stmt.Stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	return &resultImpl{Result: sqlresult}, nil
}

func (stmt *stmtImpl) ExecContext(ctx context.Context, args ...interface{}) (Result, error) {
	sqlresult, err := stmt.Stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return &resultImpl{Result: sqlresult}, nil
}

func (stmt *stmtImpl) Query(args ...interface{}) (Rows, error) {
	sqlrows, err := stmt.Stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return &rowsImpl{Rows: sqlrows}, nil
}

func (stmt *stmtImpl) QueryContext(ctx context.Context, args ...interface{}) (Rows, error) {
	sqlrows, err := stmt.Stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return &rowsImpl{Rows: sqlrows}, nil
}

func (stmt *stmtImpl) QueryRow(args ...interface{}) Row {
	return &rowImpl{Row: stmt.Stmt.QueryRow(args...)}
}

func (stmt *stmtImpl) QueryRowContext(ctx context.Context, args ...interface{}) Row {
	return &rowImpl{Row: stmt.Stmt.QueryRowContext(ctx, args...)}
}

type resultImpl struct {
	Result sql.Result
}

func (result *resultImpl) LastInsertId() (int64, error) {
	return result.Result.LastInsertId()
}

func (result *resultImpl) RowsAffected() (int64, error) {
	return result.Result.RowsAffected()
}

type rowsImpl struct {
	Rows *sql.Rows
}

func (rows *rowsImpl) Close() error {
	return rows.Rows.Close()
}

func (rows *rowsImpl) ColumnTypes() ([]ColumnType, error) {
	sqlcoltypes, err := rows.Rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	ret := make([]ColumnType, len(sqlcoltypes))
	for i, coltype := range sqlcoltypes {
		ret[i] = coltype
	}
	return ret, nil
}

func (rows *rowsImpl) Columns() ([]string, error) {
	return rows.Rows.Columns()
}

func (rows *rowsImpl) Err() error {
	return rows.Rows.Err()
}

func (rows *rowsImpl) Next() bool {
	return rows.Rows.Next()
}

func (rows *rowsImpl) NextResultSet() bool {
	return rows.Rows.NextResultSet()
}

func (rows *rowsImpl) Scan(dest ...interface{}) error {
	return rows.Rows.Scan(dest...)
}

type rowImpl struct {
	Row *sql.Row
}

func (row *rowImpl) Scan(dest ...interface{}) error {
	return row.Row.Scan(dest...)
}

type connImpl struct {
	Conn *sql.Conn
}

func (conn *connImpl) BeginTx(ctx context.Context, opts *TxOptions) (Tx, error) {
	tx, err := conn.Conn.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &txImpl{Tx: tx}, nil
}

func (conn *connImpl) Close() error {
	return conn.Conn.Close()
}

func (conn *connImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error) {
	sqlresult, err := conn.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &resultImpl{Result: sqlresult}, nil
}

func (conn *connImpl) PingContext(ctx context.Context) error {
	return conn.Conn.PingContext(ctx)
}

func (conn *connImpl) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	sqlstmt, err := conn.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &stmtImpl{Stmt: sqlstmt}, nil
}

func (conn *connImpl) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	sqlrows, err := conn.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &rowsImpl{Rows: sqlrows}, nil
}

func (conn *connImpl) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	return &rowImpl{Row: conn.Conn.QueryRowContext(ctx, query, args...)}
}
