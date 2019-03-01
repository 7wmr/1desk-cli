# 1Desk Cli


## Build

Build tool:

```bash
go build -o 1desk
```

Clone repo:

```bash
git clone git@github.com:ipsoft-tools/1desk-cli.git ~/go/src/1desk -b develop
```

Link to go src path:

```bash
ln -s ~/go/src/1desk ~/repos/1desk-cli
```

Link config file to home dir:

```bash
ln -s ~/repos/1desk-cli/.1desk.yaml ~/.1desk.yaml
```

## Deploy

```bash
brew tap ipsoft-tools/homebrew-tools
```

```bash
brew install 1desk-cli
```