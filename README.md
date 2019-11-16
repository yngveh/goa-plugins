# goa plugins
Collection of plugins to [goa](https://github.com/goadesign/goa) micro-services framework

## Overview

## Usage 
Get source
```bash
go get -u github.com/yngveh/goa-plugins
```

### logrus plugin
This plugin will setup logrus as logging framework when generating example code with
```bash
goa example github.com/user/project/design
```

To active the logrus plugin add the following to the design go file
```go
package design

import (
    // Activate logrus plugin 
    _ "github.com/yngveh/goa-plugins/logrus"
    . "goa.design/goa/v3/dsl"
)

var _ = API(........
``` 

## controllers plugin
This plugin puts controller files into the directory controllers instead of project root when doing goa example generating

To active the logrus plugin add the following to the design go file
```go
package design

import (
    // Activate controllers plugin 
    _ "github.com/yngveh/goa-plugins/controllers"
    . "goa.design/goa/v3/dsl"
)

var _ = API(........
``` 


