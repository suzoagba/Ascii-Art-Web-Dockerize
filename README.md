# Ascii Art Web Dockerize

## Description
This program starts a web server that creates ASCII art from user input. ASCII art is a form of art that uses ASCII characters to create images. 

The server can be started with the following command: `go run .`, or using the Docker image (see below).
The server starts on port 8080 and listens for requests on the following endpoints:
- `/` - the main page with a form for entering text
- `/ascii-art` - the page with the result of the conversion of the entered text to ASCII art

## Installation
To use this program, Docker must be installed on your system. If you don't have it installed, you can download it from the official Docker website.

Once Docker is installed, you can build and run the program using the following commands:

```docker build -t ascii-art .```

This command builds a Docker image called "ascii-art" using the Dockerfile provided in this repository.

```docker container run -p 8080:8080 ascii-art```

This command runs the "ascii-art" image in a Docker container, mapping port 8080 of the container to port 8080 of your host machine.

The server will start on port 8080. To access the website, open a web browser and navigate to http://localhost:8080.

### Bonus Dockerfile
A second Dockerfile is provided to build a Docker image that is optimized for size.
In this case, the image is built using a multi-stage build, and the final image is based on the `scratch` image. The command to build the image is the same as above, except that the dockerfile name must be specified:

```docker build -t ascii-art -f Dockerfile.bonus .```

## Website usage
To use the website, follow these steps:

1. Navigate to http://localhost:8080 in your web browser.
2. Select a banner among the available ones (radio buttons).
3. Enter the text to be converted to ASCII art (textarea). Note that the text can contain only  printable ASCII characters and newlines.
4. Click the "Submit" button to generate the ASCII art.

It is also possible to download the result of the conversion in plain text format by clicking on the "Download" button. A file will be downloaded, with a suggested name based on the text entered by the user.

## Authors
- Samuel Uzoagba (https://01.kood.tech/git/suzoagba)
- Jeremiah Bakere (https://01.kood.tech/git/googlee01)
- Jude Eze (https://01.kood.tech/git/Maxwell)