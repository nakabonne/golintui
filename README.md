# golintui

[![License](https://img.shields.io/github/license/nakabonne/golintui)](/LICENSE)

`golintui` is a TUI tool that helps you run various kinds of linters with ease and organize its results, with the power of [golangci-lint](https://github.com/golangci/golangci-lint).

![Screenshot](golintui.gif)

## Cool features

- Simple UI
- Selectable linters on the UI.
- Sorting out the issues for each linter
- Able to open files by specifying the issue line


## Installation

**Binary Releases**

For Mac OS or Linux, you can download a binary release [here](https://github.com/nakabonne/golintui/releases).

**Source**

```bash
go get github.com/nakabonne/golintui
```

## Usage

Requires: [golangci-lint](https://github.com/golangci/golangci-lint) executable.  
  
Be sure to change the CTYPE as shown below if your locale isn't `en_US`. The UI does not display well without it.

```bash
export LC_CTYPE=en_US.UTF-8
```

### Quick Start

```bash
golintui
```

Just press <kbd>r</kbd>, then results from the linters should be shown.

### Keybinds

**Global**

<pre>
  <kbd>r</kbd>: run selected linters against the selected directories
  <kbd>q</kbd>: quit
  <kbd>l</kbd>: next panel
  <kbd>h</kbd>: previous panel
  <kbd>j</kbd>: move down
  <kbd>k</kbd>: move up
</pre>

**Linters Panel**

<pre>
  <kbd>space</kbd>: toggle enabled
</pre>

##### Note that for users who specify `disable-all` in the config file for golangci-lint, it is impossible to disable linters that are enabled in it.

**Source File Panel**

<pre>
  <kbd>space</kbd>: toggle selected
  <kbd>o</kbd>: expand directory
</pre>

**Results Panel**

<pre>
  <kbd>o</kbd>: open a file with the reported line
</pre>

## Settings

### Editor
`golintui` refers to `$EDITOR` by default to open the problematic file. You can change the editor to your taste and habits by setting `$GOLINTUI_OPEN_COMMAND`.  

For instance, for users of VSCode:

```bash
export GOLINTUI_OPEN_COMMAND="code -r"
```

## Editors that can open by specifying a line

- vim(vi)
- emacs
- VSCode

Please let me know how to open a file at a specific line if the editor you're used to is missing.
