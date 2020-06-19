// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package beetle

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

// resourceBeetleApplication defines a server schema
func resourceBeetleApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceBeetleApplicationCreate,
		Read:   resourceBeetleApplicationRead,
		Update: resourceBeetleApplicationUpdate,
		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
				ValidateFunc: validation.NoZeroValues,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
				ValidateFunc: validation.NoZeroValues,
			},
			"application": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
				ValidateFunc: validation.NoZeroValues,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
				ValidateFunc: validation.NoZeroValues,
			},
			"strategy": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
				ValidateFunc: validation.NoZeroValues,
			},
		},
	}
}

// resourceBeetleApplicationCreate creates a server
func resourceBeetleApplicationCreate(d *schema.ResourceData, m interface{}) error {
	_ = m.(*Client).Client

	return resourceBeetleApplicationRead(d, m)
}

// resourceBeetleApplicationRead retrieves a server
func resourceBeetleApplicationRead(d *schema.ResourceData, m interface{}) error {
	_ = m.(*Client).Client

	return nil
}

// resourceBeetleApplicationUpdate updates a server
func resourceBeetleApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	_ = m.(*Client).Client

	return resourceBeetleApplicationRead(d, m)
}
