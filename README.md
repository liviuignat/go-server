#### Simple Go server with Martini

##### Go Setup
```sh
mkdir -p $HOME/go
mkdir -p $HOME/go/bin
mkdir -p $HOME/go/src
mkdir -p $HOME/go/pkg
export GOPATH=$GOPATH:$HOME/go:$HOME/go/src/go-server
export PATH=$PATH:$HOME/go/bin:$HOME/go/src/go-server/bin
cd $HOME/go/src
git clone https://github.com/liviuignat/go-server.git
cd $HOME/go/src/go-server
```

##### Setup
```sh
# Setup Postgres for local environment
docker run -d \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_USER=postgres \
  --volume $HOME/docker/postgres/data:/var/lib/postgresql/data \
  --name postgres \
  postgres

# Install package manager
brew install glide
go get github.com/codegangsta/gin

# Install dependencies
glide install # install dependencies
gin           # run app
```

##### Docker deployment
```sh
docker build -t go-server .
docker run -d -p 3000:3000 --name go-server go-server
```

Visit http://localhost:3000
