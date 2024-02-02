resource "google_artifact_registry_repository" "app" {
  depends_on = [google_project_service.main]

  location      = var.region
  repository_id = "litetodo"
  format        = "DOCKER"
}
