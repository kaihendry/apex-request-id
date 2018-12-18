deploylambda:
	apex -r ap-southeast-1 deploy

invokelambda:
	apex -r ap-southeast-1 invoke r < event.json
