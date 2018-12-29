# Go Fly

A simple application written in Go.

Go Fly hits the skyscanner api and finds the cheapest flights available for the specified destinations.

# Getting Started

1. Clone the repo. `git clone git@github.com:sdstolworthy/go-fly.git && cd go-fly/`
2. Get the dependencies `dep ensure`
3. (Optional) Seed the database
    1. Run the seed files `go run seed/*.go`
4. Run the server from the main directory `go run *.go`

For convenience, I have included my `.vscode/launch.json`. You can use this to work with the debugging functions in VSCode.

**Wishlist:**
1. ~~Daemonize the application. The application should run in the background, and send notifications through email or another means~~
2. ~~Database for analytics. Write prices to a database for historical purposes. Use this data to determine what a "good" deal is, by comparing to past price averages.~~
3. Support additional queries.
4. [In Progress] Build a web interface for customizing queries.