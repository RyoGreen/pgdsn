package pgdsn_test

import (
	"testing"

	"github.com/RyoGreen/pgdsn"
)

func TestFormatDSN(t *testing.T) {
	tests := []struct {
		Name           string
		Config         *pgdsn.Config
		ExpectedResult string
	}{
		{
			Name: "WithSSLDisableMode",
			Config: &pgdsn.Config{
				User:     "testuser",
				DbName:   "testdb",
				Password: "testpass",
				Host:     "localhost",
				Port:     5432,
				SslMode:  pgdsn.Disable,
			},
			ExpectedResult: "user=testuser dbname=testdb password=testpass host=localhost port=5432 sslmode=disable",
		},
		{
			Name: "WithSSLRequireMode",
			Config: &pgdsn.Config{
				User:     "testuser",
				DbName:   "testdb",
				Password: "testpass",
				Host:     "localhost",
				Port:     5432,
				SslMode:  pgdsn.Require,
			},
			ExpectedResult: "user=testuser dbname=testdb password=testpass host=localhost port=5432 sslmode=require",
		},
		{
			Name: "WithSSLVerifyCaMode",
			Config: &pgdsn.Config{
				User:     "testuser",
				DbName:   "testdb",
				Password: "testpass",
				Host:     "localhost",
				Port:     5432,
				SslMode:  pgdsn.VerifyCa,
			},
			ExpectedResult: "user=testuser dbname=testdb password=testpass host=localhost port=5432 sslmode=verify-ca",
		},
		{
			Name: "WithSSLVerifyFullMode",
			Config: &pgdsn.Config{
				User:     "testuser",
				DbName:   "testdb",
				Password: "testpass",
				Host:     "localhost",
				Port:     5432,
				SslMode:  pgdsn.VerifyFull,
			},
			ExpectedResult: "user=testuser dbname=testdb password=testpass host=localhost port=5432 sslmode=verify-full",
		},
		{
			Name: "WithoutSSLMode",
			Config: &pgdsn.Config{
				User:     "testuser",
				DbName:   "testdb",
				Password: "testpass",
				Host:     "localhost",
				Port:     5432,
			},
			ExpectedResult: "user=testuser dbname=testdb password=testpass host=localhost port=5432",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actualDSN := test.Config.FormatDSN()
			if actualDSN != test.ExpectedResult {
				t.Errorf("Expected DSN: %s, but got DSN: %s", test.ExpectedResult, actualDSN)
			}
		})
	}
}
