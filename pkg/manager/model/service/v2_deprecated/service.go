package v2

import (
	"fmt"
	"strings"
	"time"

	"github.com/jaehoonkim/sentinel/pkg/manager/database/vanilla"
	crypto "github.com/jaehoonkim/sentinel/pkg/manager/model/default_crypto_types/v2"
)

type Service_essential struct {
	Name              string             `column:"name"               json:"name,omitempty"`
	Summary           vanilla.NullString `column:"summary"            json:"summary,omitempty"            swaggertype:"string"`
	ClusterUuid       string             `column:"cluster_uuid"       json:"cluster_uuid,omitempty"`
	TemplateUuid      string             `column:"template_uuid"      json:"template_uuid,omitempty"`
	StepCount         int                `column:"step_count"         json:"step_count,omitempty"`
	SubscribedChannel vanilla.NullString `column:"subscribed_channel" json:"subscribed_channel,omitempty" swaggertype:"string"`
	// OnCompletion      OnCompletion      `column:"on_completion"      json:"on_completion,omitempty"`
}

func (Service_essential) TableName() string {
	return "service"
}

type Service struct {
	Uuid              string    `column:"uuid"    json:"uuid,omitempty"`    //pk
	Created           time.Time `column:"created" json:"created,omitempty"` //pk
	Service_essential `json:",inline"`
}

type ServiceStatus_essential struct {
	AssignedClientUuid string             `column:"assigned_client_uuid" json:"assigned_client_uuid,omitempty"`
	StepPosition       int                `column:"step_position"        json:"step_position,omitempty"`
	Status             StepStatus         `column:"status"               json:"status,omitempty"`
	Message            vanilla.NullString `column:"message"              json:"message,omitempty"              swaggertype:"string"`
}

func (ServiceStatus_essential) TableName() string {
	return "service_status"
}

type ServiceStatus struct {
	Uuid    string    `column:"uuid"    json:"uuid,omitempty"`    //pk
	Created time.Time `column:"created" json:"created,omitempty"` //pk

	ServiceStatus_essential `json:",inline"`
}

type ServiceResults_essential struct {
	ResultSaveType ResultSaveType      `column:"result_type" json:"result_type,omitempty"`
	Result         crypto.CryptoString `column:"result"      json:"result,omitempty"`
}

func (ServiceResults_essential) TableName() string {
	return "service_result"
}

type ServiceResult struct {
	Uuid    string    `column:"uuid"    json:"uuid,omitempty"`    //pk
	Created time.Time `column:"created" json:"created,omitempty"` //pk

	ServiceResults_essential `json:",inline"`
}

type Service_tangled struct {
	Service                  `json:",inline"` //service
	ServiceStatus_essential  `json:",inline"` //status
	ServiceResults_essential `json:",inline"` //result

	Updated vanilla.NullTime `column:"updated" json:"updated,omitempty" swaggertype:"string"` //pk
}

/*
`
SELECT A.uuid, A.created,

	       name, summary, cluster_uuid, template_uuid, step_count, subscribed_channel,
	       B.created AS updated, assigned_client_uuid, step_position, status, message, result_type, result
	  FROM service A
	  LEFT JOIN service_status B
	         ON A.uuid = B.uuid
	        AND B.created = (
	            SELECT MAX(B.created) AS MAX_created
	              FROM service_status B
	             WHERE A.uuid = B.uuid
				)
	  LEFT JOIN service_result C
	         ON A.uuid = C.uuid
	        AND C.created = (
	            SELECT MAX(C.created) AS MAX_created
	              FROM service_result C
	             WHERE A.uuid = C.uuid
				)

`
*/
func (record Service_tangled) TableName() string {
	q := `(
    SELECT %v /**columns**/
      FROM %v A /**service A**/
      LEFT JOIN %v B /**service_status B**/
             ON A.uuid = B.uuid 
            AND B.created = (
                SELECT MAX(B.created) AS MAX_created
                  FROM %v B /**service_status B**/
                 WHERE A.uuid = B.uuid 
                )
      LEFT JOIN %v C /**service_result C**/
             ON A.uuid = C.uuid
            AND C.created = (
                SELECT MAX(C.created) AS MAX_created
                  FROM %v C /**service_result C**/
                 WHERE A.uuid = C.uuid
                )
    ) X`

	columns := []string{
		"A.uuid",
		"A.created",
		"B.created AS updated",
		"name",
		"summary",
		"cluster_uuid",
		"template_uuid",
		"step_count",
		"subscribed_channel",
		fmt.Sprintf("IFNULL(assigned_client_uuid, '%v') AS assigned_client_uuid", ""),
		fmt.Sprintf("IFNULL(step_position, %v) AS step_position", 0),
		fmt.Sprintf("IFNULL(status, %v) AS status", int(StepStatusRegist)),
		"message",
		fmt.Sprintf("IFNULL(result_type, %v) AS result_type", int(ResultSaveTypeNone)),
		"result",
	}
	A := record.Service.TableName()
	B := record.ServiceStatus_essential.TableName()
	C := record.ServiceResults_essential.TableName()
	return fmt.Sprintf(q, strings.Join(columns, ", "), A, B, B, C, C)
}

type Service_status struct {
	Service                 `json:",inline"` //service
	ServiceStatus_essential `json:",inline"` //status

	Updated vanilla.NullTime `column:"updated" json:"updated,omitempty" swaggertype:"string"` //pk
}

/*
`
SELECT A.uuid, A.created,

	       name, summary, cluster_uuid, template_uuid, step_count, subscribed_channel,
	       B.created AS updated, assigned_client_uuid, step_position, status, message
	  FROM service A
	  LEFT JOIN service_status B
	         ON A.uuid = B.uuid
	        AND B.created = (
	            SELECT MAX(B.created) AS MAX_created
	              FROM service_status B
	             WHERE A.uuid = B.uuid
				)

`
*/
func (record Service_status) TableName() string {
	q := `(
    SELECT %v /**columns**/
      FROM %v A /**service A**/
      LEFT JOIN %v B /**service_status B**/
             ON A.uuid = B.uuid 
            AND B.created = (
                SELECT MAX(B.created) AS MAX_created
                  FROM %v B /**service_status B**/
                 WHERE A.uuid = B.uuid 
                )
    ) X`

	columns := []string{
		"A.uuid",
		"A.created",
		"B.created AS updated",
		"name",
		"summary",
		"cluster_uuid",
		"template_uuid",
		"step_count",
		"subscribed_channel",
		fmt.Sprintf("IFNULL(assigned_client_uuid, '%v') AS assigned_client_uuid", ""),
		fmt.Sprintf("IFNULL(step_position, %v) AS step_position", 0),
		fmt.Sprintf("IFNULL(status, %v) AS status", int(StepStatusRegist)),
		"message",
	}
	A := record.Service.TableName()
	B := record.ServiceStatus_essential.TableName()
	return fmt.Sprintf(q, strings.Join(columns, ", "), A, B, B)
}
