package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

type testHelper struct {
	name    string
	apiMethod string
	fixture string
	id      string
	out     interface{}
}

func TestGetComplianceProfile(t *testing.T) {
	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	simpleSuccessfulTestNeededValues := testHelper{
		name:    "TestComplianceProfile",
		apiMethod: "GET",
		fixture: "fixtures/complianceProfile.json",
		id:      "bbcccdddd-eeeeef-fffff",
		out: ComplianceProfile{
			ID:                  "abbcccdddd-eeeeef-fffff",
			Name:                "Foo Threat Research",
			SeedId:              "random-string-of-characters-for-seed",
			CustomerId:          "11111111-1111-1111-1111-111111111111",
			Description:         "Test description",
			Custom:              true,
			Priority:            1,
			CreatedAt:           "2022-03-24T15:09:57.539Z",
			CreatedBy:           "",
			UpdatedAt:           "2022-03-24T15:09:57.539Z",
			UpdatedBy:           "",
			Links:     			 nil,
		},
	}

	simpleSuccessfulTestMultiple := testHelper{
		name:    "Simple Successful test with no id but found by name",
		apiMethod: "GET",
		fixture: "fixtures/complianceProfile.json",
		id:      "",
		out: ComplianceProfile{
			ID:                  "compliance-profile-2",
			Name:                "Foo Threat Research 2",
			SeedId:              "random-string-of-characters-for-seed-2",
			CustomerId:          "11111111-1111-1111-1111-111111111112",
			Description:         "Test description 2",
			Custom:              true,
			Priority:            1,
			CreatedAt:           "2022-03-24T15:09:57.539Z",
			CreatedBy:           "",
			UpdatedAt:           "2022-03-24T15:09:57.539Z",
			UpdatedBy:           "",
			Links:     			 nil,
		},
	}

	
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run(simpleSuccessfulTestNeededValues.name, func(t *testing.T) {
		httpmock.RegisterResponder(simpleSuccessfulTestNeededValues.apiMethod, fmt.Sprintf("https://uptycs.foo/public/api/customers/d/complianceProfiles/%v", simpleSuccessfulTestNeededValues.id),
			func(req *http.Request) (*http.Response, error) {
				fixture, err := RespFromFixture(simpleSuccessfulTestNeededValues.fixture)
				if err != nil {
					t.Errorf(err.Error())
				}
				return fixture, err
			},
		)


		complianceProfileResp, err := c.GetComplianceProfile(ComplianceProfile{
			ID: simpleSuccessfulTestNeededValues.id,
		})

		if err != nil {
			t.Errorf(err.Error())
		}

		if !reflect.DeepEqual(complianceProfileResp, simpleSuccessfulTestNeededValues.out) {
			t.Log("Output does not match expected")
			t.Logf("Expected: %v", simpleSuccessfulTestNeededValues.out)
			t.Logf("Actual: %v", complianceProfileResp)
			t.Fail()
		}
	})
	

	
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run(simpleSuccessfulTestMultiple.name, func(t *testing.T) {
		httpmock.RegisterResponder(simpleSuccessfulTestMultiple.apiMethod, fmt.Sprintf("https://uptycs.foo/public/api/customers/d/complianceProfiles/%v", simpleSuccessfulTestMultiple.id),
			func(req *http.Request) (*http.Response, error) {
				fixture, err := RespFromFixture("fixtures/complianceProfile.json")
				if err != nil {
					t.Errorf(err.Error())
				}
				return fixture, err
			},
		)
		httpmock.RegisterResponder(simpleSuccessfulTestMultiple.apiMethod, fmt.Sprintf("https://uptycs.foo/public/api/customers/d/complianceProfiles"),
			func(req *http.Request) (*http.Response, error) {
				fixture, err := RespFromFixture("fixtures/complianceProfiles.json")
				if err != nil {
					t.Errorf(err.Error())
				}
				return fixture, err
			},
		)

		complianceProfileResp, err := c.GetComplianceProfile(ComplianceProfile{
			Name: "Foo Threat Research 2",
		})

		if err != nil {
			t.Errorf(err.Error())
		}

		if !reflect.DeepEqual(complianceProfileResp, simpleSuccessfulTestMultiple.out) {
			t.Log("Output does not match expected")
			t.Logf("Expected: %v", simpleSuccessfulTestMultiple.out)
			t.Logf("Actual: %v", complianceProfileResp)
			t.Fail()
		}
	})
}

func TestGetComplianceProfiles(t *testing.T) {
	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	simpleSuccessfulTestProfiles := testHelper{
		name:    "TestComplianceProfilesGet",
		apiMethod: "GET",
		fixture: "fixtures/complianceProfiles.json",
		id:      "",
		out: ComplianceProfiles{
			Links: []LinkItem{},
			Offset: 1, 
			Limit: 2,
			Decorators: nil,
			Items: []ComplianceProfile {
				ComplianceProfile {
					ID:                  "compliance-profile-1",
					Name:                "Foo Threat Research",
					SeedId:              "random-string-of-characters-for-seed",
					CustomerId:          "11111111-1111-1111-1111-111111111111",
					Description:         "Test description",
					Custom:              true,
					Priority:            1,
					CreatedAt:           "2022-03-24T15:09:57.539Z",
					CreatedBy:           "",
					UpdatedAt:           "2022-03-24T15:09:57.539Z",
					UpdatedBy:           "",
					Links:     			 nil,
				},
				ComplianceProfile {
					ID:                  "compliance-profile-2",
					Name:                "Foo Threat Research 2",
					SeedId:              "random-string-of-characters-for-seed-2",
					CustomerId:          "11111111-1111-1111-1111-111111111112",
					Description:         "Test description 2",
					Custom:              true,
					Priority:            1,
					CreatedAt:           "2022-03-24T15:09:57.539Z",
					CreatedBy:           "",
					UpdatedAt:           "2022-03-24T15:09:57.539Z",
					UpdatedBy:           "",
					Links:     			 nil,
				},
			},
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run(simpleSuccessfulTestProfiles.name, func(t *testing.T) {
		httpmock.RegisterResponder(simpleSuccessfulTestProfiles.apiMethod, fmt.Sprintf("https://uptycs.foo/public/api/customers/d/complianceProfiles"),
			func(req *http.Request) (*http.Response, error) {
				fixture, err := RespFromFixture("fixtures/complianceProfiles.json")
				if err != nil {
					t.Errorf(err.Error())
				}
				return fixture, err
			},
		)

		complianceProfileResp, err := c.GetComplianceProfiles()

		if err != nil {
			t.Errorf(err.Error())
		}

		if !reflect.DeepEqual(complianceProfileResp, simpleSuccessfulTestProfiles.out) {
			t.Log("Output does not match expected")
			t.Logf("Expected: %v", simpleSuccessfulTestProfiles.out)
			t.Logf("Actual: %v", complianceProfileResp)
			t.Fail()
		}
	})
}