
This application is used to wrap syncthing so that it can run in the background without a terminal open. This is meant for Windows, as it uses the windows commands
taskkill and tasklist in order to kill running processes. It can easily be ported to use ps.

You will need to run 
`go build` to create an executable to run, and then, as an administrator, run `.\main.exe --service install`.
Then if you navigate to Services, you should see the service 'Syncthing Wrapper' listed, and you can start it from there.

Sample output of tasklist command after being split by newline when the process is running

[]string len: 5, cap: 5, ["\r","Image Name                     PID Session Name        Session#    Mem Usage\r","========================= ======== ================ =========== ============\r","syncthing.exe                45204 Console                    1     61,800 K\r",""]

