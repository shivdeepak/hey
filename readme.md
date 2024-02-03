# Hey

A friendly generative AI chatbot for your terminal

![screen recording](docs/hey_hello_world_pascal.gif)

## Features

1. Uses OpenAI ChatGPT Streaming API
    1. Save on monthly ChatGPT Plus subscription
    2. Spend 1¢/question instead of $20/mo
2. Uses minimal bandwidth compared to the ChatGPT website
    1. Perfect for road warrior
3. Ask multiple questions simultaneously
    1. Ask questions over multiple terminal tabs/sessions


## Usage

Currently, there are no binaries available to download, but 
it is easy to build and use. 

### 1. Install go lang from the [official go website](https://go.dev/).

### 2. Install `github.com/shivdeepak/hey`

```shell
go install github.com/shivdeepak/hey@latest
```

### 3. Add `$GOPATH/bin` to your search path

Check if $GOPATH is set in your shell (zsh or bash)

```shell
echo $GOPATH
```

If it's empty, then extract and set that value in `.zshrc`, or `.bashrc`

```shell
go env | grep GOPATH
#> GOPATH='/Users/<username>/go'
```

```shell
echo "GOPATH='/Users/$USER/go'" >> ~/.bashrc
```

Add `$GOPATH/bin` to `$PATH` in `.zshrc`, or `.bashrc` if it doesn't exist 

```shell
PATH=$PATH:$GOPATH/bin
```

### 4. Configure and use `hey`

```shell
hey config init

# above command will ask for OpenAI API Key
```

### 5. Ask away your questions

```shell
hey why did the chicken cross the road

# note that special characters will confuse the script because
# your shell `bash` or `zsh` etc have special meanings for
# special symbols like ?, or words starting with `-` or `--`
# If you have special characters, put the entire prompt
# within double quotes
```
