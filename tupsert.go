package codegen

import "strings"

func contains(stringSlice []string, search string) bool {
	for _, value := range stringSlice {
		if value == search {
			return true
		}
	}
	return false
}

// TUpsert template
func TUpsert(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {

	// generate upsert statement
	bb.Line("// ", table.lower, "UpsertStmt helper for generating Upserts general statement")
	bb.Line("// nolint[gocyclo]")
	bb.Func(table.StoreReceiver, table.lower+"UpsertStmt")
	bb.FuncParams()
	bb.FuncReturn("*sdb.UpsertStatement")

	bb.Line("upsert := []string{}")
	{
		for _, f := range table.otherFields {
			if !contains(table.Ignores.Upsert, f.Name) {
				bb.Line(`if `, table.Initials, `.colSet == nil || `, table.Initials, `.colSet.Bit(`, table.Title+f.Title, `) == 1 {`)
				bb.Line(`upsert = append(upsert, "`, f.Name, " = VALUES(", f.Name, `)")`)
				bb.Line("}")
			}
		}
	}

	bb.Line("sql := &sdb.UpsertStatement{}")
	bb.Line(`sql.InsertInto("`, schema.Name, ".", table.Name, `")`)
	bb.S(`sql.Columns(`)
	{
		for _, f := range table.Fields {
			bb.S(`"`)
			bb.S(strings.ToLower(f.Name))
			bb.S(`",`)
		}
	}
	bb.Line(`)`)
	bb.Line(`sql.OnDuplicateKeyUpdate(upsert)`)
	bb.Line(`return sql`)
	bb.Line(`}`)

	// Upsert for a single record
	bb.Line("// UpsertOne inserts the ", table.Title, " to the database.")
	bb.Func(table.StoreReceiver, "UpsertOne")
	bb.FuncParams("data *" + table.Title)
	bb.FuncReturn("int64", "error")
	bb.Line("return ", table.Initials, ".Upsert([]*", table.Title, "{data})")
	bb.Line("}")

	// upsert for data array
	bb.Line("// Upsert executes upsert for array of ", table.Title)
	bb.Func(table.StoreReceiver, "Upsert")
	bb.FuncParams("data []*" + table.Title)
	bb.FuncReturn("int64", "error")

	bb.Line(`sql := `, table.Initials, ".", table.lower, `UpsertStmt()
	
	for _, d := range data {
		sql.Record(d)
	}

	if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {
		log.Debug().Str("fn", "`, table.Title, `Upsert").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := `, table.Initials, `.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return -1, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("rowsaffected")
		return -1, err
	}

	return affected, nil
}
`)

}
