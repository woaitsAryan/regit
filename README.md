
# Regit

Regit is a friendly CLI tool written in Golang that allows you to rewrite entire git histories. With Regit, you can make yourself or anyone else the author of all the commits in a repository.

## Installation

To run Regit, you can either use the provided executable or build your own. 

To use the provided executable, simply run:

```bash
./regit help
```

To build your own executable, ensure that Go is installed and run:
    
```bash
make build
```

## Usage
Regit currently supports two commands:
1. `./regit own <path>`: This command makes you the author of all the commits in the repository at `<path>`. `<path>` should be the absolute path to the git repository, which can be obtained by running `pwd` in the root of the git repository.

Example:
```bash
./regit own /home/user/my-git-repo
```
2.`./regit blame <path> <name> <email>`: This command makes the user specified by `<name>` and `<email>` the author of all the commits in the repository at `<path>`.

Example:
```bash
./regit blame /home/user/my-git-repo "John Doe" "johndoe@example.com"
```

## Credits
Regit uses [git-filter-repo](https://github.com/newren/git-filter-repo) under the hood to rewrite git histories. I would like to thank the authors and contributors of git-filter-repo for their work. 