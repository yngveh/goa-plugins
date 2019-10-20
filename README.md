# goa plugins
Collection of plugins to [goa](https://github.com/goadesign/goa) micro-services framework

## Overview

## Usage 
Get source
```bash
go get -u github.com/yngveh/goa-plugins
```

### logrus plugin
In the design file
```go
package design

import (
    // Activate logrus plugin 
    _ "github.com/yngveh/goa-plugins/logrus"
    . "goa.design/goa/v3/dsl"
)

var _ = API(........
``` 

