---
title: FTP PUT
weight: 4618
---

# FTP PUT
This activity invokes a FTP PUT command.

## Installation

### Flogo CLI
```bash
flogo add activity github.com/sachinroham/flogo-contrib/ftp/stor
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
			"name": "path",
			"type": "string",
			"required": true
		},
		{
			"name": "content",
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
