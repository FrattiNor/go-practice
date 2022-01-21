package gormSample

import (
	"database/sql"
	"goServer/gormSample/session"
	"goServer/logSample"
)

type Engine struct {
	db *sql.DB
}

func New(driver string, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		logSample.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		logSample.Error(err)
		return
	}
	e = &Engine{db: db}
	logSample.Info("connect success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		logSample.Error(err)
	}
	logSample.Info("close success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
