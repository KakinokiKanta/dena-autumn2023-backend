terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.51.0"
    }
  }

  backend "gcs" {
    bucket  = "dena-autumn-tfstate"
  }
}

provider "google" {
  credentials = file("key.json")

  project = "dena-autumn"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}

resource "google_sql_database" "database" {
  name     = "my-database"
  instance = google_sql_database_instance.instance.name
}

resource "google_sql_database_instance" "instance" {
  name             = "my-database-instance"
  region           = "asia-northeast1"
  database_version = "MYSQL_8_0"
  settings {
    tier = "db-f1-micro"
  }

  deletion_protection  = "true"
}
