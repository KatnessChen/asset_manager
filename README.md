# Asset Manager

Asset Manager is a Go-based web application for managing assets. It uses the Gin web framework and provides a RESTful API for asset management operations.

# How To Develop

To develop the Asset Manager application, you can use `nodemon` for automatic reloading of the server when changes are detected. Follow these steps:

1. Make sure you have `nodemon` installed globally:

   ```
   npm install -g nodemon
   ```

2. Run the application using the following command:
   ```
   nodemon --exec go run main.go --signal SIGTERM
   ```

This command will start the server and automatically restart it whenever you make changes to your Go files.
