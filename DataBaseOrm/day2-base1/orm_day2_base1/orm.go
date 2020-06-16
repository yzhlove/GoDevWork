package orm_day2_base1

import (
	"database/sql"
	"orm_day2_base1/dialect"
	"orm_day2_base1/log"
	"orm_day2_base1/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.ERROR(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.ERROR(err)
		return
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.ERRORF("dialect %s not found ", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.INFO("connection database success")
	return
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.ERROR("close db err:", err)
	}
	log.INFO("close db succeed")
}
