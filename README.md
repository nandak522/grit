# What buildout does
It has a ini syntax driven config file and that is used to bootstrap the whole environment.
Versions file is also the basis. The whole dependency management is done by [parts].
Each [part] contains a recipe and certain options required for the recipe. And each [part]
will either be a python package(setup.py driven) or a source package(CMMI driven) most of the time. If a package’s setup.py has got a dependency mentioned, its compared with the versions file and will be pulled from the index server accordingly.

Reasons why moving away from Buildout:

* Time taking
* Poor error reporting
* Editing configuration is a bit scary

# Why Go
* Speed
* Good exception handling
* Enables parallelism in a built-in way

## Lets call this new build tool.... Grit!
### How Grit works!
Fully ini driven. More like a task runner. Could be used for build activities. tasks are mentioned in the [ini] file.

```ini
[main]
index = <any pypi like index server>
pick_latest = true # default.
only_specced = true  # default = false. Opposite of pick_latest
extends = another_ini.ini
# versions are defaulted to a local versions file with [versions] content or any package(egg/cmmi) with hardcoded version specification
tasks = egg_task
        cmmi_task
# download cache is assumed to be $PWD/downloads unless explicitly mentioned

[dev_sources]
pkg_name = git git@github.com:user/repo.git   (This could be a python pkg or a cmmi package)

[egg_task]
runner = grit_egg # This downloads all pkgs and runs setup.py of each package with a custom python (So dependencies are looked up by easy_install or setuptools)
eggs = pyramid nose   # Will be pulled from the configured index server.

[cmmi_task]
runner = grit_cmmi # This downloads all pkgs and runs CMMI(configure-make-make-install) of each package with a custom prefix
pkgs = ftp://ftp.freetds.org/pub/freetds/stable/freetds-stable.tgz
```

So essentially grit_egg is more like a wrapper on top of easy_install/setuptools. Use logging to the bare minimum to start with. Have `npm` style colorful logging(With a tick(✓) mark when a task is completed). Figure out the `version conflicts` or even `failed to download` errors with a clear description of the dependency tree. Here are a few error examples(coloring will work similarly):

Failed to download zeromq because pkg server is not responding.
Dependency Tree: pkg1 ⇒ pkg2⇒ pkg3 ⇒ pkg4 ⇒ zeromq

Version conflict related to zeromq.
Required version ⇒ 2.2.0, Version mentioned in the [versions] ⇒ 2.5.0
Dependency Tree: pkg1 ⇒ pkg2⇒ pkg3 ⇒ pkg4 ⇒ zeromq


### Arsenal
```sh
$ grit boot      # Bootstraps the build
$ grit build     # Builds everything
$ grit versions  # Echoes all installed versions
$ grit shell     # similar to python shell or any project shell
```

### Installing Grit
It should work on the lines of `pythonbrew` or `pyenv`. Clone the repo and run a sh file, done! echo some bashrc snippet at the end, to be added to the user’s bashrc for hooking up the grit command when a fresh terminal is opened.

### Usage
```sh
$ git clone myproject.git
$ cd myproject
$ ls
build.ini versions.ini
$ grit build
Running Task egg_task ✓
Running Task egg_task ✓
All the tasks completed successfully :)
```

