resource "google_storage_bucket" "db" {
  name     = "${var.project}-db"
  location = var.region
}
