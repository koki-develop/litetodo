resource "google_cloud_run_v2_service" "main" {
  depends_on = [google_project_service.main]

  name     = "${var.project}-app"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    service_account = google_service_account.app.email

    scaling {
      max_instance_count = 1
    }

    containers {
      name  = "app"
      image = "${google_artifact_registry_repository.app.location}-docker.pkg.dev/${var.project}/${google_artifact_registry_repository.app.name}/app:latest"
      ports {
        container_port = 8080
      }
    }
  }
}

data "google_iam_role" "run_invoker" {
  name = "roles/run.invoker"
}

data "google_iam_policy" "cloud_run_noauth" {
  binding {
    role    = data.google_iam_role.run_invoker.name
    members = ["allUsers"]
  }
}

resource "google_cloud_run_service_iam_policy" "cloud_run_noauth" {
  location    = google_cloud_run_v2_service.main.location
  project     = google_cloud_run_v2_service.main.project
  service     = google_cloud_run_v2_service.main.name
  policy_data = data.google_iam_policy.cloud_run_noauth.policy_data
}
