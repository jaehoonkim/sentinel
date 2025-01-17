package cluster_token

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jaehoonkim/sentinel/pkg/manager/database/vanilla/ice_cream_maker"
	"github.com/jaehoonkim/sentinel/pkg/manager/macro/slicestrings"
	clusterv3 "github.com/jaehoonkim/sentinel/pkg/manager/model/cluster/v3"
	"github.com/jaehoonkim/sentinel/pkg/manager/model/tenants/v3"
)

var (
	TableNameWithTenant = tableNameWithTenant()
)

func tableNameWithTenant() func(tenant_hash string) string {
	var CT = ClusterToken{}
	var C = clusterv3.Cluster{}
	var TC = tenants.TenantClusters{}
	var T = tenants.Tenant{}

	aliasCT := CT.TableName()
	aliasC := C.TableName()
	aliasTC := TC.TableName()
	aliasT := T.TableName()

	columninfos := ice_cream_maker.ParseColumnTag(reflect.TypeOf(CT), ice_cream_maker.ParseColumnTag_opt{})

	columns := make([]string, 0, len(columninfos))
	for i := range columninfos {
		columns = append(columns, columninfos[i].Name)
	}

	columns = slicestrings.Map(columns, func(s string, i int) string {
		return aliasCT + "." + s
	})

	tables := []string{aliasCT, aliasC, aliasTC, aliasT}

	conds := []string{
		fmt.Sprintf("%v.hash = '%%v'", aliasT),
		fmt.Sprintf("%v.deleted IS NULL", aliasT),
		fmt.Sprintf("%v.tenant_id = %v.id", aliasTC, aliasT),
		fmt.Sprintf("%v.id = %v.cluster_id", aliasC, aliasTC),
		fmt.Sprintf("%v.deleted IS NULL", aliasC),
		fmt.Sprintf("%v.cluster_uuid = %v.uuid", aliasCT, aliasC),
	}

	format := fmt.Sprintf("( SELECT %v FROM %v WHERE %v ) X",
		strings.Join(columns, ", "),
		strings.Join(tables, ", "),
		strings.Join(conds, " AND "),
	)

	return func(tenant_hash string) string {
		return fmt.Sprintf(format, tenant_hash)
	}
}
