# EasyVK [![GoDoc](https://godoc.org/github.com/vorkytaka/easyvk-go/easyvk?status.svg)](https://godoc.org/github.com/vorkytaka/easyvk-go/easyvk)
Package EasyVK provides you simple way to work with VK API.

![Logo](./logo.png)

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

params := easyvk.WallPostParams{}
params.OwnerID = id
params.Message = "Test"
params.Attachments = photoID

x, err := vk.Wall.Post(params)
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

### Naming conventions:

#### API section:

Every API section must be a structure and have name as it called in API and starts with a capital letter.

*For example:*
[Account](https://vk.com/dev/account) section must be  ```type Account struct {}```

#### API methods:

Every API method must be a method that have a received type of his section.
It must have name as it called in API and starts with a capital letter.

*For example:*
[Account.banUser](https://vk.com/dev/account.banUser) method must be ```func (a *Account) BanUser() {}```

#### Parameters:

If method requests 4 or less parameters, then they must be just a parameters.

*For example:* ```func (a *Account) BanUser(userID uint) {}```

If method requests 5 or more parameters, then he must get it with a structure.
That structure must naming like ```[SectionName][MethodName]Params```.

*For example:* ```type WallPostParams struct {}```

#### Responses:

If method return only one field, then you must just return that field.

*For example:*
[Board.addTopic](https://vk.com/dev/board.addTopic) return ID of the created topic,
so we just return integer.

If method return only 1 if succeeded, then we need to return boolean.

*For example:*
[Board.closeTopic](https://vk.com/dev/board.closeTopic) return 1 is succeded, so
we check if it's 1 and return true, and false otherwise.

If method return object with 2 or more fields, then we need to create a structure for that response.
That structure must naming like ```[SectionName][MethodName]Response```.

*For example*: ```type AccountGetCountersResponse struct {}```

### List of finished methods:
* [Account](https://vk.com/dev/account)
    * [BanUser](https://vk.com/dev/account.banUser)
    * [GetAppPermissions](https://vk.com/dev/account.getAppPermissions)
    * [GetBanned](https://vk.com/dev/account.getBanned)
    * [GetCounters](https://vk.com/dev/account.getCounters)
    * [GetInfo](https://vk.com/dev/account.getInfo)
    * [GetProfileInfo](https://vk.com/dev/account.getProfileInfo)
    * [SetOffline](https://vk.com/dev/account.setOffline)
    * [SetOnline](https://vk.com/dev/account.setOnline)
    * [UnbanUser](https://vk.com/dev/account.unbanUser)
* [Board](https://vk.com/dev/board)
    * [AddTopic](https://vk.com/dev/board.addTopic)
    * [CloseTopic](https://vk.com/dev/board.closeTopic)
    * [DeleteTopic](https://vk.com/dev/board.deleteTopic)
    * [EditTopic](https://vk.com/dev/board.editTopic)
* [Fave](https://vk.com/dev/fave)
    * [GetLinks](https://vk.com/dev/fave.getLinks)
    * [GetPhotos](https://vk.com/dev/fave.getPhotos)
    * [GetUsers](https://vk.com/dev/fave.getUsers)
    * [GetVideos](https://vk.com/dev/fave.getVideos)
* [Likes](https://vk.com/dev/likes) ✓
    * [Add](https://vk.com/dev/likes.add)
    * [Delete](https://vk.com/dev/likes.delete)
    * [GetList](https://vk.com/dev/likes.getList)
    * [IsLiked](https://vk.com/dev/likes.isLiked)
* [Photos](https://vk.com/dev/photos)
    * [GetWallUploadServer](https://vk.com/dev/photos.getWallUploadServer)
    * [SaveWallPhoto](https://vk.com/dev/photos.saveWallPhoto)
* [Status](https://vk.com/dev/status) ✓
    * [Get](https://vk.com/dev/status.get)
    * [Set](https://vk.com/dev/status.set)
* [Wall](https://vk.com/dev/wall)
    * [Post](https://vk.com/dev/wall.post)
* Upload
    * PhotoWall