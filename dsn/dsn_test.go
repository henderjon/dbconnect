package dsn

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewDSN(t *testing.T) {
	d := NewDSN(nil)
	d.Username("user")
	d.Password("pass")
	d.Protocol("mysql")
	d.Address("127.0.0.1")
	d.DBName("mydb")
	d.Opts("allow-access", "true")

	expected := `user:pass@mysql(127.0.0.1)/mydb?allow-access=true`

	if diff := cmp.Diff(d.String(), expected); diff != "" {
		t.Errorf("-got/+want\n%s", diff)
	}
}

func Test_NewDSNIncomplete(t *testing.T) {
	d := NewDSN(nil)
	d.Username("user")
	// d.Password("pass")
	d.Protocol("mysql")
	d.Address("127.0.0.1:3306")
	d.DBName("mydb")
	// d.Opts("allow-access", "true")

	expected := `user@mysql(127.0.0.1:3306)/mydb`

	if diff := cmp.Diff(d.String(), expected); diff != "" {
		t.Errorf("-got/+want\n%s", diff)
	}
}

func Test_NewDSNConfig(t *testing.T) {
	d := NewDSN(&Config{
		Username: "user",
		Password: "pass",
		Protocol: "mysql",
		Address:  "127.0.0.1",
		DBName:   "mydb",
	})

	d.Opts("allow-access", "true")
	d.Opts("iam", "ironman")

	expected := `user:pass@mysql(127.0.0.1)/mydb?allow-access=true&iam=ironman`

	if diff := cmp.Diff(d.String(), expected); diff != "" {
		t.Errorf("-got/+want\n%s", diff)
	}
}
