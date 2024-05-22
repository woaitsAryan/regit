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
<p>Regit currently supports 5 commands:</p>
<ol>
  <li><code>regit recommit</code>: Reads all the commit diffs and writes better commit messages, then commits them again.</li>
  <li><code>regit own</code>: Makes you the author of all the commits.</li>
  <li><code>regit blame &lt;name&gt; &lt;email&gt;</code>: Makes the user specified by <code>&lt;name&gt;</code> and <code>&lt;email&gt;</code> the author of all the commits.</li>
  <li><code>regit nuke /path/to/file</code>: Removes the file specified from all the commits in the repository.</li>
  <li><code>regit retime &lt;duration&gt;</code>: Rewrites the commit times of all the commits in the repository to be of <code>x</code> hours in the past, evenly spaced. Can be any number of hours.</li>
</ol>
<p><strong>--path</strong> flag can be used to specify the path of the repository. If not specified, the current directory is used.</p>


## Credits
Regit uses [git-filter-repo](https://github.com/newren/git-filter-repo) under the hood to rewrite git histories. I would like to thank the authors and contributors of git-filter-repo for their work. 
