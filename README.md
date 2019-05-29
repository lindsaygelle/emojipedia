# Emojipedia

Small program that scrapes unicode.org for Emoji data. Can be used to write out specific fragements of the page to be consumed from the command line. Things like categories, subcategories and emoji are all segemented into seperated folders that can be accessed from their respective command line hooks. The core package required for the sub-packages to be build is the main HTML content that is cached from unicode.org. This file is particularly big (namely due to several base64 SVG's existing in the source). After the desired packages are built for the program, this can be removed by running `$ emojipedia unicode [-r | remove]`. Keep in mind that if this package is gone and some of the other hooks aren't built yet, the other non-built packages won't work.

## Get

`$ go get github.com/gellel/emojipedia`

## Example

Here's an example of listing out the categories using the program. 

```
$ clear && emojipedia categories -l

activities               6       5       79
animals-and-nature       3       8       127
component                2       1       4
flags                    9       3       268
food-and-drink           4       8       121
objects                  7       18      233
people-and-body          1       16      311
smileys-and-emotion      0       15      149
symbols                  8       11      217
travel-and-places        5       11      210
```

## Building

Before you can run any of the command line hooks, the program must first parse and store the HTML content from unicode.org. This can take awhile due to the size of the HTML document. Best to grab a snack and let it do its thing. Program still can access https://unicode.org/emoji/charts/emoji-list.html. If it breaks, chances are there's been a change to the URL or HTML.

`$ emojipedia unicode [-b | build]`

The program should output a status message if it succeeds fetches and downloads the page. The unicode HTML document can be found within the .emojipedia folder within your bundle inside the program GoPath. 

## Common commands
The program supports several ways of accessing the content scraped from the unicode.org site. There are a number of common commands that are shared amongst the top-level programs, these the `categories`, `emojipedia`, `subcategories`.
 
##### Building packages

```$ emojipedia <package> [-b | build]```

Builds a specific bundle. Requires the unicode HTML to be on disc. Program cannot run the bundle routines without the package being generated. Run this command beforehand so that the appropriate hooks can be explored. If missing, the program will raise a missing file status and exit.

##### Getting specific element from package

```$ emojipedia <package> [-g | get]```

Fetches a specific element from the accessed bundle (if it exists). Prints out the top-level detail for the accessed element. Is more verbose than the list command, but less-so than accessing the element directly.

##### Fetching keys from package

```$ emojipedia <package> [-k | keys]```

Lists out the available options that can be picked from. Is a fast way for looking up a collection of elements for an accessed package. The keys are sorted alphabetically for convienence. 

##### Listing elements from a package

```$ emojipedia <package> [-l | list]```

Lists all the elements and prints out top-level detail. Less verbose than `-g | get`.

##### Removing a package

```$ emojipedia <package>  [-r | remove]```

Removes all of the bundle for the selected program. Does not delete the unicode HTML file.
