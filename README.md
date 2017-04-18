# EasyVK [![GoDoc](https://godoc.org/github.com/SolidlSnake/easyvk-go/easyvk?status.svg)](https://godoc.org/github.com/SolidlSnake/easyvk-go/easyvk)
Package EasyVK provides you simple way to work with VK API.

### Installation:
#### Install:
```
$ go get -u github.com/solidlsnake/easyvk-go/easyvk
```
#### Import:
```go
import "github.com/solidlsnake/easyvk-go/easyvk"
```

### How to work:
Initialize your VK object with your access token.
```go
vk := easyvk.WithToken("token")
```
Now you can call method of VK API with your vk variable.

### Examples:
Get user profile info:
```go
info, err := vk.Account.GetProfileInfo()
```
Set user status:
```go
userID := 1
ok, err := vk.Status.Set("New status", userID)
```

### If you need to call method that not done yet:
```go
methodName := "account.banUser"
params := map[string]string{
        "user_id": "1",
}
byteArray, err := vk.Request(methodName, params)
if err != nil {
        
}
// Now work with byteArray.
// Parse it to struct or to interface.
```

#### List of finished methods:
* [Account](https://vk.com/dev/account)
    * [GetAppPermissions](https://vk.com/dev/account.getAppPermissions)
    * [GetCounters](https://vk.com/dev/account.getCounters)
    * [GetInfo](https://vk.com/dev/account.getInfo)
    * [GetProfileInfo](https://vk.com/dev/account.getProfileInfo)
* [Fave](https://vk.com/dev/fave)
    * [GetUsers](https://vk.com/dev/fave.getUsers)
    * [GetLinks](https://vk.com/dev/fave.getLinks)
* [Photos](https://vk.com/dev/photos)
    * [GetWallUploadServer](https://vk.com/dev/photos.getWallUploadServer)
* [Status](https://vk.com/dev/status)
    * [Get](https://vk.com/dev/status.get)
    * [Set](https://vk.com/dev/status.set)