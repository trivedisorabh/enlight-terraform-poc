output "public_zone_id" {
  value = "${module.iac.zone_id}"
}

output "website_bucket_id" {
  value = "${module.iac.website_bucket_id}"
}

output "website_cloudfront_id" {
  value = "${module.iac.website_cloudfront_id}"
}
