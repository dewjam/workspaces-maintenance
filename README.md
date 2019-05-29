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

#### To Do
1. Add CloudFormation stack to deploy automatically
1. Add support for Workspaces tags (only act on workspaces with a specific tag value)
1. Add 'dry-run' support

#### A Word of Caution
**Please note that I cannot commit to supporting the provided code.  This is meant to be used as an example and requires further testing to be used in a production environment**
