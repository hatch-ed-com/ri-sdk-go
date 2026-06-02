package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetConnectActions(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "sec_mgr")
		testQueryParam(t, r, "metaDataOnly", "true")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{ 
				"name": "all"
			}`,
		)
	})

	input := GetConnectActionsInput{
		Project:      "sec_mgr",
		MetaDataOnly: true,
	}
	ctx := context.Background()
	output, err := client.GetConnectActions(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Name
	want := "all"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectActionById(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/actions/{nameOrId}", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "metaDataOnly", "true")
		nameOrId := r.PathValue("nameOrId")
		if nameOrId == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "a name or id is required")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`{ 
				"id": "%s"
			}`,
			nameOrId,
		)
	})

	input := GetConnectActionByIdInput{
		Id:           "1234",
		MetaDataOnly: true,
	}
	ctx := context.Background()
	output, err := client.GetConnectActionById(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Action.Id
	want := input.Id

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectFiles(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/files/{filePath...}", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "sec_mgr")
		filePath := r.PathValue("filePath")
		if filePath == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "a file path is required")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`{ 
				"path": "%s"
			}`,
			filePath,
		)
	})

	input := GetConnectFilesInput{
		Path:    "log/jobs",
		Project: "sec_mgr",
	}
	ctx := context.Background()
	output, err := client.GetConnectFiles(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Path
	want := input.Path

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectFileContentZip(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/fileContentZip", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testHeader(t, r, "Accept", "application/zip")
		testQueryParam(t, r, "project", "sec_mgr")
		testQueryParam(t, r, "path", "/hello/world.txt")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello World")
	})

	input := GetConnectFileContentZipInput{
		PathList: []string{"/hello/world.txt", "/foo/bar.txt"},
		Project:  "sec_mgr",
	}
	ctx := context.Background()
	output, err := client.GetConnectFileContentZip(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output
	want := "Hello World"
	if string(got) != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectFileContent(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/fileContent/{filePath...}", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testHeader(t, r, "Accept", "text/plain")
		testQueryParam(t, r, "project", "sec_mgr")
		testQueryParam(t, r, "decompress", "true")
		filePath := r.PathValue("filePath")
		if filePath == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "a file path is required")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello World")
	})

	input := GetConnectFileContentInput{
		Path:         "/hello/world.txt",
		Project:      "sec_mgr",
		Decompress:   true,
		ResponseType: "text/plain",
	}
	ctx := context.Background()
	output, err := client.GetConnectFileContent(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output
	want := "Hello World"
	if string(got) != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectJobs(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{
				"jobs": [
					{ 
						"id": "1234"
					}
				]
			}`,
		)
	})

	input := GetConnectJobsInput{
		Project: MainProject,
	}
	ctx := context.Background()
	output, err := client.GetConnectJobs(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Jobs[0].Id
	want := "1234"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectProjects(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{
				"projects": [
					{ 
						"id": "1234"
					}
				]
			}`,
		)
	})

	ctx := context.Background()
	output, err := client.GetConnectProjects(ctx)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Projects[0].Id
	want := "1234"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestSearchConnectActionSets(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/search/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "sec_mgr")
		testQueryParam(t, r, "searchString", "FnAuthenticate")
		testQueryParam(t, r, "matchAction", "false")
		testQueryParam(t, r, "matchCase", "false")
		testQueryParam(t, r, "regex", "false")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{ 
				"name": "result"
			}`,
		)
	})

	input := SearchConnectActionSetsInput{
		SearchString: "FnAuthenticate",
		Project:      "sec_mgr",
	}
	ctx := context.Background()
	output, err := client.SearchConnectActionSets(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Name
	want := "result"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestConnect_MarshalJSON_ZeroValue(t *testing.T) {

	t.Parallel()

	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "GetConnectActionsOutput with nil ActionDefs",
			input: GetConnectActionsOutput{
				Name:       "test-actions",
				ActionDefs: nil,
			},
			mustContain: `"actionDefs":[]`,
		},
		{
			name: "GetConnectFileContentZipInput with nil PathList",
			input: GetConnectFileContentZipInput{
				PathList: nil,
			},
			mustContain: `"pathList":[]`,
		},
		{
			name: "GetConnectFilesOutput with nil FileEntries",
			input: GetConnectFilesOutput{
				FileEntry: FileEntry{
					Path: "test-path",
				},
				FileEntries: nil,
			},
			mustContain: `"fileEntries":[]`,
		},
		{
			name: "GetConnectJobsOutput with nil Jobs",
			input: GetConnectJobsOutput{
				Jobs: nil,
			},
			mustContain: `"jobs":[]`,
		},
		{
			name: "GetConnectProjectsOutput with nil Projects",
			input: GetConnectProjectsOutput{
				Projects: nil,
			},
			mustContain: `"projects":[]`,
		},
		{
			name: "ConnectProject with nil RestPoints",
			input: ConnectProject{
				Id: "proj-1",
				RestPoints: RestPointConfig{
					RestPoints: nil,
				},
			},
			mustContain: `"restPoints":[]`,
		},
		{
			name: "RestPoint with nil ArgMap",
			input: RestPoint{
				Id:     "rp-1",
				ArgMap: nil,
			},
			mustContain: `"argMap":[]`,
		},
		{
			name: "SearchConnectActionSetsOutput with nil ActionDefs",
			input: SearchConnectActionSetsOutput{
				Name:       "test-query",
				ActionDefs: nil,
			},
			mustContain: `"actionDefs":[]`,
		},
		{
			name: "ActionDef with nil ArgDefs and Actions",
			input: ActionDef{
				Id:      "action-1",
				ArgDefs: nil,
				Actions: nil,
			},
			mustContain: `"argDefs":[]`,
		},
		{
			name: "ConnectAction with nil Args",
			input: ConnectAction{
				Id:   "connect-1",
				Args: nil,
			},
			mustContain: `"args":[]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshaledBytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal struct: %v", err)
			}

			result := string(marshaledBytes)

			if !strings.Contains(result, tt.mustContain) {
				t.Errorf("expected JSON to contain %q, but got: %s", tt.mustContain, result)
			}

			if strings.Contains(result, ":null") {
				t.Errorf("detected unexpected 'null' value in marshaled output: %s", result)
			}
		})
	}
}

func TestSaveConnectAction(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		var action ActionDef
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		err = json.Unmarshal(reqBody, &action)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`{
				"id": "%s"
			}`,
			action.Id)
	})

	input := SaveConnectActionInput{
		Action: ActionDef{
			Id:           "A5500922-4C5D-43B8-B407-93B42DEA96E2",
			Version:      0,
			Name:         "TestRISDK",
			Description:  "This is a test from RI SDK",
			ReturnsValue: false,
			Actions: ConnectActionList{
				{
					Args: ArgDefList{
						{
							Description: "The message to be logged",
							Name:        "message",
							Type:        "string",
							Value:       "\"This is a test from the RI SDK\"",
						},
						{
							Description: "The log level (default: INFO)",
							Name:        "level",
							Type:        "string",
							Value:       "\"WARN\"",
						},
					},
					Id:   "7F3F77A5-737E-4036-A75D-8DD39A336ED1",
					Name: "log",
				},
			},
		},
	}

	ctx := context.Background()
	output, err := client.SaveConnectAction(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Action.Id
	want := input.Action.Id

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestDeleteConnectActionById(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/actions/{nameOrId}", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		nameOrId := r.PathValue("nameOrId")
		if nameOrId == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "a name or id is required")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{ 
				"success": true
			}`,
		)
	})

	input := DeleteConnectActionByIdInput{
		Id: "1234",
	}
	ctx := context.Background()
	output, err := client.DeleteConnectActionById(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.DeleteOperationStatus.Success
	want := true

	if got != want {
		t.Errorf("got %t. want %t", got, want)
	}
}
