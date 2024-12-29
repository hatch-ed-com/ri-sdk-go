package rapididentity

import (
	"encoding/json"
	"fmt"
)

type GetDelegationsForUserInput struct {
	UserId string
}

type GetDelegationsForUserOutput struct {
	AggregatedDelegation AggregatedDelegation `json:"aggregatedDelegation"`
}

type AggregatedDelegation struct {
	Id                 string              `json:"id"`
	User               User                `json:"user"`
	HelpdeskQuestions  []string            `json:"helpdeskQuestions"`
	DelegationProfiles []DelegationProfile `json:"delegationProfiles"`
}

type User struct {
	Id             string   `json:"id"`
	Dn             string   `json:"dn"`
	Username       string   `json:"username"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Email          string   `json:"email"`
	Distinguisher  string   `json:"distinguisher"`
	ImageUrl       string   `json:"imageUrl"`
	MobileNumbers  []string `json:"mobileNumbers"`
	AlternateEmail string   `json:"alternateEmail"`
}

type DelegationProfile struct {
	Delegation Delegation `json:"delegation"`
	Profile    Profile    `json:"profile"`
}

type Delegation struct {
	Id           string                `json:"id"`
	Name         string                `json:"name"`
	Type         string                `json:"type"`
	SourceBaseDN string                `json:"sourceBaseDN"`
	TargetBaseDN string                `json:"targetBaseDN"`
	Attributes   []DelegationAttribute `json:"attributes"`
	LayoutImage  string                `json:"layoutImage"`
	Layout1      string                `json:"layout1"`
	Layout2      string                `json:"layout2"`
	Layout3      string                `json:"layout3"`
	Actions      []Action              `json:"actions"`
}

type Profile struct {
	Id         string             `json:"id"`
	Attributes []ProfileAttribute `json:"attributes"`
}

type DelegationAttribute struct {
	GalItem       GalItem `json:"galItem"`
	Name          string  `json:"name"`
	Editable      bool    `json:"editable"`
	ShowInList    bool    `json:"showInList"`
	ShowInDetails bool    `json:"showInDetails"`
	Required      bool    `json:"required"`
}

type GalItem struct {
	Id              string `json:"id"`
	FriendlyName    string `json:"friendlyName"`
	AllowMultiValue bool   `json:"allowMultiValue"`
	Type            string `json:"type"`
}

type Action struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProfileAttribute struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

func (c *Client) GetDelegationsForUser(input GetDelegationsForUserInput) (*GetDelegationsForUserOutput, error) {
	var output GetDelegationsForUserOutput

	url := fmt.Sprintf("%s/profiles/aggregated/for/%s", c.baseEndpoint, input.UserId)
	req, err := c.GenerateRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := c.ReceiveResponse(res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resBody, &output)
	if err != nil {
		return nil, RapidIdentityError{
			Method:  req.Method,
			ReqUrl:  req.URL,
			Message: string(resBody),
			Reason:  err.Error(),
			Code:    res.StatusCode,
		}
	}

	return &output, nil
}
