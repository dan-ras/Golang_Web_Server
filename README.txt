/**********************************************************
* Author: Daniel Rasmussen
*
* Program Description:
*   This program is a basic web server for handling simple
*   requests. Includes both a implementation of the Go
*   networking library as well as a manual implementation.
*   Both have the same functionality in delivering webpages
*   to the client browser.
*
**********************************************************/

Instructions for execution:

    1. Must install Go on your machine.
        This is done by visiting "https://golang.org/doc/install" and downloading 
        Go for the desired operating system then installing it.


    2. Once installed, un-tar and copy the files to a directory. The package consists of the server.go
       file and a directory containing the some test files. These files
       need to be kept in the directory named "files". Navigate to this directory with the server.go file
       and the 'files' directory with a command prompt or terminal window.
    

    3. Run the server command. There is a usage statement that is printed when the server command is wrong.
       The command that needs to be entered is: 
       
       'go run server.go [serverPort#] [serverSelector]'

       Where the serverPort is an integer and the serverSelector is either '1' or '2'. A message
       stating the server is running will appear when the command is entered correctly.


    4. Open up a browser tab and type the name of the host, most likely this will be 'localhost'.

       ***The below sample command should work with both servers as long as the test files are left in
       the 'files' folder.***

            'localhost:[serverPort]/bufbomb.html'

