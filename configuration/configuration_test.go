package configuration

import (
    "testing"
    "io/ioutil"

    "github.com/stretchr/testify/require"
)

func TestNewConfiguration(t *testing.T) {
    require := require.New(t)

    c, err := New("config_test.yaml")

    require.Nil(err)

    require.NotNil(c.Database.Host)
    require.Equal(c.Database.Host, "127.0.0.1")

    require.NotNil(c.Database.Username)
    require.Equal(c.Database.Username, "postgres")

    require.NotNil(c.Database.Password)
    require.Equal(c.Database.Password, "postgres")

    require.NotNil(c.Database.DbName)
    require.Equal(c.Database.DbName, "testdb")

    require.Equal(c.Database.Port, 5432)
}

func TestStringConnection(t *testing.T) {
    require := require.New(t)

    c, err := New("config_test.yaml")
    require.Nil(err)

    require.Equal(c.StringConnection(),
        "postgres://postgres:postgres@127.0.0.1:5432?dbname=testdb&sslmode=disable")
}

func TestConfigurationNotFound(t *testing.T) {
    require := require.New(t)
    c, err := New("does-not-exist")

    require.Equal(c, Configuration{})
    require.NotNil(err)
}

func TestConfigurationInvalidFormat(t *testing.T) {
    data := []byte(`
---
database:
  host: 127.0.0.1
  username: postgres
  INVALID FORMAT
  password: postgres
  port: 5432
  dbname: testdb
`)
    testFile := "/tmp/wrong.yaml"
    if err := ioutil.WriteFile(testFile, data, 0644); err != nil {
        t.Fatal(err)
    }

    require := require.New(t)
    _, err := New(testFile)
    require.NotNil(err)
}