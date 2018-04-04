---
title: Receive Redis Message
weight: 4703
---
# tibco-redissub
This trigger allows you to start a flow with the contents of the message from Redis Channel.  It is assumed that the messages plain text. 

## Installation

```bash
flogo install github.com/sachinroham/flogo-contrib/redis/trigger/redissub
```

## Schema
Settings, Outputs :

```json
{
 "settings": [
		{
			"name": "host",
			"type": "string",
			"required": true
		},
		{
			"name": "port",
			"type": "int",
			"required": true
		},
		{
			"name": "password",
			"type": "string",
			"required": false
		},
		{
			"name": "channel",
			"type": "string",
			"required": true
		}
	],
	"output": [
		{
			"name": "message",
			"type": "string"
		},
		{
			"name": "channel",
			"type": "string"
		}
	],
```

## Example Configurations
This example flow subscribes to the syslog subject of bilbo's kafka server using a plain text connection with no authentication.

TODO