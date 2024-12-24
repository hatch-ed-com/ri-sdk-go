package rapididentity

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AuditReportOperator string

// Available operators for an audit report query
const (
	EQUAL        AuditReportOperator = "eq"
	NOT_EQUAL    AuditReportOperator = "ne"
	LESS_THAN    AuditReportOperator = "lt"
	GREATER_THAN AuditReportOperator = "gt"
	LIKE         AuditReportOperator = "like"
	AND          AuditReportOperator = "AND"
	OR           AuditReportOperator = "OR"
)

// Input for running an audit report
// query in the RapidIdentity Reporting
// module.
type RunAuditReportInput struct {
	// The query to run
	// This is a required member
	Query AuditReportQuery
}

// Audt report query component
type AuditReportQuery struct {
	// A grouping of query operations
	ChildNodes []AuditReportQuery `json:"childNodes"`

	// Query comment
	Comment string `json:"comment"`

	// The date format for the fieldValue
	CustomDateFormat string `json:"customDateFormat"`

	// The field name to filter on. For a list of
	// acceptable field names, see
	// reporting/queryBuilderColumns?type=AUDIT endpoint
	FieldName string `json:"fieldName"`

	// The secondary field name to filter on.
	FieldSecondaryName string `json:"fieldSecondaryName"`

	// The field value to filter on. In some
	// cases field values are specific, and require
	// knowing a list of values. One such example is action.displayName
	// which requires a value retrieved from the
	// reporting/possibleValues?type=AUDIT&key=action.displayName
	// endpoint
	FieldValue string `json:"fieldValue"`

	// The field values to filter on
	FieldValues []AuditReportFieldValue `json:"fieldValues"`

	// The id of the AuditReportQuery. This is
	// utilized for the ParentNode relationship
	Id string `json:"id,omitempty"`

	// The operator for the query.
	// This member is required
	OperatorType AuditReportOperator `json:"operatorType"`

	// The parent node of the operation
	ParentNode string `json:"parentNode"`
}

// Field value for relative days and users
type AuditReportFieldValue struct {
	// The dn of the user you are referencing.
	// This is also used for relative days,
	// such as LAST_7_DAYS.
	Dn string `json:"dn"`

	// Possible values are Person or
	// a relative day, such as LAST_7_DAYS
	FieldNameAndServerId string `json:"fieldNameAndServerId"`

	// The idautoID of the user or
	// a realtive day, such as LAST_7_DAYS
	Id string `json:"id"`

	// The display name of the user or
	// a relative day, such as LAST_7_DAYS
	Name string `json:"name"`
}

// The result provided by the RapidIdentity
// server when running an audit report query
type RunAuditReportOutput struct {
	// List of audit records
	AuditLogRecords []AuditReportResult `json:"auditLogRecords"`

	// Whether the limit has been reached
	// for the number of results
	AdminLimitEnforced bool `json:"adminLimitEnforced"`
}

// Result for an audit report query
type AuditReportResult struct {
	// Unique ID for the audit result record
	Id string `json:"id"`

	// The RapidIdentity component an
	// audit event corresponds
	Product AuditReportBaseDetail `json:"product"`

	// The RapidIdentity module an
	// audit event corresponds
	Module AuditReportBaseDetail `json:"module"`

	// The specific action for the audit event
	Action AuditReportActionDetail `json:"action"`

	// The time the audit event occurred
	Timestamp string `json:"timestamp"`

	// The host IP Address that the
	// audit event came.
	HostIp string `json:"hostIp"`

	// The idautoID of the entity that committed the action
	PerpetratorId string `json:"perpetratorId"`

	// The distinguished name of the entity that committed the action
	PerpetratorDn string `json:"perpetratorDn"`

	// The IP Address of the entity that
	// committed the action. If a proxy is used
	// this will be the IP Address of the proxy
	PerpetratorIp string `json:"perpetratorIp"`

	// The IP Address of the entity that
	// committed the action. This will be the
	// IP address of the originating client
	PerpetratorIpForwarded string `json:"perpetratorIpForwarded"`

	// The target system the committed action
	// was invoked
	TargetSystem string `json:"targetSystem"`

	// The idautoID of the target that
	// the committed action was invoked on
	TargetId string `json:"targetId"`

	// A flexible field that typically holds the
	// friendly name of the target the action was
	// invoked on. In some cases
	// this can be the dn.
	Target string `json:"target"`

	// Whether the audit action was successful
	Successful bool `json:"successful"`

	// Whether the audit action was synced
	// to other systems
	Synced bool `json:"synced"`

	// Additional properties for the audit event
	ExtendedProperties []AuditReportExtendedProperties `json:"extendedProperties"`
}

// Base details for several components
type AuditReportBaseDetail struct {
	// Unique ID
	Id string `json:"id"`

	// Friendly name
	DisplayName string `json:"displayName"`
}

// Details for audit actions
type AuditReportActionDetail struct {
	AuditReportBaseDetail
	// Classification group of the action.
	Classification AuditReportBaseDetail `json:"classification"`

	// Categories the action is included in.
	Categories []AuditReportBaseDetail `json:"categories"`
}

// Additional Properties syntax
type AuditReportExtendedProperties struct {
	// The field name of the additional property
	Key string `json:"key"`

	// The value of the additional property
	Value string `json:"value"`
}

// Runs an audit report query.
//
//meta:operation POST /reporting/auditQuery
func (c *Client) RunAuditReport(params RunAuditReportInput) (*RunAuditReportOutput, error) {
	url := fmt.Sprintf("%s/reporting/auditQuery", c.baseEndpoint)
	query, err := json.Marshal(params.Query)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewBuffer(query)
	req, err := c.GenerateRequest("POST", url, requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.options.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := c.ReceiveResponse(res)
	if err != nil {
		return nil, err
	}
	var output RunAuditReportOutput
	err = json.Unmarshal(resBody, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
