locals {
  services = [
    "artifactregistry.googleapis.com",
    "run.googleapis.com",
  ]
}

resource "google_project_service" "main" {
  for_each = toset(local.services)

  service = each.value
}
