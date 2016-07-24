# Untertiteln

The console application prints Untertitels (subtitle in German) from input.

## Install

```zsh
% go get github.com/Masterminds/glide
% glide install

# setup environment variables (z.B. using autoenv etc.)
% cp .env.sample .env
```

## Requirements

* Google's [cloud speech api](https://cloud.google.com/speech/)
* [SoX](http://sox.sourceforge.net/)


## Run

```zsh
% cd /path/to/untertiteln
% go run ./main.go
```

## Link

* https://cloud.google.com/speech/docs/best-practices
* https://github.com/GoogleCloudPlatform/nodejs-docs-samples/tree/master/speech


## License

GNU GPL v3

Untertiteln
Copyright (c) 2016 Yasuhiro Asaka

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
