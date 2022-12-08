package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/Ferlab-Ste-Justine/ferload-client-cli-go/config"
)

func getField(cfg *config.ConfigUser, field string) string {
	r := reflect.ValueOf(cfg)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func assertField(t *testing.T, cfg *config.ConfigUser, field string) {
	if getField(cfg, field) != "" {
		t.Errorf("%s should be empty", field)
	}
}

func createTempFile(prefix string) *os.File {
	file, err := ioutil.TempFile(os.TempDir(), config.ConfigUserFile)
	if err != nil {
		panic(err)
	}
	return file
}

func TestNewConfigUser_FileDoesntExist(t *testing.T) {
	cfg := config.LoadConfigUser("")
	assertField(t, cfg, "FerloadUrl")
	assertField(t, cfg, "UserName")
	assertField(t, cfg, "Token")
	assertField(t, cfg, "KeycloakUrl")
	assertField(t, cfg, "KeycloakClientId")
	assertField(t, cfg, "KeycloakRealm")
	assertField(t, cfg, "KeycloakAudience")
	t.Log("TestNewConfigUserEmpty OK")
}

func TestNewConfigUser_EmptyFile(t *testing.T) {

	tempCfg := createTempFile(config.ConfigUserFile).Name()
	t.Logf("Create a temp file: %s", tempCfg)
	defer os.Remove(tempCfg)

	cfg := config.LoadConfigUser(tempCfg)
	assertField(t, cfg, "FerloadUrl")
	assertField(t, cfg, "UserName")
	assertField(t, cfg, "Token")
	assertField(t, cfg, "KeycloakUrl")
	assertField(t, cfg, "KeycloakClientId")
	assertField(t, cfg, "KeycloakRealm")
	assertField(t, cfg, "KeycloakAudience")
	t.Log("TestNewConfigUserEmpty OK")
}
