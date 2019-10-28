# Cricket T20

Cricket T20 app simulate a running cricket match. It uses a weighted random number generation based on probability to determine the runs scored per ball.

## Table of Contents

  - [Installation](#install)
  - [Clean](#clean)
  - [Test](#test)
  - [Run](#run)
  - [Structure](#structure)

### Installation <a name="install"></a>

Clone this repo. And run:

```sh
make all
```

This command run everything in single shot (i.e clean, test, lint, run). To run each step seperatly execute the commands below one by one.

### Clean <a name="clean"></a>

To clean up all the generated artifacts, run:

```sh
make clean
```

### Test <a name="test"></a>

To test coverage, run:

```sh
make test
```

To see test-coverage details check test-coverage.html 

### Run <a name="test"></a>

To run app, run:

```sh
make run
```

This command runs app and prints out commentry and final result on commandline.

## Structure <a name="structure"></a>

```
.
├── Makefile
├── README.md
├── config
│   └── bengaluru.yaml
├── game
│   ├── config
│   │   └── bengaluru.yaml (test config)
│   ├── game.go
│   └── game_test.go
├── go.mod
├── go.sum
├── main.go
├── match
│   └── team.go
├── random
│   ├── random.go
│   └── random_test.go
└── start
    └── start.go
```
