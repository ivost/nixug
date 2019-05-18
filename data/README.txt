/etc/group file format:

group_name: It is the name of group. If you run ls -l command, you will see this name printed in the group field.
Password: Generally password is not used, hence it is empty/blank. ...
Group ID (GID): Each user must be assigned a group ID. ...
Group List: It is a list of user names of users who are members of the group.



The /etc/passwd file is a colon-separated file that contains the following information:
User name.
Encrypted password.
User ID number (UID)
User's group ID number (GID)
Full name of the user (GECOS)
User home directory.
Login shell.


https://github.com/fsnotify/fsnotify


watcher, err := fsnotify.NewWatcher()
if err != nil {
    log.Fatal(err)
}
defer watcher.Close()

done := make(chan bool)
go func() {
    for {
        select {
        case event, ok := <-watcher.Events:
            if !ok {
                return
            }
            log.Println("event:", event)
            if event.Op&fsnotify.Write == fsnotify.Write {
                log.Println("modified file:", event.Name)
            }
        case err, ok := <-watcher.Errors:
            if !ok {
                return
            }
            log.Println("error:", err)
        }
    }
}()

err = watcher.Add("/tmp/foo")
if err != nil {
    log.Fatal(err)
}
<-done