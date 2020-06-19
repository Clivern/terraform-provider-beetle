// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package beetle

import (
	"log"

	"github.com/clivern/beetle/app/module"
	"github.com/clivern/beetle/sdk"
)

// Client provider client
type Client struct {
	Client *sdk.Client
}

// Config provider config
type Config struct {
	APIKey string
	APIURL string
}

// Client gets an instance of the provider client
func (c *Config) Client() (*Client, error) {
	cli := &sdk.Client{}
	cli.SetHTTPClient(module.NewHTTPClient())
	cli.SetAPIKey(c.APIKey)
	cli.SetAPIURL(c.APIURL)

	log.Printf("[INFO] Upstream Client Configured")

	return &Client{Client: cli}, nil
}
