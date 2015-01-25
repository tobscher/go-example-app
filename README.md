# go-example-app

## Test deployment via kiss

Start [vagrant](https://www.vagrantup.com/) box:

```
$ vagrant up
```

Download kiss:

```
$ go get github.com/gophergala/go_ne/kiss
```

Install required plugins:

```
$ go get github.com/gophergala/go_ne/plugins/plugin-git-clone
$ go get github.com/gophergala/go_ne/plugins/plugin-whoami
$ go get github.com/gophergala/go_ne/plugins/plugin-env
```

Run deployment task `setup`:

```
$ kiss -host=127.0.0.1 -username=vagrant -password=vagrant -port=2222 -task=setup
```

Deploy/update your application:

```
$ kiss -host=127.0.0.1 -username=vagrant -password=vagrant -port=2222 -task=deploy
```

NOTE: This will pull the latest changes from GitHub. Feel free to fork this repo and try updating the code.
