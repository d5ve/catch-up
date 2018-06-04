# catch-up
My first golang webapp - a tool for arranging the best times to meet with friends

This is likely to stay simple, part working, and ugly.

## First attempt

gin/gonic, SQL boiler, ...

This was working up to a point, then I foolishly attempted to add vendoring
using `dep` to make it easier to coordinate between my two computers. Then it
all stopped working with the following error.

    $ go run main.go routes.go
    # command-line-arguments
    ./routes.go:64:13: c.ShouldBind undefined (type *gin.Context has no field or method ShouldBind)
    ./routes.go:86:31: binding.Validator.Engine undefined (type binding.StructValidator has no field or method Engine)

I can't be bothered trying to work this out, so I'll reboot and start again.
