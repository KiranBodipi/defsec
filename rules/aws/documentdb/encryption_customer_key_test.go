package documentdb

import (
    "testing"

    "github.com/aquasecurity/defsec/provider/aws/documentdb"
    "github.com/aquasecurity/defsec/state"
    "github.com/stretchr/testify/assert"
)

func TestCheckEncryptionCustomerKey(t *testing.T) {
    t.SkipNow()
    tests := []struct{
        name string
        input documentdb.DocumentDB
        expected bool
    }{
        {
            name: "positive result",
            input: documentdb.DocumentDB{},
            expected: true,
        },
        {
            name: "negative result",
            input: documentdb.DocumentDB{},
            expected: false,
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T){
            var testState state.State
            testState.AWS.DocumentDB = test.input
            results := CheckEncryptionCustomerKey.Evaluate(&testState)
            var found bool
            for _, result := range results {
                if result.Rule().LongID() == CheckEncryptionCustomerKey.Rule().LongID() {
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
