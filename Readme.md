<h1>Easy to use SFTP Uploader out of the box</h1>

##Example
````go
package main

import (
	"fmt"
	"github.com/itslearninggermany/sftp"
	)

func main ()  {
 	sftpServer :=  sftp.NewSFTPUpload("username","password", "server")
 	fmt.Println(sftpServer.SetTargetFolder("upload/").UploadAFile("./","dd.html"))
	fmt.Println(sftpServer.SetTargetFolder("upload/").SetFilenameOnServer("test.txt").UploadContent([]byte("Content")))
}
````