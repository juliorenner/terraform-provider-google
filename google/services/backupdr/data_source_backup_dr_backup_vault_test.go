// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package backupdr_test

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"testing"
)

func TestAccDataSourceGoogleBackupDRBackupVault_basic(t *testing.T) {
	t.Parallel()

	random_suffix := acctest.RandString(t, 10)
	context := map[string]interface{}{
		"random_suffix": random_suffix,
	}
	id := "tf-test-bv-" + random_suffix
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBackupDRBackupVaultDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleBackupDRBackupVault_basic(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_backup_dr_backup_vault.fetch-bv", "backup_vault_id", id),
				),
			},
		},
	})
}

func testAccDataSourceGoogleBackupDRBackupVault_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_backup_dr_backup_vault" "test-bv" {
  location = "us-central1"
  backup_vault_id = "tf-test-bv-%{random_suffix}"
  description = "This is a a backup vault built by Terraform."
  backup_minimum_enforced_retention_duration = "100000s"
  force_update = "true"
  force_delete = "true"
  allow_missing = "true"
  ignore_backup_plan_references = "false"
  ignore_inactive_datasources = "false"
}

data "google_backup_dr_backup_vault" "fetch-bv" {
  location = "us-central1"
  backup_vault_id = google_backup_dr_backup_vault.test-bv.backup_vault_id
}
`, context)
}