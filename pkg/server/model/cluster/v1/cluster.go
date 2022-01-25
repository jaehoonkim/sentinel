package v1

import (
	metav1 "github.com/NexClipper/sudory/pkg/server/model/meta/v1"
	"github.com/NexClipper/sudory/pkg/server/model/orm"
)

//Cluster Property
type ClusterProperty struct {
}

//Cluster
type Cluster struct {
	metav1.LabelMeta `json:",inline" xorm:"extends"` //inline labelmeta
	ClusterProperty  `json:",inline" xorm:"extends"` //inline property
}

//DATABASE SCHEMA: Cluster
type DbSchemaCluster struct {
	metav1.DbMeta `xorm:"extends"`
	Cluster       `xorm:"extends"`
}

var _ orm.TableName = (*DbSchemaCluster)(nil)

func (DbSchemaCluster) TableName() string {
	return "cluster"
}

//HTTP REQUEST BODY: Cluster
type HttpReqCluster struct {
	Cluster `json:",inline"`
}

//HTTP RESPONSE BODY: Cluster
type HttpRspCluster struct {
	Cluster `json:",inline"`
}

//변환 DbSchema -> Cluster
func TransFormDbSchema(s []DbSchemaCluster) []Cluster {
	var out = make([]Cluster, len(s))
	for n, it := range s {
		out[n] = it.Cluster
	}
	return out
}

//변환 Cluster -> HttpRsp
func TransToHttpRsp(s []Cluster) []HttpRspCluster {
	var out = make([]HttpRspCluster, len(s))
	for n, it := range s {
		out[n].Cluster = it
	}
	return out
}