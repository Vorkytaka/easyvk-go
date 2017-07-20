# EasyVK [![GoDoc](https://godoc.org/github.com/vorkytaka/easyvk-go/easyvk?status.svg)](https://godoc.org/github.com/vorkytaka/easyvk-go/easyvk)
Package EasyVK provides you simple way to work with VK API.

### Installation:
#### Install:
```
$ go get -u github.com/vorkytaka/easyvk-go/easyvk
```
#### Import:
```go
import "github.com/vorkytaka/easyvk-go/easyvk"
```

### How to work:
Initialize your VK object with your access token.
```go
vk := easyvk.WithToken("token")
```
Or you can log in by email and password.
```go
// put your email, password, client id and scope
// scope must be a string like "friends,wall"
vk, err := easyvk.WithAuth("my@beautiful.mail", "pa$$word", "9182736", "friends,wall,photos")
if err != nil {
	// doesn't log in
}
```

Now you can call method of VK API with your vk variable.

### Examples:
Get user profile info:
```go
info, err := vk.Account.GetProfileInfo()
```
Set user status:
```go
// If you want to update status on your page
// then set ID to 0
userID := 0
ok, err := vk.Status.Set("New status", userID)
```
Post photo on wall:
```go
id := 0

server, err := vk.Photos.GetWallUploadServer(uint(id))
if err != nil {
	log.Fatal(err)
}

// path to the image
path := "D:/x.png"
uploaded, err := vk.Upload.PhotoWall(server.UploadURL, path)
if err != nil {
	log.Fatal(err)
}

saved, err := vk.Photos.SaveWallPhoto(0, uint(id), uploaded.Photo, uploaded.Hash, "", uploaded.Server, 0, 0)
if err != nil {
	log.Fatal(err)
}

text := "Caption for the post"
photoID := "photo" + fmt.Sprint(saved[0].OwnerID) + "_" + fmt.Sprint(saved[0].ID)

// -id if you post to group wall
x, err := vk.Wall.Post(id, false, true, false, false, false, text, photoID, "", "", 0, 0, 0, 0, 0)
if err != nil {
	log.Fatal(err)
}
fmt.Println(x)
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
    * [BanUser](https://vk.com/dev/account.banUser)
    * [GetAppPermissions](https://vk.com/dev/account.getAppPermissions)
    * [GetBanned](https://vk.com/dev/account.getBanned)
    * [GetCounters](https://vk.com/dev/account.getCounters)
    * [GetInfo](https://vk.com/dev/account.getInfo)
    * [GetProfileInfo](https://vk.com/dev/account.getProfileInfo)
    * [UnbanUser](https://vk.com/dev/account.unbanUser)
* [Fave](https://vk.com/dev/fave)
    * [GetLinks](https://vk.com/dev/fave.getLinks)
    * [GetPhotos](https://vk.com/dev/fave.getPhotos)
    * [GetUsers](https://vk.com/dev/fave.getUsers)
    * [GetVideos](https://vk.com/dev/fave.getVideos)
* [Photos](https://vk.com/dev/photos)
    * [GetWallUploadServer](https://vk.com/dev/photos.getWallUploadServer)
    * [SaveWallPhoto](https://vk.com/dev/photos.photos.saveWallPhoto)
* [Status](https://vk.com/dev/status)
    * [Get](https://vk.com/dev/status.get)
    * [Set](https://vk.com/dev/status.set)
* [Wall](https://vk.com/dev/wall)
    * [Post](https://vk.com/dev/wall.post)
* Upload
    * PhotoWall