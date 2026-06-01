package rapididentity

import (
	"bytes"
	"context"
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
	Query AuditReportQuery `json:"query" jsonschema:"The query to run This is a required member"`

	// The maximum number of records to return per page
	PageSize int `json:"pageSize" jsonschema:"The maximum number of records to return per page"`

	// The token for the next page of results,
	// retrieved from RunAuditReportOutput.NextPageToken
	PageToken string `json:"pageToken" jsonschema:"The token for the next page of results, retrieved from RunAuditReportOutput.NextPageToken"`
}

type AuditReportQueryList []AuditReportQuery

func (arql AuditReportQueryList) MarshalJSON() ([]byte, error) {
	if arql == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AuditReportQuery(arql))
}

type AuditReportFieldValueList []AuditReportFieldValue

func (arfvl AuditReportFieldValueList) MarshalJSON() ([]byte, error) {
	if arfvl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AuditReportFieldValue(arfvl))
}

// Audt report query component
type AuditReportQuery struct {
	// A grouping of query operations
	ChildNodes AuditReportQueryList `json:"childNodes" jsonschema:"A grouping of query operations"`

	// Query comment
	Comment string `json:"comment" jsonschema:"Query comment"`

	// The date format for the fieldValue
	CustomDateFormat string `json:"customDateFormat" jsonschema:"The date format for the fieldValue"`

	// The field name to filter on. For a list of
	// acceptable field names, see
	// reporting/queryBuilderColumns?type=AUDIT endpoint
	FieldName string `json:"fieldName" jsonschema:"The field name to filter on. For a list of acceptable field names, see reporting/queryBuilderColumns?type=AUDIT endpoint"`

	// The secondary field name to filter on.
	FieldSecondaryName string `json:"fieldSecondaryName" jsonschema:"The secondary field name to filter on."`

	// The field value to filter on. In some
	// cases field values are specific, and require
	// knowing a list of values. One such example is action.displayName
	// which requires a value retrieved from the
	// reporting/possibleValues?type=AUDIT&key=action.displayName
	// endpoint
	FieldValue string `json:"fieldValue" jsonschema:"The field value to filter on. In some cases field values are specific, and require knowing a list of values. One such example is action.displayName which requires a value retrieved from the reporting/possibleValues?type=AUDIT&key=action.displayName endpoint"`

	// The field values to filter on
	FieldValues AuditReportFieldValueList `json:"fieldValues" jsonschema:"The field values to filter on"`

	// The id of the AuditReportQuery. This is
	// utilized for the ParentNode relationship
	Id string `json:"id,omitempty" jsonschema:"The id of the AuditReportQuery. This is utilized for the ParentNode relationship"`

	// The operator for the query.
	// This member is required
	OperatorType AuditReportOperator `json:"operatorType" jsonschema:"The operator for the query. This member is required"`

	// The parent node of the operation
	ParentNode string `json:"parentNode" jsonschema:"The parent node of the operation"`
}
// Field value for relative days and users
type AuditReportFieldValue struct {
	// The dn of the user you are referencing.
	// This is also used for relative days,
	// such as LAST_7_DAYS.
	Dn string `json:"dn" jsonschema:"The dn of the user you are referencing. This is also used for relative days, such as LAST_7_DAYS."`

	// Possible values are Person or
	// a relative day, such as LAST_7_DAYS
	FieldNameAndServerId string `json:"fieldNameAndServerId" jsonschema:"Possible values are Person or a relative day, such as LAST_7_DAYS"`

	// The idautoID of the user or
	// a realtive day, such as LAST_7_DAYS
	Id string `json:"id" jsonschema:"The idautoID of the user or a realtive day, such as LAST_7_DAYS"`

	// The display name of the user or
	// a relative day, such as LAST_7_DAYS
	Name string `json:"name" jsonschema:"The display name of the user or a relative day, such as LAST_7_DAYS"`
}

type AuditReportResultList []AuditReportResult

func (arrl AuditReportResultList) MarshalJSON() ([]byte, error) {
	if arrl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AuditReportResult(arrl))
}

// The result provided by the RapidIdentity
// server when running an audit report query
type RunAuditReportOutput struct {
	// List of audit records
	AuditLogRecords AuditReportResultList `json:"auditLogRecords" jsonschema:"List of audit records"`

	// Whether the limit has been reached
	// for the number of results
	AdminLimitEnforced bool `json:"adminLimitEnforced" jsonschema:"Whether the limit has been reached for the number of results"`

	// Token for retrieving the next page of results.
	// Pass this value as RunAuditReportInput.PageToken in the next call.
	// Empty when there are no more pages.
	NextPageToken string `json:"nextPageToken" jsonschema:"Token for retrieving the next page of results. Pass this value as RunAuditReportInput.PageToken in the next call. Empty when there are no more pages."`
}

