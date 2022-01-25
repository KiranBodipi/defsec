package apigateway

import (
    "testing"

    "github.com/aquasecurity/defsec/provider/aws/apigateway"
    "github.com/aquasecurity/defsec/state"
    "github.com/stretchr/testify/assert"
)

func TestCheckEnableTracing(t *testing.T) {
    t.SkipNow()
    tests := []struct{
        name string
        input apigateway.APIGateway
        expected bool
    }{
        {
            name: "positive result",
            input: apigateway.APIGateway{},
            expected: true,
        },
        {
            name: "negative result",
            input: apigateway.APIGateway{},
            expected: false,
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T){
            var testState state.State
            testState.AWS.APIGateway = test.input
            results := CheckEnableTracing.Evaluate(&testState)
            var found bool
            for _, result := range results {
                if result.Rule().LongID() == CheckEnableTracing.Rule().LongID() {
                    found = true
                }
            }
            if test.expected {
                assert.True(t, found, "Rule should have been found")
            } else {
                assert.False(t, found, "Rule should not have been found")
            }
        })
    }
}
