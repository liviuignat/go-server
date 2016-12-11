#### Simple Go server with Martini

##### Go Setup
```sh
mkdir -p $HOME/go
mkdir -p $HOME/go/bin
mkdir -p $HOME/go/src
mkdir -p $HOME/go/pkg
export GOPATH=$GOPATH:$HOME/go
export PATH=$PATH:$HOME/go/bin
cd $HOME/go/src
git clone https://github.com/liviuignat/go-server.git
cd go-server
```

##### Setup
```sh
glide install # install dependecies
gin           # run app
```

##### Micro Docker Image build on MacOS
```sh
chmod +x ./tasks/build
./tasks/build
# https://sebest.github.io/post/create-a-small-docker-image-for-a-golang-binary/
# at this point the docker image has max 10 MB
docker run -d -p 3000:3000 --name go-server go-server
```

Visit http://localhost:3000
