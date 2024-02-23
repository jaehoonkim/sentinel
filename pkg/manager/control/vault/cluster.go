package vault

import (
	"context"

	"github.com/jaehoonkim/morpheus/pkg/manager/database"
	"github.com/jaehoonkim/morpheus/pkg/manager/database/vanilla/excute"
	"github.com/jaehoonkim/morpheus/pkg/manager/database/vanilla/stmt"
	"github.com/jaehoonkim/morpheus/pkg/manager/model/auths/v2"
	clusters "github.com/jaehoonkim/morpheus/pkg/manager/model/cluster/v3"
	"github.com/pkg/errors"
)

// import (
// 	"github.com/jaehoonkim/morpheus/pkg/manager/database"
// 	"github.com/jaehoonkim/morpheus/pkg/manager/database/prepare"
// 	"github.com/jaehoonkim/morpheus/pkg/manager/macro/logs"
// 	clusterv1 "github.com/jaehoonkim/morpheus/pkg/manager/model/cluster/v1"
// 	"github.com/pkg/errors"
// 	"xorm.io/xorm"
// )

// type Cluster struct {
// 	tx *xorm.Session
// }

// func NewCluster(tx *xorm.Session) *Cluster {
// 	return &Cluster{tx: tx}
// }

// func (vault Cluster) Create(model clusterv1.Cluster) (*clusterv1.Cluster, error) {
// 	if err := database.XormCreate(vault.tx, &model); err != nil {
// 		return nil, errors.Wrapf(err, "create %v", model.TableName())
// 	}

// 	return &model, nil
// }

// func (vault Cluster) Get(uuid string) (*clusterv1.Cluster, error) {
// 	where := "uuid = ?"
// 	args := []interface{}{
// 		uuid,
// 	}
// 	model := &clusterv1.Cluster{}
// 	if err := database.XormGet(
// 		vault.tx.Where(where, args...), model); err != nil {
// 		return nil, errors.Wrapf(err, "get %v", model.TableName())
// 	}

// 	return model, nil
// }

// func (vault Cluster) Find(where string, args ...interface{}) ([]clusterv1.Cluster, error) {
// 	models := make([]clusterv1.Cluster, 0)
// 	if err := database.XormFind(
// 		vault.tx.Where(where, args...), &models); err != nil {
// 		return nil, errors.Wrapf(err, "find %v", new(clusterv1.Cluster).TableName())
// 	}

// 	return models, nil
// }

// func (vault Cluster) Query(query map[string]string) ([]clusterv1.Cluster, error) {
// 	//parse query
// 	preparer, err := prepare.NewParser(query)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "query %v%v", new(clusterv1.Cluster).TableName(),
// 			logs.KVL(
// 				"query", query,
// 			))
// 	}

// 	//find service
// 	models := make([]clusterv1.Cluster, 0)
// 	if err := database.XormFind(
// 		preparer.Prepared(vault.tx), &models); err != nil {
// 		return nil, errors.Wrapf(err, "query %v%v", new(clusterv1.Cluster).TableName(),
// 			logs.KVL(
// 				"query", query,
// 			))
// 	}

// 	return models, nil
// }

// func (vault Cluster) Update(model clusterv1.Cluster) (*clusterv1.Cluster, error) {
// 	where := "uuid = ?"
// 	args := []interface{}{
// 		model.Uuid,
// 	}

// 	if err := database.XormUpdate(
// 		vault.tx.Where(where, args...), &model); err != nil {
// 		return nil, errors.Wrapf(err, "update %v", model.TableName())
// 	}

// 	return &model, nil
// }

// func (vault Cluster) Delete(uuid string) error {
// 	where := "uuid = ?"
// 	args := []interface{}{
// 		uuid,
// 	}
// 	model := &clusterv1.Cluster{}
// 	if err := database.XormDelete(
// 		vault.tx.Where(where, args...), model); err != nil {
// 		return errors.Wrapf(err, "delete %v", model.TableName())
// 	}

// 	return nil
// }

func CheckCluster(ctx context.Context, tx excute.Preparer, dialect excute.SqlExcutor,
	claims *auths.TenantAccessTokenClaims,
	cluster_uuid string,
) error {
	// check cluster
	cluster_table := clusters.TableNameWithTenant(claims.Hash)
	cluster_cond := stmt.And(
		stmt.Equal("uuid", cluster_uuid),
		stmt.IsNull("deleted"),
	)
	cluster_exist, err := dialect.Exist(cluster_table, cluster_cond)(ctx, tx)
	if err != nil {
		return errors.Wrapf(err, "check cluster")
	}
	if !cluster_exist {
		return errors.Wrapf(database.ErrorRecordWasNotFound, "check cluster")
	}
	return nil
}
