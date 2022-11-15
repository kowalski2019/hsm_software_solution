
# HSM (Hardware Security Model) Software Solution by Claude Stephane M. Kouame

In order to realise the encryption and decryption of the data or text, I offer here a solution on the software level of the HSM technology. This solution will be nothing more than a simple instance with an API whose task is to encode and decode.This instance or service can be hosted anywhere. The most important thing is that the configuration file that contains the key and other important elements like an authentication key is present.

This instance will run in a Docker container or a Linux service, both are possible. The most important thing is that you have a folder "/crypto/config/" on its server and in this folder a file called "config.toml", because without this file the application will not work.

The content of the configuration file should look like this:


```
[constants]
crypto_key = "4bda55a55fcc2678e72ba07c14f3eeb3e59a6fc523fc6ae594d232762dbcd8dd"
auth_key = "dummy"
```

**crypto_key** is the key used for encryption and decryption.

**auth_key** is the key that authenticates the person making a request to encrypt or decrypt a text.

**! Important: Most of the commands executed in this tutorial are suitable for a Linux environment, a small adaptation is required for a Windows environment.**

## Pull docker image

```
sudo docker pull csmk59/crypto:latest
```

## Create the configuration folder

```
sudo mkdir -p /crypto/config
```

## Create and edit the configuration file "config.toml".

Editing the file can be done with nano or vim, the choice is yours ;)

```
touch /crypto/config/config.toml
vim /crypto/config/config.toml
```

You add the various information as in the example configuration above.

It is essential that the encryption and decryption key is at least 32 bytes in size (32 characters, each character contains 8 bits, so 1 byte), as the encryption technique is AES-256.

## Start the Docker container

```
sudo docker run -d -p <host_port>:8008 -v /crypto/config:/crypto/config csmk59/crypto:latest
```

<host_port> must be chosen by you. It can simply be 8008 or 8000 etc....

## Test API

The test can be done with postman or curl.

### Example of encryption

```
 curl -d 'auth_key=Dummy' \
        -d 'text=Hello World!' \
        -d 'crypto_type=encode' \
        -X POST \
        -H 'Content-Type: application/xwww-form-urlencoded' \
         http://127.0.0.1:<host_port>/api/v1/encryption

```

### Example of decryption

```
curl -d 'auth_key=dummy' \
        -d 'text=hewfj3jh3981239832bj09djn' \
        -d 'crypto_type=decode' \
        -X POST \
        -H 'Content-Type: application/xwww-form-urlencoded' \
         http://127.0.0.1:<host_port>/api/v1/encryption


```


