## Workspaces Maintenance Lambda Function
This code is designed to be executed via a scheduled CloudWatch event and will start any Workspaces that are currently in a "Stopped" state.  This is intended to give administrators more control around when Workspaces receive Windows and application updates.  

#### Requirements
- Golang 1.x (https://golang.org/dl/)
- glide (https://glide.sh/)

#### Build Steps
1. Execute the `local_build_script.sh` script
1. ZIP up the resulting binary file `bin/wm`

#### Deployment steps
1. Create a new Lambda function from scratch (make sure the IAM role has the capability to describe/start/stop AWS Workspaces)
1. Upload the zip (from the Build Step) as the Lambda "Function package" and set the handler to `wm`
1. Test the function (event content doesn't matter at this point)
