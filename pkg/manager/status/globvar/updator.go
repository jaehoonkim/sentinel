package globvar

import (
	"context"
	"database/sql"
	"time"

	// "github.com/NexClipper/sudory/pkg/manager/control/vault"

	"github.com/NexClipper/logger"
	"github.com/NexClipper/sudory/pkg/manager/database/vanilla"
	"github.com/NexClipper/sudory/pkg/manager/database/vanilla/excute"
	"github.com/NexClipper/sudory/pkg/manager/database/vanilla/stmt"
	"github.com/NexClipper/sudory/pkg/manager/macro"
	"github.com/NexClipper/sudory/pkg/manager/macro/logs"
	globvarv2 "github.com/NexClipper/sudory/pkg/manager/model/global_variables/v2"
	"github.com/NexClipper/sudory/pkg/manager/status/state"
	"github.com/pkg/errors"
)

type GlobalVariableUpdate struct {
	*sql.DB
	dialect excute.SqlExcutor
	offset  time.Time //updated column
}

func NewGlobalVariablesUpdate(db *sql.DB, dialect excute.SqlExcutor) *GlobalVariableUpdate {
	return &GlobalVariableUpdate{
		DB:      db,
		dialect: dialect,
	}
}

// func (worker *GlobalVariableUpdate) Dialect() string {
// 	return worker.dialect
// }

// Update
//
//	Update = read db -> global_variables
func (worker *GlobalVariableUpdate) Update() (err error) {
	records := make([]globvarv2.GlobalVariables, 0, state.ENV__INIT_SLICE_CAPACITY__())
	globvar := globvarv2.GlobalVariables{}
	globvar.Updated = *vanilla.NewNullTime(worker.offset)

	globvar_cond := stmt.GT("updated", globvar.Updated)

	err = worker.dialect.QueryRows(globvar.TableName(), globvar.ColumnNames(), globvar_cond, nil, nil)(
		context.Background(), worker)(
		func(scan excute.Scanner, _ int) error {
			err := globvar.Scan(scan)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}

			records = append(records, globvar)

			return err
		})
	if err != nil {
		return
	}

	for i := range records {
		record := &records[i]
		gv, err := ParseKey(record.Name)

		switch err {
		case nil:
			if err := storeManager.Call(gv, record.Value.String); err != nil {
				return errors.Wrapf(err, "store global_variables%v",
					logs.KVL(
						"key", record.Name,
						"value", record.Value.String,
					))
			}
		default:
			logger.Warningf("%v: parse record name to key%v", err.Error(), logs.KVL(
				"key", record.Name,
			))
		}
	}

	//update offset
	worker.offset = time.Now()

	return
}

// WhiteListCheck
//
//	리스트 체크
func (worker *GlobalVariableUpdate) WhiteListCheck() (err error) {
	records := make([]globvarv2.GlobalVariables, 0, state.ENV__INIT_SLICE_CAPACITY__())

	globvar := globvarv2.GlobalVariables{}
	globvar.Updated = *vanilla.NewNullTime(worker.offset)
	globvar_cond := stmt.IsNull("deleted")

	err = worker.dialect.QueryRows(globvar.TableName(), globvar.ColumnNames(), globvar_cond, nil, nil)(
		context.Background(), worker)(
		func(scan excute.Scanner, _ int) error {
			err := globvar.Scan(scan)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}

			records = append(records, globvar)

			return err
		})
	if err != nil {
		return
	}

	count := 0
	push, pop := macro.StringBuilder()
	for _, key := range KeyNames() {
		var found bool = false
	LOOP:
		for i := range records {
			if key == records[i].Name {
				found = true
				break LOOP
			}
		}
		if !found {
			count++
			push(key)
		}
	}
	if 0 < count {
		return errors.Errorf("not exists global_variables keys=['%s']", pop("', '"))
	}

	return nil
}

func (worker *GlobalVariableUpdate) Merge() (err error) {
	records := make([]globvarv2.GlobalVariables, 0, state.ENV__INIT_SLICE_CAPACITY__())

	globvar := globvarv2.GlobalVariables{}
	globvar.Updated = *vanilla.NewNullTime(worker.offset)

	globvar_cond := stmt.IsNull("deleted")

	err = worker.dialect.QueryRows(globvar.TableName(), globvar.ColumnNames(), globvar_cond, nil, nil)(
		context.Background(), worker)(
		func(scan excute.Scanner, _ int) error {
			err := globvar.Scan(scan)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}

			records = append(records, globvar)

			return err
		})
	if err != nil {
		return
	}

	for _, key := range KeyNames() {
		var found bool = false
	LOOP:
		for i := range records {
			if key == records[i].Name {
				found = true
				break LOOP
			}
		}
		if !found {
			globvar_key, err := ParseKey(key)
			if err != nil {
				return errors.Wrapf(err, "ParseGlobVar%v",
					logs.KVL(
						"key", key,
					))
			}

			globvar, updated_columns, ok := GetDefaultGlobalVariable(globvar_key, time.Now())
			if !ok {
				return errors.Errorf("not found global_variables%v",
					logs.KVL(
						"key", key,
					))
			}

			_, _, err = worker.dialect.InsertOrUpdate(globvar.TableName(), globvar.ColumnNames(), updated_columns, globvar.Values())(
				context.Background(), worker)
			if err != nil {
				return errors.Wrapf(err, "failed to create or update global_variables")
			}
		}
	}

	return nil
}

// func foreach_environment(elems []envv1.Environment, fn func(elem envv1.Environment) bool) {
// 	for n := range elems {
// 		ok := fn(elems[n])
// 		if !ok {
// 			return
// 		}
// 	}
// }