type AuditReportExtendedPropertiesList []AuditReportExtendedProperties

func (arepl AuditReportExtendedPropertiesList) MarshalJSON() ([]byte, error) {
	if arepl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AuditReportExtendedProperties(arepl))
}

// Result for an audit report query
type AuditReportResult struct {
	// Unique ID for the audit result record
	Id string `json:"id" jsonschema:"Unique ID for the audit result record"`

	// The RapidIdentity component an
	// audit event corresponds
	Product AuditReportBaseDetail `json:"product" jsonschema:"The RapidIdentity component an audit event corresponds"`

	// The RapidIdentity module an
	// audit event corresponds
	Module AuditReportBaseDetail `json:"module" jsonschema:"The RapidIdentity module an audit event corresponds"`

	// The specific action for the audit event
	Action AuditReportActionDetail `json:"action" jsonschema:"The specific action for the audit event"`

	// The time the audit event occurred
	Timestamp string `json:"timestamp" jsonschema:"The time the audit event occurred"`

	// The host IP Address that the
	// audit event came.
	HostIp string `json:"hostIp" jsonschema:"The host IP Address that the audit event came."`

	// The idautoID of the entity that committed the action
	PerpetratorId string `json:"perpetratorId" jsonschema:"The idautoID of the entity that committed the action"`

	// The distinguished name of the entity that committed the action
	PerpetratorDn string `json:"perpetratorDn" jsonschema:"The distinguished name of the entity that committed the action"`

	// The IP Address of the entity that
	// committed the action. If a proxy is used
	// this will be the IP Address of the proxy
	PerpetratorIp string `json:"perpetratorIp" jsonschema:"The IP Address of the entity that committed the action. If a proxy is used this will be the IP Address of the proxy"`

	// The IP Address of the entity that
	// committed the action. This will be the
	// IP address of the originating client
	PerpetratorIpForwarded string `json:"perpetratorIpForwarded" jsonschema:"The IP Address of the entity that committed the action. This will be the IP address of the originating client"`

	// The target system the committed action
	// was invoked
	TargetSystem string `json:"targetSystem" jsonschema:"The target system the committed action was invoked"`

	// The idautoID of the target that
	// the committed action was invoked on
	TargetId string `json:"targetId" jsonschema:"The idautoID of the target that the committed action was invoked on"`

	// A flexible field that typically holds the
	// friendly name of the target the action was
	// invoked on. In some cases
	// this can be the dn.
	Target string `json:"target" jsonschema:"A flexible field that typically holds the friendly name of the target the action was invoked on. In some cases this can be the dn."`

	// Whether the audit action was successful
	Successful bool `json:"successful" jsonschema:"Whether the audit action was successful"`

	// Whether the audit action was synced
	// to other systems
	Synced bool `json:"synced" jsonschema:"Whether the audit action was synced to other systems"`

	// Additional properties for the audit event
	ExtendedProperties AuditReportExtendedPropertiesList `json:"extendedProperties" jsonschema:"Additional properties for the audit event"`
}

type AuditReportBaseDetailList []AuditReportBaseDetail

func (arbdl AuditReportBaseDetailList) MarshalJSON() ([]byte, error) {
	if arbdl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AuditReportBaseDetail(arbdl))
}

// Base details for several components
type AuditReportBaseDetail struct {
	// Unique ID
	Id string `json:"id" jsonschema:"Unique ID"`

	// Friendly name
	DisplayName string `json:"displayName" jsonschema:"Friendly name"`
}

// Details for audit actions
type AuditReportActionDetail struct {
	AuditReportBaseDetail
	// Classification group of the action.
	Classification AuditReportBaseDetail `json:"classification" jsonschema:"Classification group of the action."`

	// Categories the action is included in.
	Categories AuditReportBaseDetailList `json:"categories" jsonschema:"Categories the action is included in."`
}

// Additional Properties syntax
type AuditReportExtendedProperties struct {
	// The field name of the additional property
	Key string `json:"key" jsonschema:"The field name of the additional property"`

	// The value of the additional property
	Value string `json:"value" jsonschema:"The value of the additional property"`
}

// Runs an audit report query.
//
//meta:operation POST /reporting/auditQuery
func (c *Client) RunAuditReport(ctx context.Context, params RunAuditReportInput) (*RunAuditReportOutput, error) {
	url := fmt.Sprintf("%s/reporting/auditQuery", c.baseEndpoint)
	sep := "?"
	if params.PageSize > 0 {
		url += fmt.Sprintf("%spage_size=%d", sep, params.PageSize)
		sep = "&"
	}
	if params.PageToken != "" {
		url += fmt.Sprintf("%spage_token=%s", sep, params.PageToken)
	}
	query, err := json.Marshal(params.Query)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewBuffer(query)
	req, err := c.GenerateRequest(ctx, "POST", url, requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
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
