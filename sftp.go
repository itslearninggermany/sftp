package sftp

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

type Sftp struct {
	username string
	password string
	server string
	filename string
}


func NewSFTPUpload (username, password, server string) *Sftp{
	out := new(Sftp)
	out.username = username
	out.password = password
	out.server = server
	return out
}

func (p *Sftp) setFilenameOnServer (filename string) *Sftp {
	p.filename = filename
	return p

}

/*
stores the contend on the sftp-server
 */
func (p *Sftp) UploadContent (uploadcontent []byte) (counted int, err error){
	config := &ssh.ClientConfig{
		User:            p.username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(p.password),
		},
	}

	config.SetDefaults()
	sshConn, err := ssh.Dial("tcp", p.server+":22", config) //sftpServer
	if err != nil {
		return 0,err
	}
	defer sshConn.Close()

	c, err := sftp.NewClient(sshConn)
	if err != nil {
		return 0,err
	}
	defer c.Close()

	if p.filename == "" {
		p.filename = "upload.txt"
	}
	// Uploading the file
	remoteFile, err := c.Create(p.filename)
	//Start Error-Handling
	//1. Log error
	if err != nil {
		return 0,err
	}

	counted, err = remoteFile.Write(uploadcontent)
	if err != nil {
		return 0,err
	}

	return counted,err
}

/*
upload a file to sftpserver
 */
func (p *Sftp) UploadAFile (path,filename string) (counted int, err error) {
	config := &ssh.ClientConfig{
		User:            p.username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(p.password),
		},
	}

	config.SetDefaults()
	sshConn, err := ssh.Dial("tcp", p.server+":22", config) //sftpServer
	if err != nil {
		return 0,err
	}
	defer sshConn.Close()

	c, err := sftp.NewClient(sshConn)
	if err != nil {
		return 0,err
	}
	defer c.Close()

	remoteFile, err := c.Create(p.filename)
	if err != nil {
		return 0,err
	}

	uploadcontent, err := ioutil.ReadFile(fmt.Sprint(path,filename))
	if err != nil {
		return 0,err
	}
	counted, err = remoteFile.Write(uploadcontent)
	if err != nil {
		return 0,err
	}
	return counted,err
}