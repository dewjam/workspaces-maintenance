package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/stretchr/testify/assert"
)

// var workspaceIds = CreateStartRequest(27)

// func CreateStartRequest(c int) []*workspaces.StartRequest {
// 	var startRequest []*workspaces.StartRequest
// 	for c > 0 {
// 		workspaceId := fmt.Sprintf("ws-%d", c)
// 		startRequest = append(startRequest, &workspaces.StartRequest{WorkspaceId: &workspaceId})
// 		c -= 1
// 	}
// 	workspaceId := fmt.Sprintf("ws-%d", c)
// 	startRequest = append(startRequest, &workspaces.StartRequest{WorkspaceId: &workspaceId})
// 	return startRequest
// }

func TestChunkStartRequest(t *testing.T) {
	assert.Equal(t, case1Expected, chunkStartRequest(case1WorkspaceIds))
	assert.Equal(t, case2Expected, chunkStartRequest(case2WorkspaceIds))
	assert.Equal(t, case3Expected, chunkStartRequest(case3WorkspaceIds))
}

func returnPointer(s string) *string {
	return &s
}

var case1Expected = [][]*workspaces.StartRequest{
	{
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-1")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-2")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-3")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-4")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-5")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-6")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-7")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-8")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-9")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-10")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-11")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-12")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-13")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-14")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-15")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-16")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-17")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-18")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-19")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-20")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-21")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-22")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-23")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-24")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-25")},
	},
}

var case1WorkspaceIds = []*workspaces.StartRequest{
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-1")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-2")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-3")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-4")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-5")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-6")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-7")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-8")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-9")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-10")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-11")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-12")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-13")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-14")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-15")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-16")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-17")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-18")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-19")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-20")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-21")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-22")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-23")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-24")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-25")},
}

var case2Expected = [][]*workspaces.StartRequest{
	{
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-3")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-4")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-5")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-6")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-7")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-8")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-9")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-10")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-11")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-12")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-13")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-14")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-15")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-16")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-17")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-18")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-19")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-20")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-21")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-22")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-23")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-24")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-25")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-26")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-27")}},
	{
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-1")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-2")},
	},
}

var case2WorkspaceIds = []*workspaces.StartRequest{
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-1")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-2")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-3")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-4")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-5")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-6")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-7")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-8")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-9")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-10")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-11")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-12")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-13")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-14")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-15")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-16")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-17")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-18")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-19")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-20")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-21")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-22")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-23")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-24")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-25")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-26")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-27")},
}

var case3Expected = [][]*workspaces.StartRequest{
	{
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-28")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-29")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-30")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-31")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-32")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-33")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-34")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-35")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-36")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-37")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-38")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-39")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-40")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-41")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-42")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-43")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-44")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-45")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-46")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-47")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-48")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-49")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-50")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-51")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-52")}},
	{
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-3")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-4")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-5")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-6")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-7")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-8")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-9")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-10")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-11")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-12")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-13")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-14")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-15")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-16")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-17")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-18")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-19")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-20")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-21")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-22")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-23")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-24")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-25")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-26")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-27")}},
	{
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-1")},
		&workspaces.StartRequest{
			WorkspaceId: returnPointer("w-2")}},
}

var case3WorkspaceIds = []*workspaces.StartRequest{
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-1")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-2")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-3")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-4")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-5")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-6")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-7")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-8")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-9")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-10")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-11")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-12")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-13")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-14")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-15")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-16")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-17")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-18")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-19")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-20")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-21")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-22")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-23")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-24")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-25")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-26")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-27")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-28")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-29")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-30")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-31")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-32")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-33")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-34")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-35")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-36")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-37")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-38")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-39")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-40")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-41")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-42")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-43")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-44")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-45")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-46")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-47")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-48")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-49")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-50")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-51")},
	&workspaces.StartRequest{
		WorkspaceId: returnPointer("w-52")},
}
