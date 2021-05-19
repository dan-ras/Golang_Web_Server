/**********************************************************
* Author: Daniel Rasmussen
*
* 	Hrs: 12
*
* Program Description:
*	This program is a basic webserver for handling simple
*	requests. Includes both a implementation of the Go
*   networking library as well as a manual implementation.
*   Both have the same functionality in delivering webpages
*   to the client browser.
*
**********************************************************/
package main // Used for the compilation of main()

// Import the necessary libraries.
import (
	"fmt"      // Printing to the console.
	"io"       // Input/Output
	"net"      // Used for the manual implementation
	"net/http" // Used for http. functions.
	"os"       // Command-line arguments.
	"strings"  // Strings.
)

/**********************************************************
* Handles the specific instance when a user requests the
* '/textfile.txt' file.
**********************************************************/
func textFunc(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "403: That page is down for maintenance.", http.StatusForbidden)
	return
}

/**********************************************************
* Prints a usage statement specifying the startup command.
**********************************************************/
func printUsage() {
	println("\nFormat for command to start the server program is: \n")
	println("\t'go run server.go [serverPort#] [serverSelector]'")
	println("\nPlease enter one of the following as the serverSelector:\n")
	println("\t'1' for network library.")
	println("\t'2' for a manual implementation.")
	println("\nRecommended serverPort number is '80'")
}

/**********************************************************
* Checks the last three chars of the file extension.
* Returns the content type for the TCP header.
**********************************************************/
func fileExtHandler(message string) string {
	// Create a map of valid extensions.
	extensions := map[string]string{
		".txt": "text/txt;",
		".gif": "image/gif;",
		".jpg": "image/jpeg;",
		"html": "text/html;",
	}

	// Get the file extension
	fileExt := message[len(message)-4:]
	//var fileExt string = message[len(message)-4:]

	// Specify encoding type:
	encodeType := " charset=UTF-8"

	// Check for a file with the specified extension.
	if contentType, exists := extensions[fileExt]; exists {
		return contentType + encodeType
	} else {
		return "Invalid"
	}
}

/**********************************************************
* Handles page requests from the user. This is what a
* manual implementation of the 'http.Handle() could look
* like.
**********************************************************/
func sendFileToClient(connection net.Conn) {
	fmt.Println("\nA client has connected!")

	// Close the connection on any return statement.
	defer connection.Close()

	// Get the filename from the GET request.
	b := make([]byte, 2048)
	length, err := connection.Read(b)
	if err != nil {
		println(err)
	}
	fullRequest := string(b[:length])
	clientRequest := strings.Split(fullRequest, " ")
	//fileName := strings.TrimPrefix(clientRequest[1], "/")
	fileName := clientRequest[1]

	// Declare a 404 message with http line feed.
	error404 := "404: File not found.\r\n\r\n"
	if fileName == "" {
		connection.Write([]byte(error404))
		println("Unable to process empty requests.\n")
		return
	}

	// Print out our client's GET request.
	println("Client request: ", clientRequest[0]+" "+clientRequest[1])

	// Verify the file extension.
	FileContentType := fileExtHandler(fileName)
	if FileContentType == "Invalid" {
		connection.Write([]byte(error404))
		println("Invalid content type.\n")
		return
	}

	// Send the header.
	fmt.Println("Sending header...")
	Header := "HTTP/1.1 200 OK\r\n"
	ConType := "Content-Type: " + FileContentType + "\r\n" + "\r\n"
	connection.Write([]byte(Header))
	connection.Write([]byte(ConType))

	// Send the file. First open it and check to make sure it exists:
	file, err := os.Open("files" + fileName)
	if err != nil {
		connection.Write([]byte(error404))
		fmt.Println(err)
		return
	}
	sendBuffer := make([]byte, 1024) // Create a sending buffer.
	fmt.Println("Sending file...")
	for { // Send the file one buffer at a time.
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	fmt.Println("File has been sent, closing connection!")
	return // Defer statement from line 84 is executed.
}

/**********************************************************
* This is main.
*
*	Main contains two different implementations of a web
*	server in Go. One uses the networking library while
*   the other one was written by me.
**********************************************************/
func main() {
	// Grab the number of command-line arguments.
	argLength := len(os.Args[1:])

	// Check to see which server we are running.
	if argLength == 2 {
		serverPort := os.Args[1]

		if os.Args[2] == "1" {
			// Easy way (with library)
			//////////////////////////////////////
			fmt.Printf("\n\tStarting library server on port %s.\n\n", serverPort)
			http.Handle("/", http.FileServer(http.Dir("./files"))) // Handle all requests except 'textfile.txt'
			http.HandleFunc("/textfile.txt", textFunc)
			// Set up a socket
			// The line below is a loop.
			if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
				fmt.Println(err)
			}
		} else if os.Args[2] == "2" {
			// Hard way (manual)
			///////////////////////////////////////
			// Create a server and listen on port 80.
			server, err := net.Listen("tcp", ":"+serverPort)
			if err != nil {
				fmt.Println("Error listening: ", err)
				os.Exit(1)
			}

			// Close server port when finished executing.
			defer server.Close()

			// Continuously serve client requests.
			fmt.Printf("\n\tStarting manual server on port %s.\n\n", serverPort)
			for {
				connection, err := server.Accept()
				if err != nil {
					fmt.Println("Error: ", err)
					os.Exit(1)
				}
				go sendFileToClient(connection) // Go run this function.
			}
		} else {
			printUsage()
		}
	} else { // Default case; print a usage statement.
		printUsage()
	}
}
