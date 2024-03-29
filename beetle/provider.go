// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package beetle

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider returns instance of the provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BEETLE_API_KEY", nil),
			},
			"api_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "api.beetle.com",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"beetle_server": resourceBeetleApplication(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure configures the provider
func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	log.Println("[INFO] Initializing Client")

	config := Config{
		APIKey: data.Get("api_key").(string),
		APIURL: data.Get("api_url").(string),
	}

	return config.Client()
}
