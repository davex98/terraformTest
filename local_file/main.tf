

resource "local_file" "hello_world" {
  filename = var.file_name
  content  = var.file_content
}
