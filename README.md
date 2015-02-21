# rightFootLeftFoot_Go
Identifies unused localization key/value pairs in a project.
This is a variation on a small utility program called CommandL10n. Recursively reads the directory and its sub-directories looking for view files, eventually comparing localization string variables from the view files with the locale file key/value pairs
I used GoLang for this version. I want to re-write this simple utility in a few different languages to help me gauge my interest in learning more about each of them.

### example of returned array of unused localization string keys (  returns the unused localization keys only )
```javascript
['started']
['exported']
['starting']
['demand']
['delivered']
['planned']
['scheduled']
```

### Running from the command line
```javascript
	> .\getunusedL10n -localepath="path/to/your/en.json" -viewpath="path/to/your/view/file"
```

##### Options:
```javascript
	-v, --viewfolder     [string] - path to your view file
	-l, --localefile     [string] - path and file name to locale file you want to read
```
