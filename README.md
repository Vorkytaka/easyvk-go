# EasyVK
Package EasyVK provides you simple way work with VK API.

#### How to work:
Initialize your VK object with your access token.
```go
vk := easyvk.WithToken("token")
```
Now you can call method of VK API with your vk variable.

##### Examples:
Get user profile info:
```go
info, err := vk.Account.GetProfileInfo()
```
Set user status:
```go
userID := 1
ok, err := vk.Status.Set("New status", 1)
```

##### If you need to call method that not done yet:
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
* [Photos](https://vk.com/dev/photos)
    * [GetWallUploadServer](https://vk.com/dev/photos.getWallUploadServer)
* [Status](https://vk.com/dev/status)
    * [Get](https://vk.com/dev/status.get)
    * [Set](https://vk.com/dev/status.set)