---
title: REDIS PUBLISH
weight: 4618
---

# REDIS PUBLISH
This activity publishes a message to a Redis channel.

## Installation

### Flogo CLI
```bash
flogo add activity github.com/sachinroham/flogo-contrib/redis/activity/redispub
```

## Schema
Inputs and Outputs:

```json
{
  "input": [
		{
			"name": "host",
			"type": "string",
			"required": true
		},
		{
			"name": "port",
			"type": "integer",
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
		},
		{
			"name": "message",
			"type": "string",
			"required": true
		}
	],
	"output": [
		{
			"name": "result",
			"type": "string"
		}
	]
}
```
## Settings


## Examples
