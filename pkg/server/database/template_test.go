package database_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/NexClipper/sudory/pkg/server/database"
	. "github.com/NexClipper/sudory/pkg/server/macro"
	"github.com/NexClipper/sudory/pkg/server/macro/newist"
	templatev1 "github.com/NexClipper/sudory/pkg/server/model/template/v1"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

//Test Template CRUD
//테스트 목적: Template 테이블 기능
//테스트 조건:
//		로컬 호스트 mariadb 인스턴스
//		v1.Template 테이블
func TestTemplateCRUD(t *testing.T) {

	newEngine := func() *xorm.Engine {
		const (
			driver = "mysql"
			dsn    = "sudory:sudory@tcp(127.0.0.1:3306)/sudory?charset=utf8mb4&parseTime=True&loc=Local"
		)

		engine, err := xorm.NewEngine(driver, dsn)
		if ErrorWithHandler(err) {
			panic(err)
		}

		engine.SetLogLevel(log.LOG_DEBUG)

		return engine
	}

	ctx := database.NewContext(newEngine())

	Are := func(expect interface{}, actual interface{}, equal, diff func(string)) {

		ex := fmt.Sprint(expect)
		ac := fmt.Sprint(actual)
		if ex == ac {
			if equal != nil {
				equal(fmt.Sprintf("Equal: expect='%v', actual='%v'", expect, actual))
			}
		} else {
			if equal != nil {
				equal(fmt.Sprintf("NotEq: expect='%v', actual='%v'", expect, actual))
			}
		}
	}

	create_model := func() *templatev1.DbSchemaTemplate {
		out := templatev1.DbSchemaTemplate{}

		out.Uuid = "00001111222233334444555566667777"
		out.Name = "(NEW) test-template"
		out.Summary = newist.String("(NEW) test template")
		out.ApiVersion = "v1"
		out.Origin = newist.String("origin")

		return &out
	}
	update_model := func() *templatev1.DbSchemaTemplate {
		out := templatev1.DbSchemaTemplate{}

		out.Uuid = "00001111222233334444555566667777"
		out.Name = "(UPDATED) test-template"
		out.Summary = newist.String("(UPDATED) test template")
		out.ApiVersion = "v1"
		out.Origin = newist.String("changed-origin")

		return &out
	}

	empty_model := func() *templatev1.DbSchemaTemplate {
		return new(templatev1.DbSchemaTemplate)
	}
	create := func() error {
		err := ctx.CreateTemplate(*create_model())
		if err != nil {
			return fmt.Errorf("created error with; %w", err)
		}

		return nil
	} //생성

	read := func() error {
		record, err := ctx.GetTemplate(create_model().Uuid)
		if err != nil {
			return fmt.Errorf("read error with; %w", err)
		}

		j, err := json.MarshalIndent(record, "", " ")
		if err != nil {
			return fmt.Errorf("read error with json; %w", err)
		}
		t.Logf("\nverbose: %s\n", j)

		return nil
	} //조회
	validate := func(model *templatev1.DbSchemaTemplate) func() error {
		return func() error {
			record, err := ctx.GetTemplate(model.Uuid)
			if err != nil {
				return fmt.Errorf("validate error with; %w", err)
			}

			are_equl := (func(msg string))(nil)
			are_diff := func(msg string) { t.Error(err) }
			//valied
			Are(model.Uuid, record.Uuid, are_equl, are_diff)
			Are(model.Name, record.Name, are_equl, are_diff)
			Are(model.Summary, record.Summary, are_equl, are_diff)
			Are(model.ApiVersion, record.ApiVersion, are_equl, are_diff)
			Are(model.Origin, record.Origin, are_equl, are_diff)

			return nil
		}
	} //데이터 정합 확인

	update := func() error {
		err := ctx.UpdateTemplate(*update_model())
		if err != nil {
			return fmt.Errorf("updated error with; %w", err)
		}

		return nil
	} //데이터 갱신
	delete := func() error {
		err := ctx.DeleteTemplate(update_model().Uuid)
		if err != nil {
			return fmt.Errorf("deleted error with; %w", err)
		}

		return nil
	} //삭제
	clear := func() error {
		table := empty_model().TableName()

		sqlrst, err := newEngine().
			Exec(fmt.Sprintf("delete from %s where uuid = ?", table), create_model().Uuid)
		if err != nil {
			return fmt.Errorf("no record clear; %w", err)
		}

		affect, err := sqlrst.RowsAffected()
		if err != nil {
			return fmt.Errorf("clear error with; %w", err)
		}
		if !(0 < affect) {
			return fmt.Errorf("clear error with affect %d;", affect)
		}
		return nil
	} //데이터 정리

	//시나리오 구성
	TestScenarios([]TestChapter{
		{Subject: "create newist record", Action: create},
		{Subject: "vaild created record", Action: validate(create_model())},
		{Subject: "read created record", Action: read},
		{Subject: "update record", Action: update},
		{Subject: "vaild updated record", Action: validate(update_model())},
		{Subject: "read updated record", Action: read},
		{Subject: "delete record", Action: delete},    //테이블에서 지워지는것이 아니라 삭제 플래그가 업데이트 됨 (레코드가 남아있음)
		{Subject: "clear test record", Action: clear}, //남아있는 레코드 정리
	}).Foreach(func(err error) { t.Error(err) })
}
