package sql

import (
    "testing"

    "github.com/aquasecurity/defsec/provider/google/sql"
    "github.com/aquasecurity/defsec/state"
    "github.com/stretchr/testify/assert"
)

func TestCheckNoCrossDbOwnershipChaining(t *testing.T) {
    t.SkipNow()
    tests := []struct{
        name string
        input sql.SQL
        expected bool
    }{
        {
            name: "positive result",
            input: sql.SQL{},
            expected: true,
        },
        {
            name: "negative result",
            input: sql.SQL{},
            expected: false,
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T){
            var testState state.State
            testState.Google.SQL = test.input
            results := CheckNoCrossDbOwnershipChaining.Evaluate(&testState)
            var found bool
            for _, result := range results {
                if result.Rule().LongID() == CheckNoCrossDbOwnershipChaining.Rule().LongID() {
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
