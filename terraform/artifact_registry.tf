resource "google_artifact_registry_repository" "app" {
  depends_on = [google_project_service.main]

  location      = var.region
  repository_id = "app"
  format        = "DOCKER"
}

resource "google_artifact_registry_repository" "replicate" {
  depends_on = [google_project_service.main]

  location      = var.region
  repository_id = "replicate"
  format        = "DOCKER"
}

resource "google_artifact_registry_repository" "restore" {
  depends_on = [google_project_service.main]

  location      = var.region
  repository_id = "restore"
  format        = "DOCKER"
}
