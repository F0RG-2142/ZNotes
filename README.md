# ZNotes
The Official Repository for the ZNotes App<br>

The induvidual repos can be visited at:<br>
https://github.com/F0RG-2142/znotes-frontend-2<br>
https://github.com/F0RG-2142/znotes-backend<br>
https://github.com/F0RG-2142/LogVanguard (a log analyser for the future)<br>

I will be creating a local installer in the future for self hosting.<br>
And now, the _flavor text_:<br>

## Why I decided to make this

## Current WIP Installation Guide

Create the main directory, clone the two repos, in the backend directory install dependancies and run with Go 
```bash
mkdir znotes
cd znotes
git clone https://github.com/F0RG-2142/znotes-backend@latest
git clone https://github.com/F0RG-2142/znotes-frontend@latest
cd znotes-backend
go install github.com/pressly/goose/v3/cmd/goose@latest
go run .
```
And in a new terminal
```bash
cd znotes/znotes-frontend
npm i
npm run dev
```
If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
