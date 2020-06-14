
provider "beetle" {
  api_key = "key"
  api_url = "http://127.0.0.1:8080"
}

data "beetle_image" "image" {
  slug = "UBUNTU_18_04_64BIT"
}

resource "beetle_server" "web" {
    name = "web"
    image = data.beetle_image.image.id
    region = "eu"
    size = "small"
}

