# sumaxmlrpc - written in Go
This application has functions which can:
list active systems
list available packages of systems
schedule upgrade jobs
update GPG or SSL content defined in SUSE Manager.

This program needs a SUSE Manager login config file in which the server, userid and password are stored.
The password will be encrypted. Use generate a encrypted password in the config file follow steps below:

## Use your own keystring for password encryption:
Generate a 32 bytes keystring on linux:
```
head -c16 </dev/urandom|xxd -p -u
```
Set the OS environment variable:
```
export SUMAPROM_ENCRYPT=`head -c16 </dev/urandom|xxd -p -u`
```
Or use one-liner command:
```
export SUMAPROM_ENCRYPT=`head -c16 </dev/urandom|xxd -p -u`
```

## Generate sumaconfig:
```
go run main.go -create-sumaconf suma.conf
```

## Download the binary: Change the Url to the release you want to download.
```
wget https://github.com/bjin01/sumaxmlrpc/releases/download/v0.1.2/sumaxmlrpc_Linux_x86_64.tar.gz
tar -xvf sumaxmlrpc_Linux_x86_64.tar.gz
sudo cp sumaxmlrpc /usr/local/bin
```

## Usage:
Run the binary with -h for help and see the subcommands and arguments to continue.
```
sumaxmlrpc --help
```
__For instance:__
Update SSL definition in SUSE Manager:
```
sumaxmlrpc -sumaconf sumaconfg.yaml gpgsslupdate -file /home/bjin/temp/testgpg.key -description "testkey" -type "SSL"
```