# KLED-Server

## Motivation

The abbreviation KLED stands for "Korean Learner's English Dictionary" which is based on [Korean Leaner's Dictionary](https://krdict.korean.go.kr/mainAction). The dictionary even provides an extensive API which can be used after registering account on their website.

I found two particular problems with their API that bothered me a lot though:
1. For some reason the requested multimedia part (videos, pronunciation files etc.) is not delivered.
2. The searched keyword is not highlighted in any way in the example sentences.[^1]

For one of my other [projects](https://github.com/Mxngls/kled-scraper) I wrote an API endpoint that parses the html for a given word. For future use I thought it might be great to to provide an API endpoint to be able to get single dictionary entries instead of the whole dict.

## Description

The given script does exactly that. As the dictionary itself is licensed under the Creative Common License there are no copyright issues to worry about at all.


## Installation & Usage

Given that Go is installed:
```zsh
% go version go1.17.6 darwin/amd64
```

Just clone the repository:
```zsh
% git clone git@github.com:Mxngls/kled-server.git
````

Then run:
```zsh
% go run .
```

This will `tart a server that runs on localhost and accepts two different kind of requests under the following paths:
1. ```/search```
2. ```/view```

The result for either one will a http response with the specific data encoded in JSON.

[^1]:To see why this might be a problem see the various possible conjugations of the *regular* verb [건네다](https://en.wiktionary.org/wiki/%EA%B1%B4%EB%84%A4%EB%8B%A4#Conjugation).