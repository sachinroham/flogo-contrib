---
title: FTP GET
weight: 4618
---

# FTP GET
This activity invokes a FTP GET command.

## Installation

### Flogo CLI
```bash
flogo add activity github.com/sachinroham/flogo-contrib/ftp/retr
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
			"required": false
		},
		{
			"name": "user",
			"type": "string",
			"required": true
		},
		{
			"name": "password",
			"type": "string",
			"required": false
		},
		{
			"name": "localpath",
			"type": "string",
			"required": false
		},
		{
			"name": "remotepath",
			"type": "any",
			"required": true
		}
	],
	"output": [
		{
			"name": "result",
			"type": "any"
		}
	]
}
```
## Settings


## Examples
