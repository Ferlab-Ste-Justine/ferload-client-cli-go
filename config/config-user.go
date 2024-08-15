package config

import (
	"fmt"

	"github.com/magiconair/properties"
)

const ConfigUserFile = ".ferload-client.properties"
const ConfigUserPath = "${HOME}/" + ConfigUserFile

type ConfigUser struct {
	FerloadUrl       string `properties:"ferload-url,default="`
	UserName         string `properties:"username,default="`
	Token            string `properties:"token,default="`
	KeycloakUrl      string `properties:"keycloak-url,default="`
	KeycloakRealm    string `properties:"keycloak-realm,default="`
	KeycloakClientId string `properties:"keycloak-client-id,default="`
	KeycloakAudience string `properties:"keycloak-audience,default="`
}

func LoadDefaultConfigUser() *ConfigUser {
	return LoadConfigUser(ConfigUserPath)
}

func LoadConfigUser(path string) *ConfigUser {
	cfg := &ConfigUser{}
	p := properties.MustLoadFiles([]string{path}, properties.UTF8, true)
	if err := p.Decode(cfg); err != nil {
		panic(fmt.Sprintf("Failed to read user configuration file: %s reason: %s", ConfigUserPath, err))
	}
	return cfg
}
