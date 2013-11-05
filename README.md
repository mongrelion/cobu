Cobu
====

Code Base Updater
Simple Go application for updating code base using [Git] and [Github] hooks.

---
*Requirements*
* [Go]
* [Git]

---
**Instructions**

1. Make sure that you have the **PPATH** environment variable set on your terminal,
pointing to the directory you want to be updated whenever [Github] sends a
post-hook when pushing code to your repo. ```export PPATH=/home/deploy/apps/foobar/```
2. If you have Go installed in your server, you can just run the application with `go run cobu.go`.  
   If not, search for instructions on the Internets on how to install Go on  
   your distro or how to cross compile.
3. Run the app.
That's it.

[Go]: http://golang.org
[Git]: http://git-scm.com/
[Github]: http://github.com
