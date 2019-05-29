# Emojipedia

Small program that scrapes unicode.org for Emoji data. Can be used to write out specific fragements of the page to be consumed from the command line. Things like categories, subcategories and emoji are all segemented into seperated folders that can be accessed from their respective command line hooks. The core package required for the sub-packages to be build is the main HTML content that is cached from unicode.org. This file is particularly big (namely due to several base64 SVG's existing in the source). After the desired packages are built for the program, this can be removed by running `emojipedia unicode [-r, remove]`. Keep in mind that if this package is gone and some of the other hooks aren't built yet, the other non-built packages won't work.

#### Get

`go get github.com/gellel/emojipedia`

#### Building

Before you can run any of the command line hooks, the program must first parse and store the HTML content from unicode.org. This can take awhile due to the size of the HTML document. Best to grab a snack and let it do its thing. Program still can access https://unicode.org/emoji/charts/emoji-list.html. If it breaks, chances are there's been a change to the URL or HTML.

`$ emojipedia unicode -b`

The program should output a status message if it succeeds parsing the page. The contents can be found in the hidden folder within your GoPath. 

## Hooks
The program supports several ways of accessing the content scraped from the unicode.org site. 

#### Categories
Categories explores the collection of category data fetched from the HTML. Prints out information at a high-level view of each found category. Assumes that all categories are stored on the disc, however specific categories can be removed from the category hook.
`$ emojipedia categories [options]`
#### Category
`$ emojipedia category [options]`
#### Emojipedia
`$ emojipedia emojipedia [options]`
#### Emoji
`$ emojipedia emoji [options]`
#### Subcategories 
`$ emojipedia subcategories [options]`
#### Subcategory
`$ emojipedia subcategory [options]`
