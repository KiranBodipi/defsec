---
additional_links: 
  - "https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#user_data"
---

Remove sensitive data from the EC2 instance user-data

```hcl
resource "aws_iam_instance_profile" "good_example" {
		 // ...
 }
 
 resource "aws_instance" "good_example" {
	 ami           = "ami-12345667"
	 instance_type = "t2.small"
 
	 iam_instance_profile = aws_iam_instance_profile.good_profile.arn
 
	 user_data = <<EOF
	 export GREETING=hello
 EOF
 }
```