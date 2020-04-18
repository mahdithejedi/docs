# python modules and packages

## how import works

in import python search your modules in `sys.path` and if the modules you search was not found we face exception

so how we can add a path to `sys.path`?

*   **[`sys.path.append`](https://docs.python.org/3/library/sys.html#sys.path)**

eg:
`sys.path.append('MY_DIRECTORY')

* **[`PYTHONPATH`](https://docs.python.org/3/using/cmdline.html#envvar-PYTHONPATH)**

you can say to linux to add the path to python!

eg: `export PYTHONPATH=MY_DIRECTORY`
or 
[`export PYTHONPATH="${PYTHONPATH}:/my/other/path"`](https://stackoverflow.com/a/3402176/9651641)

or
the most efficent way is

[`export PYTHONPATH="$PWD/MY_DIRECTORY"`](https://bic-berkeley.github.io/psych-214-fall-2016/using_pythonpath.html)

and to make it permenent you can add it to ~/.bashrc

## execute a directory

with `__main__.py` you can make a python package  executable  like you can get argument from command line and use it in hole package 

**even** you can create a zip file and tell python to execute that zip file and you can like add a setup script and say to python in \__main__.py file to run that script!
