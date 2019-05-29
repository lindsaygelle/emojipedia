# Emojipedia

Small program that scrapes unicode.org for Emoji data. Can be used to write out specific fragements of the page to be consumed from the command line. Things like categories, subcategories and emoji are all segemented into seperated folders that can be accessed from their respective command line hooks. The core package required for the sub-packages to be build is the main HTML content that is cached from unicode.org. This file is particularly big (namely due to several base64 SVG's existing in the source). After the desired packages are built for the program, this can be removed by running `emojipedia unicode [-r, remove]`. Keep in mind that if this package is gone and some of the other hooks aren't built yet, the other non-built packages won't work.

#### Get

`go get github.com/gellel/emojipedia`

#### Building

`$ emojipedia unicode -b`
