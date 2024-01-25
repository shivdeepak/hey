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

1. Install go lang from the [official go website](https://go.dev/).

2. Clone this repository, and cd in to the project directory

```
git clone git@github.com:shivdeepak/hey.git
cd hey
```

3. Build the project. This will create `hey` executable in the `./bin` directory

```
make build
```

4. Configure and use `hey`

```
./bin/hey config init

# above command will ask for OpenAI API Key
```

5. Ask away your questions

```
./bin/hey why did the chicken cross the road
# note that special characters will confuse the script because
# your shell `bash` or `zsh` etc have special meanings for
# special symbols like ?, or words starting with `-` or `--`
# If you have special characters, put the entire prompt
# within double quotes
```

6. Install `hey` to your search path

This can be done in several ways:

6.1. Add project's `./bin` directory to `$PATH` environment variable in `.zshrc`, or `.bashrc`

For Example, add a line similar to the following at bottom of your shell's `.*rc` file.

```
PATH=$PATH:$HOME/Workspace/hey/bin
```

- Copy or Symlink `./bin/hey` to one of the directories in your search path.

For Example, you may want to run something similar to one of the two following commands
```
cp ./bin/hey /usr/bin/
# or
ln -s $HOME/Workspace/hey/bin/hey /usr/bin/hey
```
