# sumaxmlrpc - written in Go
This application has functions which can:
* list active systems
* list available packages of systems
* schedule upgrade jobs
* update GPG or SSL content defined in SUSE Manager. (Useful for RHEL users to update rotating certs)

This program uses config file in which the server, userid and password are stored for authenticating to SUSE Manager.
The password will be encrypted. To generate a encrypted password in the config file follow below steps:

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
Run below command to start generating the sumaconfig file.
```
sumaxmlrpc -create-sumaconf sumaconfig.yaml
```

## Download the binary: Change the Url to the release you want to download.
```
wget https://github.com/bjin01/sumaxmlrpc/releases/latest/download/sumaxmlrpc_Linux_x86_64.tar.gz
tar -xvf sumaxmlrpc_Linux_x86_64.tar.gz
sudo cp sumaxmlrpc /usr/local/bin
```

## Usage:
Run the binary with -h for help and see the subcommands and arguments to continue.
The binary can be executed as non-sudo user.

```
sumaxmlrpc --help
```
__For instance:__
Update SSL definition in SUSE Manager:
```
sumaxmlrpc -sumaconf sumaconfig.yaml gpgsslupdate -file /home/bjin/temp/testgpg.key -description "testkey" -type "SSL"
```
