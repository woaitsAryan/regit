<br />
<div align="center">
    <img src="assets/icon_transparent.png" alt="Logo" width="150" height="150">

  <h1 align="center">Regit</h3>

  <p align="center">
    CLI tool to manage git repositories and histories
    <br />
    <br />
    <a href="#installation">Installation</a>
    ·
    <a href="#docs">Docs</a>
    ·
    <a href="https://github.com/woaitsAryan/regit/issues/new?labels=enhancement&template=feature-request---.md">Request Features</a>
    ·
   <a href="https://github.com/woaitsAryan/regit/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
  </p>
</div>


<p align="center">
Regit is a CLI tool written in Go that allows you to rewrite git histories. Changing ownership, timestamps, even commit messages to follow conventions, all in a single command
</p>

## Warning
Using regit multiple times might cause object corruption in your git repository. Please use it after backup up your .git folder.

<h2 name="installation">Installation</h2>

<h3>Windows</h3>

<pre><code>python3 -m pip install --user git-filter-repo</pre></code>
<pre><code>winget install regit</code></pre>

<h3>macOS</h3>

<pre><code>git clone https://github.com/woaitsAryan/regit && cd regit</pre></code>
<pre><code>make setup</code></pre>

<h3>Linux</h3>
<pre><code>git clone https://github.com/woaitsAryan/regit && cd regit</pre></code>
<pre><code>make setup</code></pre>

<h2 name="docs">Docs</h2>

Regit currently supports 5 commands:

1. `regit recommit`: Reads all the commit diffs and writes better commit messages, then commits them again.
2. `regit own`: Makes you the author of all the commits.
3. `regit blame <name> <email>`: Makes the user specified by `<name>` and `<email>` the author of all the commits.
4. `regit nuke /path/to/file`: Removes the file specified from all the commits in the repository.
5. `regit retime <duration>`: Rewrites the commit times of all the commits in the repository to be of `x` hours in the past, evenly spaced. Can be any number of hours.
6. `regit rewind <duration>`: Rewinds the commit times of all the commits to be pulled `x` hours in the past. Can be any number of hours.
7. `regit fastforward <duration>`: Fast forwards the commit times of all the commits to be pushed `x` hours in the future. Can be any number of hours.

## Common Flags

The following flags can be used with any command:

- `--source` or `-s`: Specify the path to the git repo. If not specified, the current directory is used.
- `--branch` or `-b`: Specify a branch.
- `--verbose` or `-v`: Enable verbose output. This will print additional details about the operations being performed.
- `--quiet` or `-q`: Enable quiet output. This will suppress most output, printing only essential information.

## Credits
Regit uses [git-filter-repo](https://github.com/newren/git-filter-repo) under the hood to rewrite git histories. I would like to thank the authors and contributors of git-filter-repo for their work. 
