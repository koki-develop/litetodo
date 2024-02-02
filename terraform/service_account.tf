locals {
  app_roles = ["roles/storage.admin"]
}

resource "google_service_account" "app" {
  account_id = "${var.project}-app"
}

resource "google_project_iam_member" "app" {
  for_each = toset(local.app_roles)

  project = var.project
  role    = each.value
  member  = "serviceAccount:${google_service_account.app.email}"
}
