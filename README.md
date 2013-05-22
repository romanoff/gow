Go Watch
========
Go watch is simple command line tool intended to monitor file events in specified folder and execute console commands appropriately.

#### Config example:
```toml
[rules]
  [rules.sass]
  pattern="*.sass,*.css"
  command="make regenerate_css"

  [rules.js]
  path="/path/to/folder/with/javascripts"
  pattern="*.js"
  command="make regenerate_js"
```

Config file has to be in `.gow` file located in current folder.

Example config file is watching for changes in .sass and .css files (in current folder) and if those are changes, executes "make regenerate_css" console command. It also contains other monitoring rule (you can have as many as you want). It watches for changes in *.js files in /path/to/folder/with/javascripts folder (path parameter overrides current path used by default) and executes "make regenerate_js" command.
