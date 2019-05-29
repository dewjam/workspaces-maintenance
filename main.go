package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

type ScheduledEvent struct {
	Tag string `json:"tag"`
}

func HandleRequest(ctx context.Context, tag ScheduledEvent) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := workspaces.New(sess)

	workspaceOutput, err := svc.DescribeWorkspaces(&workspaces.DescribeWorkspacesInput{})
	if err != nil {
		return "An error occurred", err
	}

	stoppedWorkspaces := getStoppedWorkspaces(workspaceOutput.Workspaces)

	for w := range stoppedWorkspaces {
		svc.StartWorkspaces(&workspaces.StartWorkspacesInput{StartWorkspaceRequests: stoppedWorkspaces[w]})
	}

	return fmt.Sprintf("started workspaces: %s", stoppedWorkspaces), nil
}

// getStoppedWorkspaces loops through the list of workspaces and gets those
// in a stopped state
func getStoppedWorkspaces(w []*workspaces.Workspace) [][]*workspaces.StartRequest {
	var stoppedWorkspaces []*workspaces.StartRequest
	for i := range w {
		if *w[i].State == "STOPPED" {
			stoppedWorkspaces = append(stoppedWorkspaces, &workspaces.StartRequest{WorkspaceId: w[i].WorkspaceId})
			fmt.Printf("Starting workspace: %s", *w[i].WorkspaceId)
		}
	}
	return chunkStartRequest(stoppedWorkspaces)
}

// chunkStartRequest takes the list of workspaces and breaks them up into
// groups of 25.  This is the max number of workspaces that can be started at
// one time.
func chunkStartRequest(w []*workspaces.StartRequest) [][]*workspaces.StartRequest {
	var chunkedStartRequest [][]*workspaces.StartRequest
	var i int
	for i = len(w); i > 25; i -= 25 {
		chunkedStartRequest = append(chunkedStartRequest, w[i-25:i])
	}
	chunkedStartRequest = append(chunkedStartRequest, w[:i])
	return chunkedStartRequest
}

func main() {
	lambda.Start(HandleRequest)
}
