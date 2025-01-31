/*
Use this data source to query detailed information of ssm ssh_key_pair_value

~> **NOTE:** Must set at least one of `secret_name` or `ssh_key_id`.

Example Usage

```hcl
data "tencentcloud_ssm_ssh_key_pair_value" "example" {
  secret_name = "keep_terraform"
  ssh_key_id  = "skey-2ae2snwd"
}
```

Or

```hcl
data "tencentcloud_ssm_ssh_key_pair_value" "example" {
  secret_name = "keep_terraform"
}
```

Or

```hcl
data "tencentcloud_ssm_ssh_key_pair_value" "example" {
  ssh_key_id  = "skey-2ae2snwd"
}
```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ssm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm/v20190923"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func dataSourceTencentCloudSsmSshKeyPairValue() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudSsmSshKeyPairValueRead,
		Schema: map[string]*schema.Schema{
			"secret_name": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Secret name.",
			},
			"ssh_key_id": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The key pair ID is the unique identifier of the key pair in the cloud server.",
			},
			"public_key": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Public key plain text, encoded using base64.",
			},
			"private_key": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Private key plain text, encoded using base64.",
			},
			"project_id": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "The project ID to which this key pair belongs.",
			},
			"ssh_key_name": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "SSH key name.",
			},
			"ssh_key_description": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Description of the SSH key pair. Users can modify the description information of the key pair in the CVM console.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudSsmSshKeyPairValueRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_ssm_ssh_key_pair_value.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId           = getLogId(contextNil)
		ctx             = context.WithValue(context.TODO(), logIdKey, logId)
		service         = SsmService{client: meta.(*TencentCloudClient).apiV3Conn}
		sshKeyPairValue *ssm.GetSSHKeyPairValueResponseParams
		sshKeyID        string
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("secret_name"); ok {
		paramMap["SecretName"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("ssh_key_id"); ok {
		paramMap["SSHKeyId"] = helper.String(v.(string))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeSsmSshKeyPairValueByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		sshKeyPairValue = result
		return nil
	})

	if err != nil {
		return err
	}

	if sshKeyPairValue.SSHKeyID != nil {
		_ = d.Set("ssh_key_id", sshKeyPairValue.SSHKeyID)
		sshKeyID = *sshKeyPairValue.SSHKeyID
	}

	if sshKeyPairValue.SSHKeyName != nil {
		_ = d.Set("ssh_key_name", sshKeyPairValue.SSHKeyName)
	}

	if sshKeyPairValue.PublicKey != nil {
		_ = d.Set("public_key", sshKeyPairValue.PublicKey)
	}

	if sshKeyPairValue.PrivateKey != nil {
		_ = d.Set("private_key", sshKeyPairValue.PrivateKey)
	}

	if sshKeyPairValue.ProjectID != nil {
		_ = d.Set("project_id", sshKeyPairValue.ProjectID)
	}

	if sshKeyPairValue.SSHKeyDescription != nil {
		_ = d.Set("ssh_key_description", sshKeyPairValue.SSHKeyDescription)
	}

	d.SetId(sshKeyID)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), d); e != nil {
			return e
		}
	}

	return nil
}
