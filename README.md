<h1>README</h1>
The project is created under Jetbrain GoLand Go(GOPATH) project configuration and "Go SDK 1.21rc1". 

OS: Windows 11 Home 22H2

The name of the directory is the default name of project created in GoLand.

<h2>Main.go</h2>
main.go contains code for the web server, receipt processing, and point calculations. The receipt processing function process all receipts file in the "Receipt" folder. Points are calculated as soon the receipts is processed.

Command used to start main.go:

```
go run main.go 
````


<h2>unit_test.go</h2>
unit_test.go contains the test for the point calculations. It uses the JSON provided in the README of the github to validate the result.

Command used to start unit tests: 

``` 
go test -v unit_test go main.go
```

#   G o A P I P r o j e c t 
 
 
