# Golang_Test_Project

This is a test project using GO programming language.

This project aims to receive some rectangles information (one of them is the main one) in the POST request in the server and returns those rectangles having overlaps with the main rectangle as a result of the GET request to the client.

To run the code, you can use the follwoing command(you should have GO installed on your system):
```
go run server.go
```

Sample POST request:

```
{
"main": {"x": 0, "y": 0, "width": 10, "height": 20},
"input": [
{"x": 2, "y": 18, "width": 5, "height": 4},
{"x": 12, "y": 18, "width": 5, "height": 4},
{"x": -1, "y": -1, "width": 5, "height": 4}
]
}
```

Sample GET response:

```
[
  {
    "x": 2,
    "y": 18,
    "width": 5,
    "height": 4,
    "time": "2022-08-17 12:59:18 AM"
  },
  {
    "x": -1,
    "y": -1,
    "width": 5,
    "height": 4,
    "time": "2022-08-17 12:59:18 AM"
  }
]
```
