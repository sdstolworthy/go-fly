# Go Fly

A simple application written in Go.

Go Fly hits the skyscanner api and finds the cheapest flights available for the specified destinations.

**Roadmap:**
1. Daemonize the application. The application should run in the background, and send notifications through email or another means
2. Database for analytics. Write prices to a database for historical purposes. Use this data to determine what a "good" deal is, by comparing to past price averages.
3. Support additional queries.
4. Build a web interface for customizing queries.