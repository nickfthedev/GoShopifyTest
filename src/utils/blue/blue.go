package blue

import "github.com/microcosm-cc/bluemonday"

// Init bluemonday policys
// Do this once for each unique policy, and use the policy for the life of the program
// Policy creation/editing is not safe to use in multiple goroutines

var UGC *bluemonday.Policy
var StrictPolicy *bluemonday.Policy

func InitBlueMondayPolicies() {

	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	UGC = bluemonday.UGCPolicy()
	StrictPolicy = bluemonday.StrictPolicy()

}
