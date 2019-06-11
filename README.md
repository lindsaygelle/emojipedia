# Emojipedia

A simple, small program that scrubs the unicode.org site for emoji data. Can be used to write out specific fragments of the page to be consumed from the command line or something else. Things like emoji categories, emoji subcategories and emoji are all segmented into seperated folders that can be accessed from their respective command line routines. The core package required for the sub-packages to be build is the main HTML content that is cached from unicode.org. This file is particularly big (namely due to several base64 SVG's existing in the source). After the desired packages are built for the program, this can be removed by running the remove command. Keep in mind that if the unicode.org HTML package is gone and some of the other hooks aren't yet build, the other non-built packages won't work and the content will need to be fetched again.

## Install

Grabbing the program is as easy as running Go's `get` command. After everythings collected, you'll probably want to run the `go build` command to create a binary. After that, you're almost good to go.

```go get github.com/gellel/emojipedia && go build```

## Building

Assuming you've created the binary, you'll next want to run the unicode `build` command. This option attempts to request and store the current unicode.org chart to the programs local folder. As mentioned above, this site is pretty big, so you might want to queue up something else to do while everything downloads.

```emojipedia [-u unicode] [-b build]```

As of writing this documentation, the program assumes that the source content is still hosted under the URL https://unicode.org/emoji/charts/emoji-list.html. Should this page be moved, removed or auth protected, chances are the program will not work. If this is the case, please raise a issue. Otherwise, the program should just download and store the file (eventually).

## Packages

The emojipedia program separates the contents of the unicode.org HTML file in several different subsets. Given the amount of content that is contained at each level, the emojipedia program does not automatically create each and every one for you on install. To create a new package, run the `build` command for the content desired. Currently, there are four main package directories that can be built out of HTML file. These are `categories`, `emojipedia`, `keywords` and `subcategories`. Each of these can be built individually and are not interdepenant, but all require the unicode.org HTML file to exists before they can be created.

```emojipedia [<package-name>] [-b build]```

After you have created the desired number of packages, you can removed the unicode.org HTML file.

```emojipedia [-u unicode] [-r remove]```


## Usage (collections)

The program will show you a set of available commands if no input is given at runtime. However, here are some basic commands to help get you started.

Each of the main packages can be browsed using a common set of commands. These are `get`, `keys`, `list`, and `number`. 

```emojipedia [<[-c],[-k],[-e],[-s]>] [<[-g get],[-k keys],[-l list],[-n number]>]```

Here's an example of using the list command on the categories package.

```
emojipedia [-c categories] [-l list] [<options>]

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

And another; in this example we're listing out the emoji, but using the same command.

```
emojipedia [-e emojipedia] [-l list] [<options>]

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

Using the `get` command lets you specify the specific package contents you'd like to list. It takes _n_ positional arguments after the keyword `g` or `get`. Each argument refers to the name of a corresponding file. This command is similar to the `list` variant, but offers more detail than its sibling.

```
emojipedia [-c categories] [-g get] flags symbols
...

```

## Usage (specifics)

For most of the main packages, there also exists a detail level that enables browsing a particular package element in more detail. These detail routines require the main package to exists. This means that if you have built the `categories` package, you also have access to the `category` program. For each of these detail specific programs, there are a few common commands that exist. These generally include `anchor`, `category`, `emoji`, `href`, `number`, `position`, `subcategory` and `table`. Some programs, like the `emoji` routine have more options than are listed here. For more run the emoji command.

Here's an example or running the shared `table` command for the emoji detail routine.

```
emojipedia [-ee emoji] boar [-t table] [<options>]
...
🐗     |animals-and-nature     |U+1F417|http://....     |boar pig       |boar   |494    |animal-mammal

```

