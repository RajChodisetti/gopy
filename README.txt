A program that logs the relative filepaths along with timestamp when called.

Requirements:
Python 3.9.6
Go go1.21.5
Flask Version: 3.0.2

How to run :
Step 1 : Start app.py with command "python app.py"
Step 2 : Now run the go program main.go with "go run main.go" in  another terminal.
Step 3 : Go program listens on port 8080. 
         You can search for pwd with "localhost:8080/api"
         (OR)
         You can post the directory you want to search for as following:
         "localhost:8080/api?directory=Path_you_want_to_search"
Step 4 : You can terminate the programs with control+c in both terminals.
