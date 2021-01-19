variable "human" {
  type        = object({ name = string, age = number })
  description = "Human object"
}

variable "magic_tuple" {
  type        = tuple([string, number, bool])
  description = "Magic tuple"
}

variable "file_content" {
  type        = string
  description = "File content"
  default     = "Hello World!"
}

variable "file_name" {
  type = string
  default = "file.txt"
  description = "File name"
}