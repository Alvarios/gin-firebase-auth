# Gin-Firebase-Auth (GFA)

GFA is a package of [gin](https://github.com/gin-gonic/gin) middlewares to handle authentification using [Google Firebase](https://github.com/firebase/firebase-admin-go).

## Setup

First, you need to [initialize your firebase app](https://firebase.google.com/docs/admin/setup).
Once you have your `app` variable, create a new GFA config object:

```go
config := gfa.Handler{}
config.Init(app) // Pass your firebase app variable.
```

gfaHandler will automatically declare some firebase objects to use inside middlewares.

## Middlewares

Each gfa middlewares use a wrapper, to pass in some parameters.
To declare them, use :

```go
router.METHOD(
    path,
    mid(config, ...opts),
    ...
)
```

Each section comes with a how-to-use example, for better clarity.

### VerifyToken

| Option | Type | Description | 
| :--- | :--- | :--- |
| nonRestrictivePolicy | bool | When authentication fails, the request is aborted with an unauthorized status.<br/>Set this flag to true to override this behavior. |

Go handles auth with credentials from its own api. Once an user is logged, each
secure request will be issued with an Authorization header, containing a token id
that handles credentials.

VerifyToken will check and decode this token.

```go
package main

import "github.com/Alvarios/gin-firebase-auth"

func main() {
    // Set your gin and gfa apps here.
    
    router.GET(
        "/path/to/resource1",
        gfa.VerifyToken(config, false),
        GetResource1,   // Use your own functions here.
        Render,
    )
}
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/gin-firebase-auth/blob/master/LICENSE)
