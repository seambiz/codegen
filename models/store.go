package models

import (
	"database/sql"
	"math/big"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// nolint[structcheck]
type Store struct {
	db           Execer
	withJoin     bool
	joinType     string
	where        string
	orderBy      string
	groupBy      string
	prependField string
	limit        int
	offset       int
	batch        int
	colSet       *big.Int
	stmt         *SQLStatement

	// TODO vielleicht hier weg, da unnötig bezogen auf stmt. Nur für wirkliche Store Methoden relevant
	dtos     []DTO
	dtoslice *Result
}

type Row []DTO
type Result []Row

// NewStore return DAO Storr
func NewStore(conn *sql.Tx) *Store {
	s := &Store{}
	s.db = conn
	s.stmt = NewSQLStatement()
	return s
}

func (s *Store) BindSlice(res *Result, rowStructs Row) *Store {
	for i := range rowStructs {
		s.dtos = append(s.dtos, rowStructs[i])
	}
	s.dtoslice = res
	return s
}
func (s *Store) Bind(datas ...DTO) *Store {
	for i := range datas {
		s.dtos = append(s.dtos, datas[i])
	}
	return s
}
func (s *Store) Select(columns string) *Store {
	s.stmt.Append("SELECT", columns)
	return s
}
func (s *Store) From(table string) *Store {
	s.stmt.Append("FROM", table)
	return s
}
func (s *Store) Join(table, condition string) *Store {
	s.stmt.Append("INNER JOIN", table, "ON", condition)
	return s
}
func (s *Store) LeftJoin(table, condition string) *Store {
	s.stmt.Append("LEFT JOIN", table, "ON", condition)
	return s
}
func (s *Store) RightJoin(table, condition string) *Store {
	s.stmt.Append("RIGHT JOIN", table, "ON", condition)
	return s
}
func (s *Store) Where(condition string) *Store {
	s.stmt.Append("WHERE", condition)
	return s
}
func (s *Store) OrderBy(columns string) *Store {
	s.stmt.Append("ORDER BY", columns)
	return s
}
func (s *Store) Limit(limits ...int) *Store {
	if len(limits) == 0 {
		return s
	}
	s.stmt.Append("LIMIT", limits[0])
	if len(limits) > 1 {
		s.stmt.Append(",", limits[1])
	}
	return s
}

// Columns to be used for various statements.
func (s *Store) Columns(cols ...int) *Store {
	if s.colSet == nil {
		s.colSet = big.NewInt(0)
	}
	for _, col := range cols {
		s.colSet.SetBit(s.colSet, col, 1)
	}
	return s
}

func (s *Store) Fields(prefix string, fieldFunc func(*big.Int) []string) *Store {
	s.stmt.Fields(s.prependField, prefix, fieldFunc(s.colSet))
	s.prependField = ","
	return s
}

func (s *Store) queryBegin(stmt string, args ...interface{}) (*sql.Rows, []sql.RawBytes, []interface{}, error) {
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "Store.One").Str("stmt", stmt).Interface("args", args).Msg("sql")
	}
	rows, err := s.db.Query(stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("query")
		return nil, nil, nil, err
	}
	colNames, err := rows.Columns()
	if err != nil {
		log.Error().Err(err).Msg("colNames")
		return nil, nil, nil, err
	}
	columnValues := make([]sql.RawBytes, len(colNames))
	columnPointers := make([]interface{}, len(colNames))
	for i := range columnValues {
		columnPointers[i] = &columnValues[i]
	}
	return rows, columnValues, columnPointers, err
}

// One retrieves a row from 'best.bestellung' as a Bestellung with possible joined data.
func (s *Store) one(data DTO, stmt *SQLStatement, args ...interface{}) error {
	rows, values, valuePointers, err := s.queryBegin(stmt.Query(), args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	if rows.Next() {
		err = rows.Scan(valuePointers...)
		if err != nil {
			log.Error().Err(err).Msg("scan")
			return err
		}
		col := 0
		data.bind(values, s.withJoin, s.colSet, &col)
	} else {
		return sql.ErrNoRows
	}
	return err
}

// One retrieves a row from 'best.bestellung' as a Bestellung with possible joined data.
func (s *Store) MapScan(dest map[string]sql.RawBytes, args ...interface{}) error {
	rows, values, valuePointers, err := s.queryBegin(s.stmt.Query(), args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	if rows.Next() {
		err = rows.Scan(valuePointers...)
		if err != nil {
			log.Error().Err(err).Msg("scan")
			return err
		}

		var columns []string
		columns, err = rows.Columns()
		if err != nil {
			log.Error().Err(err).Msg("columns")
			return err
		}

		for i, column := range columns {
			dest[column] = values[i]
		}
	} else {
		return sql.ErrNoRows
	}
	return err
}

// One retrieves a row from 'best.bestellung' as a Bestellung with possible joined data.
func (s *Store) One(args ...interface{}) error {
	rows, values, valuePointers, err := s.queryBegin(s.stmt.Query(), args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	if rows.Next() {
		err = rows.Scan(valuePointers...)
		if err != nil {
			log.Error().Err(err).Msg("scan")
			return err
		}
		col := 0
		for i := range s.dtos {
			s.dtos[i].bind(values, s.withJoin, s.colSet, &col)
		}
	} else {
		return sql.ErrNoRows
	}
	return err
}

// QueryCustom retrieves many rows from 'best.bestellung' as a slice of Bestellung with possible joined data.
func (s *Store) queryCustom(res DTOSlice, d DTO, stmt string, args ...interface{}) error {
	rows, values, valuePointers, err := s.queryBegin(stmt, args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(valuePointers...)
		if err != nil {
			log.Error().Err(err).Msg("scan")
			return err
		}
		data := d.new()
		col := 0
		data.bind(values, s.withJoin, s.colSet, &col)
		res.append(data)
	}
	return err
}

// QueryCustom retrieves many rows from 'best.bestellung' as a slice of Bestellung with possible joined data.
func (s *Store) Query(args ...interface{}) error {
	rows, values, valuePointers, err := s.queryBegin(s.stmt.Query(), args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(valuePointers...)
		if err != nil {
			log.Error().Err(err).Msg("scan")
			return err
		}
		col := 0
		row := Row{}
		for i := range s.dtos {
			d := s.dtos[i].new()
			d.bind(values, false, s.colSet, &col)
			row = append(row, d)
		}
		*s.dtoslice = append(*s.dtoslice, row)
	}
	return err
}