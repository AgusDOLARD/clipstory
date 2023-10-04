# Clipstory

ClipStory is a simple app for managing clipboard history.

## About

It utilizes a sqlite database to store clipboard entries and provides a simple command-line interface for accessing and viewing the clipboard history.

## Installation

To install and run ClipStory, follow these steps:

```sh
git clone https://github.com/AgusDOLARD/clipstory.git
cd clipstory
make install
```

## Usage

After installation, you can run the Clipstory application to retrieve and display your clipboard history.

1. Run the deamon using the command.

```sh
clipstoryd
```

2. Now you can view the clipboard history using the command.

```sh
clipstory
```

3. You can pipe the `clipstory` output to some other program/s to select an entry and copy it to your clipboard.

```sh
clipstory | fzf | xsel -ib
```
