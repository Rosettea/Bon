# Bon
> 🍮  A small, simple URL shortener.

Bon is a very tiny and simple self hosted URL shortener written in Go.  
![](https://modeus.is-inside.me/geC6rUKI.png)

# Install
## With go get
```
go get -u github.com/Rosettea/Bon
```

## Compile from Source
```
git clone https://github.com/Rosettea/Bon
cd Bon
go build
```

## Using Docker

Build:
```
git clone https://github.com/Rosettea/Bon
cd Bon
docker build -t bon:latest .
```
Run:
```
docker run -it -p 3000:3000 bon:latest
```

# Usage
Run the `Bon` executable, then it'll be running at port 3000.

# License
BSD 3-Clause

