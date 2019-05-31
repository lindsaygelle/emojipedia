# Emojipedia

A simple, small program that scrubs the unicode.org site for emoji data. Can be used to write out specific fragments of the page to be consumed from the command line or something else. Things like emoji categories, emoji subcategories and emoji are all segmented into seperated folders that can be accessed from their respective command line routines. The core package required for the sub-packages to be build is the main HTML content that is cached from unicode.org. This file is particularly big (namely due to several base64 SVG's existing in the source). After the desired packages are built for the program, this can be removed by running the remove command. Keep in mind that if the unicode.org HTML package is gone and some of the other hooks aren't yet build, the other non-built packages won't work and the content will need to be fetched again.

## Get

Installing the package is as easy as running Go's get command.

`$ go get github.com/gellel/emojipedia`

## Example

Here's an example of listing out the categories using the program. 

```
$ clear && emojipedia [-c|categories] [-l|list]

Name                    |Number |Emoji  |Subcategories
activities              |6      |79     |5
animals-and-nature      |3      |127    |8
component               |2      |4      |1
flags                   |9      |268    |3
food-and-drink          |4      |121    |8
objects                 |7      |233    |18
people-and-body         |1      |311    |16
smileys-and-emotion     |0      |149    |15
symbols                 |8      |217    |11
travel-and-places       |5      |210    |11
```
Or perhaps you'd like the emoji.

```
$ clear && emojipedia [-e|emojipedia] [-l|list]

Name                                            |Number |Category               |Subcategory            |Keywords
1st-place-medal                                 |947    |activities             |award-medal            |4
2nd-place-medal                                 |948    |activities             |award-medal            |4
3rd-place-medal                                 |949    |activities             |award-medal            |4
a-button-blood-type                             |1384   |symbols                |alphanum               |3
ab-button-blood-type                            |1385   |symbols                |alphanum               |3
abacus                                          |1088   |objects                |computer               |2
adhesive-bandage                                |1212   |objects                |medical                |2
admission-tickets                               |942    |activities             |event                  |3
aerial-tramway                                  |839    |travel-and-places      |transport-air          |5
airplane                                        |830    |travel-and-places      |transport-air          |2
airplane-arrival                                |833    |travel-and-places      |transport-air          |6
airplane-departure                              |832    |travel-and-places      |transport-air          |5
... 
```

And of course, the emoji in detail.

```
clear && emojipedia [-ee|emoji] boar [-t|table]

        |category               |codes  |href            |keywords       |name   |number |subcategory
🐗     |animals-and-nature     |U+1F417|http://....     |boar pig       |boar   |494    |animal-mammal

```

## Building

Before you can run any of the command line hooks, the program must first parse and store the HTML content from unicode.org. This can take awhile due to the size of the HTML document. Best to grab a snack and let it do its thing. Program still can access https://unicode.org/emoji/charts/emoji-list.html. If it breaks, chances are there's been a change to the URL or HTML.

`$ emojipedia unicode [-b | build]`

The program should output a status message if it succeeds fetches and downloads the page. The unicode HTML document can be found within the .emojipedia folder within your bundle inside the program GoPath. 

## Full install

This command will _"quick"_ install all of the separate packages. Futher update will compress this into a single command.

`$ emojipedia unicode -b && emojipedia categories -b && emojipedia emojipedia -b && emojipedia subcategories -b && emojipedia unicode -r`

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


## Category commands

```$ emojipedia category <name> [-a | anchor]```
```$ emojipedia category <name> [-e | emoji]```
```$ emojipedia category <name> [-h | href]```
