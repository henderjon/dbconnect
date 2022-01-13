package dsn

import (
	"net/url"
	"strings"
)

// DSN is a struct that represents a typical DSN
type Config struct {
	Username string
	Password string
	Protocol string
	Address  string
	DBName   string
	Opts     *url.Values
}

type DSN struct {
	c *Config
}

func NewDSN(c *Config) *DSN {
	if c == nil {
		c = &Config{
			Opts: &url.Values{},
		}
	}

	if c.Opts == nil {
		c.Opts = &url.Values{}
	}

	return &DSN{c}
}

func (dsn *DSN) Username(u string) {
	dsn.c.Username = u
}

func (dsn *DSN) Password(p string) {
	dsn.c.Password = p
}

func (dsn *DSN) Protocol(p string) {
	dsn.c.Protocol = p
}

func (dsn *DSN) Address(a string) {
	dsn.c.Address = a
}

func (dsn *DSN) DBName(db string) {
	dsn.c.DBName = db
}

func (dsn *DSN) Opts(key, value string) {
	dsn.c.Opts.Set(key, value)
}

// String creates a string of the DSN, the output is formatteed `[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]`
func (dsn *DSN) String() string {
	var final strings.Builder

	if dsn.c.Username != "" {
		final.WriteString(dsn.c.Username)
	}

	if dsn.c.Password != "" {
		final.WriteString(":")
		final.WriteString(dsn.c.Password)
	}

	if final.Len() > 0 {
		final.WriteString("@")
	}

	if dsn.c.Protocol != "" {
		final.WriteString(dsn.c.Protocol)
	} else {
		final.WriteString("tcp")
	}

	if dsn.c.Address != "" {
		final.WriteString("(")
		final.WriteString(dsn.c.Address)
		final.WriteString(")")
	}

	if dsn.c.DBName != "" {
		final.WriteString("/")
		final.WriteString(dsn.c.DBName)
	}

	params := dsn.c.Opts.Encode()
	if params != "" {
		final.WriteString("?")
		final.WriteString(params)
	}

	return final.String()
}

// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
