# NAME

lf - terminal file manager

# SYNOPSIS

**lf**
[**-command** *command*]
[**-config** *path*]
[**-cpuprofile** *path*]
[**-doc**]
[**-last-dir-path** *path*]
[**-log path**]
[**-memprofile** *path*]
[**-print-last-dir**]
[**-print-selection**]
[**-remote** *command*]
[**-selection-path** *path*]
[**-server**]
[**-single**]
[**-version**]
[**-help**]
[*cd-or-select-path*]

# DESCRIPTION

lf is a terminal file manager.

The source code can be found in the repository at https://github.com/gokcehan/lf

This documentation can either be read from the terminal using `lf -doc` or online at https://github.com/gokcehan/lf/blob/master/doc.md
You can also use the `doc` command (default `<f-1>`) inside lf to view the documentation in a pager.
A man page with the same content is also available in the repository at https://github.com/gokcehan/lf/blob/master/lf.1

You can run `lf -help` to see descriptions of command line options.

# QUICK REFERENCE

The following commands are provided by lf:

	quit                     (default 'q')
	up                       (default 'k' and '<up>')
	half-up                  (default '<c-u>')
	page-up                  (default '<c-b>' and '<pgup>')
	scroll-up                (default '<c-y>')
	down                     (default 'j' and '<down>')
	half-down                (default '<c-d>')
	page-down                (default '<c-f>' and '<pgdn>')
	scroll-down              (default '<c-e>')
	updir                    (default 'h' and '<left>')
	open                     (default 'l' and '<right>')
	jump-next                (default ']')
	jump-prev                (default '[')
	top                      (default 'gg' and '<home>')
	bottom                   (default 'G' and '<end>')
	high                     (default 'H')
	middle                   (default 'M')
	low                      (default 'L')
	toggle
	visual                   (default 'V')
	invert                   (default 'v')
	unselect                 (default 'u')
	glob-select
	glob-unselect
	calcdirsize
	clearmaps
	copy                     (default 'y')
	cut                      (default 'd')
	paste                    (default 'p')
	clear                    (default 'c')
	sync
	draw
	redraw                   (default '<c-l>')
	load
	reload                   (default '<c-r>')
	echo
	echomsg
	echoerr
	cd
	select
	delete         (modal)
	rename         (modal)   (default 'r')
	source
	push
	read           (modal)   (default ':')
	shell          (modal)   (default '$')
	shell-pipe     (modal)   (default '%')
	shell-wait     (modal)   (default '!')
	shell-async    (modal)   (default '&')
	find           (modal)   (default 'f')
	find-back      (modal)   (default 'F')
	find-next                (default ';')
	find-prev                (default ',')
	search         (modal)   (default '/')
	search-back    (modal)   (default '?')
	search-next              (default 'n')
	search-prev              (default 'N')
	filter         (modal)
	setfilter
	mark-save      (modal)   (default 'm')
	mark-load      (modal)   (default "'")
	mark-remove    (modal)   (default '"')
	tag
	tag-toggle               (default 't')
	addcustominfo
	tty-write

The following Visual mode commands are provided by lf:

	visual-accept            (default 'V')
	visual-unselect
	visual-discard           (default '<esc>')
	visual-change            (default 'o')

The following Command-line mode commands are provided by lf:

	cmd-escape               (default '<esc>')
	cmd-complete             (default '<tab>')
	cmd-menu-complete
	cmd-menu-complete-back
	cmd-menu-accept
	cmd-enter                (default '<c-j>' and '<enter>')
	cmd-interrupt            (default '<c-c>')
	cmd-history-next         (default '<c-n>' and '<down>')
	cmd-history-prev         (default '<c-p>' and '<up>')
	cmd-left                 (default '<c-b>' and '<left>')
	cmd-right                (default '<c-f>' and '<right>')
	cmd-home                 (default '<c-a>' and '<home>')
	cmd-end                  (default '<c-e>' and '<end>')
	cmd-delete               (default '<c-d>' and '<delete>')
	cmd-delete-back          (default '<backspace>' and '<backspace2>')
	cmd-delete-home          (default '<c-u>')
	cmd-delete-end           (default '<c-k>')
	cmd-delete-unix-word     (default '<c-w>')
	cmd-yank                 (default '<c-y>')
	cmd-transpose            (default '<c-t>')
	cmd-transpose-word       (default '<a-t>')
	cmd-word                 (default '<a-f>')
	cmd-word-back            (default '<a-b>')
	cmd-delete-word          (default '<a-d>')
	cmd-delete-word-back     (default '<a-backspace>' and '<a-backspace2>')
	cmd-capitalize-word      (default '<a-c>')
	cmd-uppercase-word       (default '<a-u>')
	cmd-lowercase-word       (default '<a-l>')

The following options can be used to customize the behavior of lf:

	anchorfind        bool      (default true)
	autoquit          bool      (default true)
	borderfmt         string    (default "\033[0m")
	cleaner           string    (default '')
	copyfmt           string    (default "\033[7;33m")
	cursoractivefmt   string    (default "\033[7m")
	cursorparentfmt   string    (default "\033[7m")
	cursorpreviewfmt  string    (default "\033[4m")
	cutfmt            string    (default "\033[7;31m")
	dircache          bool      (default true)
	dircounts         bool      (default false)
	dirfirst          bool      (default true)
	dironly           bool      (default false)
	dirpreviews       bool      (default false)
	drawbox           bool      (default false)
	dupfilefmt        string    (default '%f.~%n~')
	errorfmt          string    (default "\033[7;31;47m")
	filesep           string    (default "\n")
	filtermethod      string    (default 'text')
	findlen           int       (default 1)
	hidden            bool      (default false)
	hiddenfiles       []string  (default '.*' for Unix and '' for Windows)
	history           bool      (default true)
	icons             bool      (default false)
	ifs               string    (default '')
	ignorecase        bool      (default true)
	ignoredia         bool      (default true)
	incfilter         bool      (default false)
	incsearch         bool      (default false)
	info              []string  (default '')
	infotimefmtnew    string    (default 'Jan _2 15:04')
	infotimefmtold    string    (default 'Jan _2  2006')
	locale            string    (default '')
	mouse             bool      (default false)
	number            bool      (default false)
	numberfmt         string    (default "\033[33m")
	period            int       (default 0)
	preserve          []string  (default "mode")
	preview           bool      (default true)
	previewer         string    (default '')
	promptfmt         string    (default "\033[32;1m%u@%h\033[0m:\033[34;1m%d\033[0m\033[1m%f\033[0m")
	ratios            []int     (default '1:2:3')
	relativenumber    bool      (default false)
	reverse           bool      (default false)
	roundbox          bool      (default false)
	rulerfmt          string    (default "  %a|  %p|  \033[7;31m %m \033[0m|  \033[7;33m %c \033[0m|  \033[7;35m %s \033[0m|  \033[7;34m %f \033[0m|  %i/%t")
	scrolloff         int       (default 0)
	searchmethod      string    (default 'text')
	selectfmt         string    (default "\033[7;35m")
	selmode           string    (default 'all')
	shell             string    (default 'sh' for Unix and 'cmd' for Windows)
	shellflag         string    (default '-c' for Unix and '/c' for Windows)
	shellopts         []string  (default '')
	showbinds         bool      (default true)
	sixel             bool      (default false)
	smartcase         bool      (default true)
	smartdia          bool      (default false)
	sortby            string    (default 'natural')
	statfmt           string    (default "\033[36m%p\033[0m| %c| %u| %g| %S| %t| -> %l")
	tabstop           int       (default 8)
	tagfmt            string    (default "\033[31m")
	tempmarks         string    (default '')
	timefmt           string    (default 'Mon Jan _2 15:04:05 2006')
	truncatechar      string    (default '~')
	truncatepct       int       (default 100)
	visualfmt         string    (default "\033[7;36m")
	waitmsg           string    (default 'Press any key to continue')
	watch             bool      (default false)
	wrapscan          bool      (default true)
	wrapscroll        bool      (default false)
	user_{option}     string    (default none)

The following environment variables are exported for shell commands:

	f
	fs
	fv
	fx
	id
	PWD
	OLDPWD
	LF_LEVEL
	OPENER
	VISUAL
	EDITOR
	PAGER
	SHELL
	lf
	lf_{option}
	lf_user_{option}
	lf_width
	lf_height
	lf_count
	lf_mode

The following special shell commands are used to customize the behavior of lf when defined:

	open
	paste
	rename
	delete
	pre-cd
	on-cd
	on-load
	on-focus-gained
	on-focus-lost
	on-init
	on-select
	on-redraw
	on-quit

The following commands/keybindings are provided by default:

	Unix
	cmd open &$OPENER "$f"
	map e $$EDITOR "$f"
	map i $$PAGER "$f"
	map w $$SHELL
	cmd doc $$lf -doc | $PAGER
	map <f-1> doc
	cmd maps $lf -remote "query $id maps" | $PAGER
	cmd nmaps $lf -remote "query $id nmaps" | $PAGER
	cmd vmaps $lf -remote "query $id vmaps" | $PAGER
	cmd cmaps $lf -remote "query $id cmaps" | $PAGER
	cmd cmds $lf -remote "query $id cmds" | $PAGER

	Windows
	cmd open &%OPENER% %f%
	map e $%EDITOR% %f%
	map i !%PAGER% %f%
	map w $%SHELL%
	cmd doc !%lf% -doc | %PAGER%
	map <f-1> doc
	cmd maps !%lf% -remote "query %id% maps" | %PAGER%
	cmd nmaps !%lf% -remote "query %id% nmaps" | %PAGER%
	cmd vmaps !%lf% -remote "query %id% vmaps" | %PAGER%
	cmd cmaps !%lf% -remote "query %id% cmaps" | %PAGER%
	cmd cmds !%lf% -remote "query %id% cmds" | %PAGER%

The following additional keybindings are provided by default:

	map zh set hidden!
	map zr set reverse!
	map zn set info
	map zs set info size
	map zt set info time
	map za set info size:time
	map sn :set sortby natural; set info
	map ss :set sortby size; set info size
	map st :set sortby time; set info time
	map sa :set sortby atime; set info atime
	map sb :set sortby btime; set info btime
	map sc :set sortby ctime; set info ctime
	map se :set sortby ext; set info
	map gh cd ~
	nmap <space> :toggle; down

If the `mouse` option is enabled, mouse buttons have the following default effects:

	Left mouse button
	    Click on a file or directory to select it.

	Right mouse button
	    Enter a directory or open a file. Also works on the preview window.

	Scroll wheel
	    Move up or down. If Ctrl is pressed, scroll up or down.

# CONFIGURATION

Configuration files should be located at:

	OS       system-wide               user-specific
	Unix     /etc/lf/lfrc              ~/.config/lf/lfrc
	Windows  C:\ProgramData\lf\lfrc    C:\Users\<user>\AppData\Roaming\lf\lfrc

The colors file should be located at:

	OS       system-wide               user-specific
	Unix     /etc/lf/colors            ~/.config/lf/colors
	Windows  C:\ProgramData\lf\colors  C:\Users\<user>\AppData\Roaming\lf\colors

The icons file should be located at:

	OS       system-wide               user-specific
	Unix     /etc/lf/icons             ~/.config/lf/icons
	Windows  C:\ProgramData\lf\icons   C:\Users\<user>\AppData\Roaming\lf\icons

The selection file should be located at:

	Unix     ~/.local/share/lf/files
	Windows  C:\Users\<user>\AppData\Local\lf\files

The marks file should be located at:

	Unix     ~/.local/share/lf/marks
	Windows  C:\Users\<user>\AppData\Local\lf\marks

The tags file should be located at:

	Unix     ~/.local/share/lf/tags
	Windows  C:\Users\<user>\AppData\Local\lf\tags

The history file should be located at:

	Unix     ~/.local/share/lf/history
	Windows  C:\Users\<user>\AppData\Local\lf\history

You can configure these locations with the following variables given with their order of precedences and their default values:

	Unix
	    $LF_CONFIG_HOME
	    $XDG_CONFIG_HOME
	    ~/.config

	    $LF_DATA_HOME
	    $XDG_DATA_HOME
	    ~/.local/share

	Windows
	    %LF_CONFIG_HOME%
	    %APPDATA%

	    %LF_DATA_HOME%
	    %LOCALAPPDATA%

A sample configuration file can be found at
https://github.com/gokcehan/lf/blob/master/etc/lfrc.example

# COMMANDS

This section shows information about built-in commands.
Modal commands do not take any arguments, but instead change the operation mode to read their input conveniently, and so they are meant to be assigned to keybindings.

## quit (default `q`)

Quit lf and return to the shell.

## up (default `k` and `<up>`), half-up (default `<c-u>`), page-up (default `<c-b>` and `<pgup>`), scroll-up (default `<c-y>`), down (default `j` and `<down>`), half-down (default `<c-d>`), page-down (default `<c-f>` and `<pgdn>`), scroll-down (default `<c-e>`)

Move/scroll the current file selection upwards/downwards by one/half a page/full page.

## updir (default `h` and `<left>`)

Change the current working directory to the parent directory.

## open (default `l` and `<right>`)

If the current file is a directory, then change the current directory to it, otherwise, execute the `open` command.
A default `open` command is provided to call the default system opener asynchronously with the current file as the argument.
A custom `open` command can be defined to override this default.

## jump-next (default `]`), jump-prev (default `[`)

Change the current working directory to the next/previous jumplist item.

## top (default `gg` and `<home>`), bottom (default `G` and `<end>`)

Move the current file selection to the top/bottom of the directory.
A count can be specified to move to a specific line, for example, use `3G` to move to the third line.

## high (default `H`), middle (default `M`), low (default `L`)

Move the current file selection to the high/middle/low of the screen.

## toggle

Toggle the selection of the current file or files given as arguments.

## visual (default `V`)

Switch to Visual mode.
If already in Visual mode, discard the visual selection and stay in Visual mode.

## visual-accept (default `V`)

Add the visual selection to the selection list, quit Visual mode and return to Normal mode.

## visual-unselect

Remove the visual selection from the selection list, quit Visual mode and return to Normal mode.

## visual-discard (default `<esc>`)

Discard the visual selection, quit Visual mode and return to Normal mode.

## visual-change (default `o`)

Go to the other end of the current Visual mode selection.

## invert (default `v`)

Reverse the selection of all files in the current directory (i.e. `toggle` all files).
Selections in other directories are not affected by this command.
You can define a new command to select all files in the directory by combining `invert` with `unselect` (i.e. `cmd select-all :unselect; invert`), though this will also remove selections in other directories.

## unselect (default `u`)

Remove the selection of all files in all directories.

## glob-select, glob-unselect

Select/unselect files that match the given glob.

## calcdirsize

Calculate the total size for each of the selected directories.
Option `info` should include `size` and option `dircounts` should be disabled to show this size.
If the total size of a directory is not calculated, it will be shown as `-`.

## clearmaps

Remove all keybindings associated with the `map`, `nmap` and `vmap` command.
This command can be used in the config file to remove the default keybindings.
For safety purposes, `:` is left mapped to the `read` command, and `cmap` keybindings are retained so that it is still possible to exit `lf` using `:quit`.

## copy (default `y`)

If there are no selections, save the path of the current file to the copy buffer, otherwise, copy the paths of selected files.

## cut (default `d`)

If there are no selections, save the path of the current file to the cut buffer, otherwise, copy the paths of selected files.

## paste (default `p`)

Copy/Move files in the copy/cut buffer to the current working directory.
A custom `paste` command can be defined to override this default.

## clear (default `c`)

Clear file paths in copy/cut buffer.

## sync

Synchronize copied/cut files with the server.
This command is automatically called when required.

## draw

Draw the screen.
This command is automatically called when required.

## redraw (default `<c-l>`)

Synchronize the terminal and redraw the screen.

## load

Load modified files and directories.
This command is automatically called when required.

## reload (default `<c-r>`)

Flush the cache and reload all files and directories.

## echo

Print the given arguments to the message line at the bottom.

## echomsg

Print the given arguments to the message line at the bottom and also to the log file.

## echoerr

Print given arguments to the message line at the bottom as `errorfmt` and also to the log file.

## cd

Change the working directory to the given argument.

## select

Change the current file selection to the given argument.

## delete (modal)

Remove the current file or selected file(s).
A custom `delete` command can be defined to override this default.

## rename (modal) (default `r`)

Rename the current file using the built-in method.
A custom `rename` command can be defined to override this default.

## source

Read the configuration file given in the argument.

## push

Simulate key pushes given in the argument.

## read (modal) (default `:`)

Read a command to evaluate.

## shell (modal) (default `$`)

Read a shell command to execute.

## shell-pipe (modal) (default `%`)

Read a shell command to execute piping its standard I/O to the bottom statline.

## shell-wait (modal) (default `!`)

Read a shell command to execute and wait for a key press in the end.

## shell-async (modal) (default `&`)

Read a shell command to execute asynchronously without standard I/O.

## find (modal) (default `f`), find-back (modal) (default `F`), find-next (default `;`), find-prev (default `,`)

Read key(s) to find the appropriate file name match in the forward/backward direction and jump to the next/previous match.

## search (default `/`), search-back (default `?`), search-next (default `n`), search-prev (default `N`)

Read a pattern to search for a file name match in the forward/backward direction and jump to the next/previous match.

## filter (modal), setfilter

Command `filter` reads a pattern to filter out and only view files matching the pattern.
Command `setfilter` does the same but uses an argument to set the filter immediately.
You can supply an argument to `filter` to use as the starting prompt.

## mark-save (modal) (default `m`)

Save the current directory as a bookmark assigned to the given key.

## mark-load (modal) (default `'`)

Change the current directory to the bookmark assigned to the given key.
A special bookmark `'` holds the previous directory after a `mark-load`, `cd`, or `select` command.

## mark-remove (modal) (default `"`)

Remove a bookmark assigned to the given key.

## tag

Tag a file with `*` or a single-width character given in the argument.
You can define a new tag-clearing command by combining `tag` with `tag-toggle` (i.e. `cmd tag-clear :tag; tag-toggle`).

## tag-toggle (default `t`)

Tag a file with `*` or a single width character given in the argument if the file is untagged, otherwise remove the tag.

## addcustominfo

Update the `custom` info field of the given file with the given string.
The info string may contain ANSI escape codes to further customize its appearance.
If no info is provided, clear the file's info instead.

## tty-write

Write the given string to the tty.
This is useful for sending escape sequences to the terminal to control its behavior (e.g. OSC 0 to set the window title).
Using `tty-write` is preferred over directly writing to `/dev/tty` because the latter is not synchronized and can interfere with drawing the UI.

# COMMAND LINE COMMANDS

The prompt character specifies which of the several Command-line modes you are in.
For example, the `read` command takes you to the `:` mode.

When the cursor is at the first character in `:` mode, pressing one of the keys `!`, `$`, `%`, or `&` takes you to the corresponding mode.
You can go back with `cmd-delete-back` (`<backspace>` by default).

The command line commands should be mostly compatible with readline keybindings.
A character refers to a Unicode code point, a word consists of letters and digits, and a unix word consists of any non-blank characters.

## cmd-escape (default `<esc>`)

Quit Command-line mode and return to Normal mode.

## cmd-complete (default `<tab>`)

Autocomplete the current word.

## cmd-menu-complete, cmd-menu-complete-back

Autocomplete the current word with the menu selection.
You need to assign keys to these commands (e.g. `cmap <tab> cmd-menu-complete; cmap <backtab> cmd-menu-complete-back`).
You can use the assigned keys to display the menu and then cycle through completion options.

## cmd-menu-accept

Accept the currently selected match in menu completion and close the menu.

## cmd-enter (default `<c-j>` and `<enter>`)

Execute the current line.

## cmd-interrupt (default `<c-c>`)

Interrupt the current shell-pipe command and return to the Normal mode.

## cmd-history-next (default `<c-n>` and `<down>`), cmd-history-prev (default `<c-p>` and `<up>`)

Go to the next/previous item in the history.

## cmd-left (default `<c-b>` and `<left>`), cmd-right (default `<c-f>` and `<right>`)

Move the cursor to the left/right.

## cmd-home (default `<c-a>` and `<home>`), cmd-end (default `<c-e>` and `<end>`)

Move the cursor to the beginning/end of the line.

## cmd-delete (default `<c-d>` and `<delete>`)

Delete the next character.

## cmd-delete-back (default `<backspace>` and `<backspace2>`)

Delete the previous character.
When at the beginning of a prompt, returns either to Normal mode or to `:` mode.

## cmd-delete-home (default `<c-u>`), cmd-delete-end (default `<c-k>`)

Delete everything up to the beginning/end of the line.

## cmd-delete-unix-word (default `<c-w>`)

Delete the previous unix word.

## cmd-yank (default `<c-y>`)

Paste the buffer content containing the last deleted item.

## cmd-transpose (default `<c-t>`), cmd-transpose-word (default `<a-t>`)

Transpose the positions of the last two characters/words.

## cmd-word (default `<a-f>`), cmd-word-back (default `<a-b>`)

Move the cursor by one word in the forward/backward direction.

## cmd-delete-word (default `<a-d>`)

Delete the next word in the forward direction.

## cmd-delete-word-back (default `<a-backspace>` and `<a-backspace2>`)

Delete the previous word in the backward direction.

## cmd-capitalize-word (default `<a-c>`), cmd-uppercase-word (default `<a-u>`), cmd-lowercase-word (default `<a-l>`)

Capitalize/uppercase/lowercase the current word and jump to the next word.

# OPTIONS

This section shows information about options to customize the behavior.
Character `:` is used as the separator for list options `[]int` and `[]string`.

## anchorfind (bool) (default true)

When this option is enabled, the find command starts matching patterns from the beginning of file names, otherwise, it can match at an arbitrary position.

## autoquit (bool) (default true)

Automatically quit the server when there are no clients left connected.

## borderfmt (string) (default `\033[0m`)

Format string of the box drawing characters enabled by the `drawbox` option.

## cleaner (string) (default ``) (not called if empty)

Set the path of a cleaner file.
The file should be executable.
This file is called if previewing is enabled, the previewer is set, and the previously selected file has its preview cache disabled.
The following arguments are passed to the file, (1) current file name, (2) width, (3) height, (4) horizontal position, (5) vertical position of preview pane and (6) next file name to be previewed respectively.
Preview cleaning is disabled when the value of this option is left empty.

## copyfmt (string) (default `\033[7;33m`)

Format string of the indicator for files to be copied.

## cursoractivefmt (string) (default `\033[7m`), cursorparentfmt string (default `\033[7m`), cursorpreviewfmt string (default `\033[4m`)

Format strings for highlighting the cursor.
`cursoractivefmt` applies in the current directory pane,
`cursorparentfmt` applies in panes that show parents of the current directory,
and `cursorpreviewfmt` applies in panes that preview directories.

The default is to make the active cursor and the parent directory cursor inverted. The preview cursor is underlined.

Some other possibilities to consider for the preview or parent cursors: an empty string for no cursor, `\033[7;2m` for dimmed inverted text (visibility varies by terminal), `\033[7;90m` for inverted text with grey (aka "brightblack") background.

If the format string contains the characters `%s`, it is interpreted as a format string for `fmt.Sprintf`. Such a string should end with the terminal reset sequence.
For example, `\033[4m%s\033[0m` has the same effect as `\033[4m`.

## cutfmt (string) (default `\033[7;31m`)

Format string of the indicator for files to be cut.

## dircache (bool) (default true)

Cache directory contents.

## dircounts (bool) (default false)

When this option is enabled, directory sizes show the number of items inside instead of the total size of the directory, which needs to be calculated for each directory using `calcdirsize`.
This information needs to be calculated by reading the directory and counting the items inside.
Therefore, this option is disabled by default for performance reasons.
This option only has an effect when `info` has a `size` field and the pane is wide enough to show the information.
999 items are counted per directory at most, and bigger directories are shown as `999+`.

## dirfirst (bool) (default true)

Show directories first above regular files.
With `dircounts` enabled, sorting by `size` always separates directories and files, regardless of `dirfirst`.

## dironly (bool) (default false)

Show only directories.

## dirpreviews (bool) (default false)

If enabled, directories will also be passed to the previewer script. This allows custom previews for directories.

## drawbox (bool) (default false)

Draw boxes around panes with box drawing characters.

## dupfilefmt (string) (default `%f.~%n~`)

Format string of file name when creating duplicate files. With the default format, copying a file `abc.txt` to the same directory will result in a duplicate file called `abc.txt.~1~`.
Special expansions are provided, `%f` as the file name, `%b` for the basename (file name without extension), `%e` as the extension (including the dot) and `%n` as the number of duplicates.

## errorfmt (string) (default `\033[7;31;47m`)

Format string of error messages shown in the bottom message line.

If the format string contains the characters `%s`, it is interpreted as a format string for `fmt.Sprintf`. Such a string should end with the terminal reset sequence.
For example, `\033[4m%s\033[0m` has the same effect as `\033[4m`.

## filesep (string) (default `\n`)

File separator used in environment variables `fs`, `fv` and `fx`.

## filtermethod (string) (default `text`)

How filter command patterns are treated.
Currently supported methods are `text` (i.e. string literals), `glob` (i.e shell globs) and `regex` (i.e. regular expressions).
See `SEARCHING FILES` for more details.

## findlen (int) (default 1)

Number of characters prompted for the find command.
When this value is set to 0, find command prompts until there is only a single match left.

## hidden (bool) (default false)

Show hidden files.
On Unix systems, hidden files are determined by the value of `hiddenfiles`.
On Windows, files with hidden attributes are also considered hidden files.

## hiddenfiles ([]string) (default `.*` for Unix and `` for Windows)

List of hidden file glob patterns.
Patterns can be given as relative or absolute paths.
Globbing supports the usual special characters, `*` to match any sequence, `?` to match any character, and `[...]` or `[^...]` to match character sets or ranges.
In addition, if a pattern starts with `!`, then its matches are excluded from hidden files. To add multiple patterns, use `:` as a separator. Example: `.*:lost+found:*.bak`

## history (bool) (default true)

Save command history.

## icons (bool) (default false)

Show icons before each item in the list.

## ifs (string) (default ``)

Sets `IFS` variable in shell commands.
It works by adding the assignment to the beginning of the command string as `IFS=...; ...`.
The reason is that `IFS` variable is not inherited by the shell for security reasons.
This method assumes a POSIX shell syntax so it can fail for non-POSIX shells.
This option has no effect when the value is left empty.
This option does not have any effect on Windows.

## ignorecase (bool) (default true)

Ignore case in sorting and search patterns.

## ignoredia (bool) (default true)

Ignore diacritics in sorting and search patterns.

## incsearch (bool) (default false)

Jump to the first match after each keystroke during searching.

## incfilter (bool) (default false)

Apply filter pattern after each keystroke during filtering.

## info ([]string)  (default ``)

A list of information that is shown for directory items at the right side of the pane.
Currently supported information types are `size`, `time`, `atime`, `btime`, `ctime`, `perm`, `user`, `group` and `custom`.
The `custom` type is empty by default and can be updated using the `addcustominfo` command.
Information is only shown when the pane width is more than twice the width of information.

## infotimefmtnew (string) (default `Jan _2 15:04`)

Format string of the file time shown in the info column when it matches this year.

## infotimefmtold (string) (default `Jan _2  2006`)

Format string of the file time shown in the info column when it doesn't match this year.

## locale (string) (default ``)

An IETF BCP 47 language tag (e.g. `zh-CN`) for specifying the locale used when using sort type `natural` and `name`.
An empty string means disable locale ordering, and the special value `*` is used to indicate reading the locale setting from the system environment.
This feature is currently experimental.

## mouse (bool) (default false)

Send mouse events as input.

## number (bool) (default false)

Show the position number for directory items on the left side of the pane.
When the `relativenumber` option is enabled, only the current line shows the absolute position and relative positions are shown for the rest.

## numberfmt (string) (default `\033[33m`)

Format string of the position number for each line.

## period (int) (default 0)

Set the interval in seconds for periodic checks of directory updates.
This works by periodically calling the `load` command.
Note that directories are already updated automatically in many cases.
This option can be useful when there is an external process changing the displayed directory and you are not doing anything in lf.
Periodic checks are disabled when the value of this option is set to zero.

## preserve ([]string) (default `mode`)

List of attributes that are preserved when copying files.
Currently supported attributes are `mode` (i.e. access mode) and `timestamps` (i.e. modification time and access time).
Note that preserving other attributes like ownership of change/birth timestamp is desirable, but not portably supported in Go.

## preview (bool) (default true)

Show previews of files and directories at the rightmost pane.
If the file has more lines than the preview pane, the rest of the lines are not read.
Files containing the null character (U+0000) in the read portion are considered binary files and displayed as `binary`.

## previewer (string) (default ``) (not filtered if empty)

Set the path of a previewer file to filter the content of regular files for previewing.
The file should be executable.
The following arguments are passed to the file, (1) current file name, (2) width, (3) height, (4) horizontal position, and (5) vertical position of preview pane respectively.
SIGPIPE signal is sent when enough lines are read.
If the previewer returns a non-zero exit code, then the preview cache for the given file is disabled.
This means that if the file is selected in the future, the previewer is called once again.
Preview filtering is disabled and files are displayed as they are when the value of this option is left empty.

## promptfmt (string) (default `\033[32;1m%u@%h\033[0m:\033[34;1m%d\033[0m\033[1m%f\033[0m`)

Format string of the prompt shown in the top line.
Special expansions are provided, `%u` as the user name, `%h` as the hostname, `%w` as the working directory, `%d` as the working directory with a trailing path separator, `%f` as the file name, and `%F` as the current filter. `%S` may be used once and will provide a spacer so that the following parts are right aligned on the screen.
The home folder is shown as `~` in the working directory expansion.
Directory names are automatically shortened to a single character starting from the leftmost parent when the prompt does not fit the screen.

## ratios ([]int) (default `1:2:3`)

List of ratios of pane widths.
Number of items in the list determines the number of panes in the UI.
When the `preview` option is enabled, the rightmost number is used for the width of the preview pane.

## relativenumber (bool) (default false)

Show the position number relative to the current line.
When `number` is enabled, the current line shows the absolute position, otherwise nothing is shown.

## reverse (bool) (default false)

Reverse the direction of sort.

## roundbox (bool) (default false)

Draw rounded outer corners when the `drawbox` option is enabled.

## rulerfmt (string) (default `  %a|  %p|  \033[7;31m %m \033[0m|  \033[7;33m %c \033[0m|  \033[7;35m %s \033[0m|  \033[7;36m %v \033[0m|  \033[7;34m %f \033[0m|  %i/%t`)

Format string of the ruler shown in the bottom right corner.
Special expansions are provided, `%a` as the pressed keys, `%p` as the progress of file operations, `%m` as the number of files to be cut (moved), `%c` as the number of files to be copied, `%s` as the number of selected files, `%v` as the number of visually selected files, `%f` as the filter, `%i` as the position of the cursor, `%t` as the number of files shown in the current directory, `%h` as the number of files hidden in the current directory, `%P` as the scroll percentage, and `%d` as the amount of free disk space remaining.
Additional expansions are provided for environment variables exported by lf, in the form `%{lf_<name>}` (e.g. `%{lf_selmode}`). This is useful for displaying the current settings.
Expansions are also provided for user-defined options, in the form `%{lf_user_<name>}` (e.g. `%{lf_user_foo}`).
The `|` character splits the format string into sections. Any section containing a failed expansion (result is a blank string) is discarded and not shown.

## selectfmt (string) (default `\033[7;35m`)

Format string of the indicator for files that are selected.

## selmode (string) (default `all`)

Selection mode for commands.
When set to `all` it will use the selected files from all directories.
When set to `dir` it will only use the selected files in the current directory.

## scrolloff (int) (default 0)

Minimum number of offset lines shown at all times at the top and bottom of the screen when scrolling.
The current line is kept in the middle when this option is set to a large value that is bigger than the half of number of lines.
A smaller offset can be used when the current file is close to the beginning or end of the list to show the maximum number of items.

## searchmethod (string) default `text`)

How search command patterns are treated.
Currently supported methods are `text` (i.e. string literals), `glob` (i.e shell globs) and `regex` (i.e. regular expressions).
See `SEARCHING FILES` for more details.

## shell (string) (default `sh` for Unix and `cmd` for Windows)

Shell executable to use for shell commands.
Shell commands are executed as `shell shellopts shellflag command -- arguments`.

## shellflag (string) (default `-c` for Unix and `/c` for Windows)

Command line flag used to pass shell commands.

## shellopts ([]string)  (default ``)

List of shell options to pass to the shell executable.

## showbinds (bool) (default true)

Show bindings associated with pressed keys.

## sixel (bool) (default false)

Render sixel images in preview.

## smartcase (bool) (default true)

Override `ignorecase` option when the pattern contains an uppercase character.
This option has no effect when `ignorecase` is disabled.

## smartdia (bool) (default false)

Override `ignoredia` option when the pattern contains a character with diacritic.
This option has no effect when `ignoredia` is disabled.

## sortby (string) (default `natural`)

Sort type for directories.
Currently supported sort types are `natural`, `name`, `size`, `time`, `atime`, `btime`, `ctime`, `ext` and `custom`.

Meaning of each sort type:

	natural   file name (track_2.flac comes before track_10.flac)
	name      file name (track_10.flac comes before track_2.flac)
	size      file size
	time      time of last data modification
	atime     time of last access
	btime     time of file birth
	ctime     time of last status (inode) change
	ext       file extension
	custom    property defined via `addcustominfo`

## statfmt (string) (default `\033[36m%p\033[0m| %c| %u| %g| %S| %t| -> %l`)

Format string of the file info shown in the bottom left corner.
Special expansions are provided, `%p` as the file permissions, `%c` as the link count, `%u` as the user, `%g` as the group, `%s` as the file size, `%S` as the file size but with a fixed width of five characters (left-padded with spaces), `%t` as the last modified time, `%l` as the link target, `%m` as the current mode and `%M` as the current mode but also shown in Normal mode (displaying `NORMAL` instead of a blank string).

The `|` character splits the format string into sections. Any section containing a failed expansion (result is a blank string) is discarded and not shown.

File size is formatted using first letter of IEC 80000-13:2025 prefixes for binary multiples (i.e. 1024 bytes is `1.0K`).

## tabstop (int) (default 8)

Number of space characters to show for horizontal tabulation (U+0009) character.

## tagfmt (string) (default `\033[31m`)

Format string of the tags.

If the format string contains the characters `%s`, it is interpreted as a format string for `fmt.Sprintf`. Such a string should end with the terminal reset sequence.
For example, `\033[4m%s\033[0m` has the same effect as `\033[4m`.

## tempmarks (string) (default ``)

Marks to be considered temporary (e.g. `abc` refers to marks `a`, `b`, and `c`).
These marks are not synced to other clients and they are not saved in the bookmarks file.
Note that the special bookmark `` ` `` is always treated as temporary and it does not need to be specified.

## timefmt (string) (default `Mon Jan _2 15:04:05 2006`)

Format string of the file modification time shown in the bottom line.

## truncatechar (string) (default `~`)

The truncate character that is shown at the end when the file name does not fit into the pane.

## truncatepct (int) (default 100)

When a filename is too long to be shown completely, the available space is
partitioned into two pieces. truncatepct defines a fraction (in percent
between 0 and 100) for the size of the first piece, which will show the
beginning of the filename. The second piece will show the end of the filename
and will use the rest of the available space. Both pieces are separated by the
truncation character (truncatechar).
A value of 100 will only show the beginning of the filename,
while a value of 0 will only show the end of the filename, e.g.:

- `set truncatepct 100` -> `very-long-filename-tr~` (default)

- `set truncatepct 50`  -> `very-long-f~-truncated`

- `set truncatepct 0`   -> `~ng-filename-truncated`

## visualfmt (string) (default `\033[7;36m`)

Format string of the indicator for files that are visually selected.

## waitmsg (string) (default `Press any key to continue`)

String shown after commands of shell-wait type.

## watch (bool) (default false)

Watch the filesystem for changes using `fsnotify` to automatically refresh file information.
FUSE is currently not supported due to limitations in `fsnotify`.

## wrapscan (bool) (default true)

Searching can wrap around the file list.

## wrapscroll (bool) (default false)

Scrolling can wrap around the file list.

## user_{option} (string) (default none)

Any option that is prefixed with `user_` is a user-defined option and can be set to any string.
Inside a user-defined command, the value will be provided in the `lf_user_{option}` environment variable.
These options are not used by lf and are not persisted.

# ENVIRONMENT VARIABLES

The following variables are exported for shell commands:
These are referred to with a `$` prefix on POSIX shells (e.g. `$f`), between `%` characters on Windows cmd (e.g. `%f%`), and with a `$env:` prefix on Windows PowerShell (e.g. `$env:f`).

## f

Current file selection as a full path.

## fs

Selected file(s) separated with the value of `filesep` option as full path(s).

## fv

Visually selected file(s) separated with the value of `filesep` option as full path(s).

## fx

Selected file(s) (i.e. `fs`, never `fv`) if there are any selected files, otherwise current file selection (i.e. `f`).

## id

Id of the running client.

## PWD

Present working directory.

## OLDPWD

Initial working directory.

## LF_LEVEL

The value of this variable is set to the current nesting level when you run lf from a shell spawned inside lf.
You can add the value of this variable to your shell prompt to make it clear that your shell runs inside lf.
For example, with POSIX shells, you can use `[ -n "$LF_LEVEL" ] && PS1="$PS1""(lf level: $LF_LEVEL) "` in your shell configuration file (e.g. `~/.bashrc`).

## OPENER

If this variable is set in the environment, use the same value. Otherwise, this is set to `start` in Windows, `open` in MacOS, `xdg-open` in others.

## EDITOR

If VISUAL is set in the environment, use its value. Otherwise, use the value of the environment variable EDITOR. If neither variable is set, this is set to `vi` on Unix, `notepad` in Windows.

## PAGER

If this variable is set in the environment, use the same value. Otherwise, this is set to `less` on Unix, `more` in Windows.

## SHELL

If this variable is set in the environment, use the same value. Otherwise, this is set to `sh` on Unix, `cmd` in Windows.

## lf

Absolute path to the currently running lf binary, if it can be found. Otherwise, this is set to the string `lf`.

## lf_{option}

Value of the {option}.

## lf_user_{option}

Value of the user_{option}.

## lf_width, lf_height

Width/Height of the terminal.

## lf_count

Value of the count associated with the current command.

## lf_mode

Current mode that `lf` is operating in.
This is useful for customizing keybindings depending on what the current mode is.
Possible values are `delete`, `rename`, `filter`, `find`, `mark`, `search`, `command`, `shell`, `pipe` (when running a shell-pipe command), `normal`, `visual` and `unknown`.

# SPECIAL COMMANDS

This section shows information about special shell commands.

## open

This shell command can be defined to override the default `open` command when the current file is not a directory.

## paste

This shell command can be defined to override the default `paste` command.

## rename

This shell command can be defined to override the default `rename` command.

## delete

This shell command can be defined to override the default `delete` command.

## pre-cd

This shell command can be defined to be executed before changing a directory.

## on-cd

This shell command can be defined to be executed after changing a directory.

## on-load

This shell command can be defined to be executed after loading a directory.
It provides the files inside the directory as arguments.

## on-focus-gained

This shell command can be defined to be executed when the terminal gains focus.

## on-focus-lost

This shell command can be defined to be executed when the terminal loses focus.

## on-init

This shell command can be defined to be executed after initializing and connecting to the server.

## on-select

This shell command can be defined to be executed after the selection changes.

## on-redraw

This shell command can be defined to be executed after the screen is redrawn or if the terminal is resized.

## on-quit

This shell command can be defined to be executed before quitting.

# PREFIXES

The following command prefixes are used by lf:

	:  read (default)  builtin/custom command
	$  shell           shell command
	%  shell-pipe      shell command running with the UI
	!  shell-wait      shell command waiting for a key press
	&  shell-async     shell command running asynchronously

The same evaluator is used for the command line and the configuration file for reading shell commands.
The difference is that prefixes are not necessary in the command line.
Instead, different modes are provided to read corresponding commands.
These modes are mapped to the prefix keys above by default.
Visual mode mappings are defined the same way Normal mode mappings are defined.

# SYNTAX

Characters from `#` to newline are comments and ignored:

	# comments start with `#`

The following commands (`set`, `setlocal`, `map`, `nmap`, `vmap`, `cmap`, and `cmd`) are used for configuration.

Command `set` is used to set an option which can be a boolean, integer, or string:

	set hidden         # boolean enable
	set hidden true    # boolean enable
	set nohidden       # boolean disable
	set hidden false   # boolean disable
	set hidden!        # boolean toggle
	set scrolloff 10   # integer value
	set sortby time    # string value w/o quotes
	set sortby 'time'  # string value with single quotes (whitespaces)
	set sortby "time"  # string value with double quotes (backslash escapes)

Command `setlocal` is used to set a local option for a directory which can be a boolean or string.
Currently supported local options are `dircounts`, `dirfirst`, `dironly`, `hidden`, `info`, `reverse`, `sortby` and `locale`.
Adding a trailing path separator (i.e. `/` for Unix and `\` for Windows) sets the option for the given directory along with its subdirectories:

	setlocal /foo/bar hidden         # boolean enable
	setlocal /foo/bar hidden true    # boolean enable
	setlocal /foo/bar nohidden       # boolean disable
	setlocal /foo/bar hidden false   # boolean disable
	setlocal /foo/bar hidden!        # boolean toggle
	setlocal /foo/bar sortby time    # string value w/o quotes
	setlocal /foo/bar sortby 'time'  # string value with single quotes (whitespaces)
	setlocal /foo/bar sortby "time"  # string value with double quotes (backslash escapes)
	setlocal /foo/bar  hidden        # for only '/foo/bar' directory
	setlocal /foo/bar/ hidden        # for '/foo/bar' and its subdirectories (e.g. '/foo/bar/baz')

Command `map` is used to bind a key in Normal and Visual mode to a command which can be a builtin command, custom command, or shell command:

	map gh cd ~        # builtin command
	map D trash        # custom command
	map i $less $f     # shell command
	map U !du -csh *   # waiting shell command

Command `nmap` does the same but for Normal mode only.

Command `vmap` does the same but for Visual mode only.

Overview of which map command works in which mode:

	map                Normal, Visual
	nmap               Normal
	vmap               Visual
	cmap               Command-line

Command `cmap` is used to bind a key on the command line to a command line command or any other command:

	cmap <c-g> cmd-escape
	cmap <a-i> set incsearch!

You can delete an existing binding by leaving the expression empty:

	map gh             # deletes 'gh' mapping in Normal and Visual mode
	nmap v             # deletes 'v' mapping in Normal mode
	vmap o             # deletes 'o' mapping in Visual mode
	cmap <c-g>         # deletes '<c-g>' mapping

Command `cmd` is used to define a custom command:

	cmd usage $du -h -d1 | less

You can delete an existing command by leaving the expression empty:

	cmd trash          # deletes 'trash' command

If there is no prefix then `:` is assumed:

	map zt set info time

An explicit `:` can be provided to group statements until a newline which is especially useful for `map` and `cmd` commands:

	map st :set sortby time; set info time

If you need multiline you can wrap statements in `{{` and `}}` after the proper prefix.

	map st :{{
	    set sortby time
	    set info time
	}}

# KEY MAPPINGS

Regular keys are assigned to a command with the usual syntax:

	map a down

Keys combined with the shift key simply use the uppercase letter:

	map A down

Special keys are written in between `<` and `>` characters and always use lowercase letters:

	map <enter> down

Angle brackets can be assigned with their special names:

	map <lt> down
	map <gt> down

Function keys are prefixed with `f` character:

	map <f-1> down

Keys combined with the control key are prefixed with a `c` character:

	map <c-a> down

Keys combined with the alt key are assigned in two different ways depending on the behavior of your terminal.
Older terminals (e.g. xterm) may set the 8th bit of a character when the alt key is pressed.
On these terminals, you can use the corresponding byte for the mapping:

	map á down

Newer terminals (e.g. gnome-terminal) may prefix the key with an escape key when the alt key is pressed.
lf uses the escape delaying mechanism to recognize alt keys in these terminals (delay is 100ms).
On these terminals, keys combined with the alt key are prefixed with an `a` character:

	map <a-a> down

It is possible to combine special keys with modifiers:

	map <a-enter> down

WARNING: Some key combinations will likely be intercepted by your OS, window manager, or terminal.
Other key combinations cannot be recognized by lf due to the way terminals work (e.g. `Ctrl+h` combination sends a backspace key instead).
The easiest way to find out the name of a key combination and whether it will work on your system is to press the key while lf is running and read the name from the `unknown mapping` error.

Mouse buttons are prefixed with an `m` character:

	map <m-1> down  # primary
	map <m-2> down  # secondary
	map <m-3> down  # middle
	map <m-4> down
	map <m-5> down
	map <m-6> down
	map <m-7> down
	map <m-8> down

Mouse wheel events are also prefixed with an `m` character:

	map <m-up>    down
	map <m-down>  down
	map <m-left>  down
	map <m-right> down

# PUSH MAPPINGS

The usual way to map a key sequence is to assign it to a named or unnamed command.
While this provides a clean way to remap built-in keys as well as other commands, it can be limiting at times.
For this reason, the `push` command is provided by lf.
This command is used to simulate key pushes given as its arguments.
You can `map` a key to a `push` command with an argument to create various keybindings.

This is mainly useful for two purposes.
First, it can be used to map a command with a command count:

	map <c-j> push 10j

Second, it can be used to avoid typing the name when a command takes arguments:

	map r push :rename<space>

One thing to be careful of is that since the `push` command works with keys instead of commands it is possible to accidentally create recursive bindings:

	map j push 2j

These types of bindings create a deadlock when executed.

# SHELL COMMANDS

Regular shell commands are the most basic command type that is useful for many purposes.
For example, we can write a shell command to move the selected file(s) to trash.
A first attempt to write such a command may look like this:

	cmd trash ${{
	    mkdir -p ~/.trash
	    if [ -z "$fs" ]; then
	        mv "$f" ~/.trash
	    else
	        IFS="$(printf '\n\t')"; mv $fs ~/.trash
	    fi
	}}

We check `$fs` to see if there are any selected files.
Otherwise, we just delete the current file.
Since this is such a common pattern, a separate `$fx` variable is provided.
We can use this variable to get rid of the conditional:

	cmd trash ${{
	    mkdir -p ~/.trash
	    IFS="$(printf '\n\t')"; mv $fx ~/.trash
	}}

The trash directory is checked each time the command is executed.
We can move it outside of the command so it would only run once at startup:

	${{ mkdir -p ~/.trash }}

	cmd trash ${{ IFS="$(printf '\n\t')"; mv $fx ~/.trash }}

Since these are one-liners, we can drop `{{` and `}}`:

	$mkdir -p ~/.trash

	cmd trash $IFS="$(printf '\n\t')"; mv $fx ~/.trash

Finally, note that we set the `IFS` variable manually in these commands.
Instead, we could use the `ifs` option to set it for all shell commands (i.e. `set ifs "\n"`).
This can be especially useful for interactive use (e.g. `$rm $f` or `$rm $fs` would simply work).
This option is not set by default as it can behave unexpectedly for new users.
However, use of this option is highly recommended and it is assumed in the rest of the documentation.

# PIPING SHELL COMMANDS

Regular shell commands have some limitations in some cases.
When an output or error message is given and the command exits afterwards, the ui is immediately resumed and there is no way to see the message without dropping to shell again.
Also, even when there is no output or error, the UI still needs to be paused while the command is running.
This can cause flickering on the screen for short commands and similar distractions for longer commands.

Instead of pausing the UI, piping shell commands connect stdin, stdout, and stderr of the command to the statline at the bottom of the UI.
This can be useful for programs following the Unix philosophy to give no output in the success case, and brief error messages or prompts in other cases.

For example, following rename command prompts for overwrite in the statline if there is an existing file with the given name:

	cmd rename %mv -i $f $1

You can also output error messages in the command and it will show up in the statline.
For example, an alternative rename command may look like this:

	cmd rename %[ -e $1 ] && printf "file exists" || mv $f $1

Note that input is line buffered and output and error are byte buffered.

# WAITING SHELL COMMANDS

Waiting shell commands are similar to regular shell commands except that they wait for a key press when the command is finished.
These can be useful to see the output of a program before the UI is resumed.
Waiting shell commands are more appropriate than piping shell commands when the command is verbose and the output is best displayed as multiline.

# ASYNCHRONOUS SHELL COMMANDS

Asynchronous shell commands are used to start a command in the background and then resume operation without waiting for the command to finish.
Stdin, stdout, and stderr of the command are neither connected to the terminal nor the UI.

# REMOTE COMMANDS

One of the more advanced features in lf is remote commands.
All clients connect to a server on startup.
It is possible to send commands to all or any of the connected clients over the common server.
This is used internally to notify file selection changes to other clients.

To use this feature, you need to use a client which supports communicating with a Unix domain socket.
OpenBSD implementation of netcat (nc) is one such example.
You can use it to send a command to the socket file:

	echo 'send echo hello world' | nc -U ${XDG_RUNTIME_DIR:-/tmp}/lf.${USER}.sock

Since such a client may not be available everywhere, lf comes bundled with a command line flag to be used as such.
When using lf, you do not need to specify the address of the socket file.
This is the recommended way of using remote commands since it is shorter and immune to socket file address changes:

	lf -remote 'send echo hello world'

In this command `send` is used to send the rest of the string as a command to all connected clients.
You can optionally give it an ID number to send a command to a single client:

	lf -remote 'send 1234 echo hello world'

All clients have a unique id number but you may not be aware of the id number when you are writing a command.
For this purpose, an `$id` variable is exported to the environment for shell commands.
The value of this variable is set to the process ID of the client.
You can use it to send a remote command from a client to the server which in return sends a command back to itself.
So now you can display a message in the current client by calling the following in a shell command:

	lf -remote "send $id echo hello world"

Since lf does not have control flow syntax, remote commands are used for such needs.
For example, you can configure the number of columns in the UI with respect to the terminal width as follows:

	cmd recol %{{
	    if [ $lf_width -le 80 ]; then
	        lf -remote "send $id set ratios 1:2"
	    elif [ $lf_width -le 160 ]; then
	        lf -remote "send $id set ratios 1:2:3"
	    else
	        lf -remote "send $id set ratios 1:2:3:5"
	    fi
	}}

In addition, the `query` command can be used to obtain information about a specific lf instance by providing its id:

	lf -remote "query $id maps"

The following types of information are supported:

	maps     list of mappings created by the 'map', 'nmap' and 'vmap' command
	nmaps    list of mappings created by the 'nmap' and 'map' command
	vmaps    list of mappings created by the 'vmap' and 'map' command
	cmaps    list of mappings created by the 'cmap' command
	cmds     list of commands created by the 'cmd' command
	jumps    contents of the jump list, showing previously visited locations
	history  list of previously executed commands on the command line
	files    list of files in the currently open directory as displayed by lf, empty if dir is still loading

When listing mappings the characters in the first column are:

	n  Normal
	v  Visual
	c  Command-line

This is useful for scripting actions based on the internal state of lf.
For example, to select a previous command using fzf and execute it:

	map <a-h> ${{
	    clear
	    cmd=$(
	        lf -remote "query $id history" |
	        awk -F'\t' 'NR > 1 { print $NF}' |
	        sort -u |
	        fzf --reverse --prompt='Execute command: '
	    )
	    lf -remote "send $id $cmd"
	}}

There is also a `quit` command to quit the server when there are no connected clients left, and a `quit!` command to force quit the server by closing client connections first:

	lf -remote 'quit'
	lf -remote 'quit!'

Lastly, there is a `conn` command to connect the server to a client.
This should not be needed for users.

# FILE OPERATIONS

lf uses its own built-in copy and move operations by default.
These are implemented as asynchronous operations and progress is shown in the bottom ruler.
These commands do not overwrite existing files or directories with the same name.
Instead, a suffix that is compatible with the `--backup=numbered` option in GNU cp is added to the new files or directories.
Only file modes and (some) timestamps can be preserved (see `preserve` option), all other attributes are ignored including ownership, context, and xattr.
Special files such as character and block devices, named pipes, and sockets are skipped and links are not followed.
Moving is performed using the rename operation of the underlying OS.
For cross-device moving, lf falls back to copying and then deletes the original files if there are no errors.
Operation errors are shown in the message line as well as the log file and they do not preemptively finish the corresponding file operation.

File operations can be performed on the currently selected file or on multiple files by selecting them first.
When you `copy` a file, lf doesn't actually copy the file on the disk, but only records its name to a file.
The actual file copying takes place when you `paste`.
Similarly `paste` after a `cut` operation moves the file.

You can customize copy and move operations by defining a `paste` command.
This is a special command that is called when it is defined instead of the built-in implementation.
You can use the following example as a starting point:

	cmd paste %{{
	    load=$(cat ~/.local/share/lf/files)
	    mode=$(echo "$load" | sed -n '1p')
	    list=$(echo "$load" | sed '1d')
	    if [ $mode = 'copy' ]; then
	        cp -R $list .
	    elif [ $mode = 'move' ]; then
	        mv $list .
	        rm ~/.local/share/lf/files
	        lf -remote 'send clear'
	    fi
	}}

Some useful things to be considered are to use the backup (`--backup`) and/or preserve attributes (`-a`) options with `cp` and `mv` commands if they support it (i.e. GNU implementation), change the command type to asynchronous, or use `rsync` command with progress bar option for copying and feed the progress to the client periodically with remote `echo` calls.

By default, lf does not assign `delete` command to a key to protect new users.
You can customize file deletion by defining a `delete` command.
You can also assign a key to this command if you like.
An example command to move selected files to a trash folder and remove files completely after a prompt is provided in the example configuration file.

# SEARCHING FILES

There are two mechanisms implemented in lf to search a file in the current directory.
Searching is the traditional method to move the selection to a file matching a given pattern.
Finding is an alternative way to search for a pattern possibly using fewer keystrokes.

The searching mechanism is implemented with commands `search` (default `/`), `search-back` (default `?`), `search-next` (default `n`), and `search-prev` (default `N`).
You can set `searchmethod` to `glob` to match using a glob pattern.
Globbing supports `*` to match any sequence, `?` to match any character, and `[...]` or `[^...]` to match character sets or ranges.
You can set `searchmethod` to `regex` to match using a regex pattern.
For a full overview of Go's RE2 syntax, see https://pkg.go.dev/regexp/syntax.
You can enable `incsearch` option to jump to the current match at each keystroke while typing.
In this mode, you can either use `cmd-enter` to accept the search or use `cmd-escape` to cancel the search.
You can also map some other commands with `cmap` to accept the search and execute the command immediately afterwards.
For example, you can use the right arrow key to finish the search and open the selected file with the following mapping:

	cmap <right> :cmd-enter; open

The finding mechanism is implemented with commands `find` (default `f`), `find-back` (default `F`), `find-next` (default `;`), `find-prev` (default `,`).
You can disable `anchorfind` option to match a pattern at an arbitrary position in the filename instead of the beginning.
You can set the number of keys to match using `findlen` option.
If you set this value to zero, then the keys are read until there is only a single match.
The default values of these two options are set to jump to the first file with the given initial.

Some options effect both searching and finding.
You can disable `wrapscan` option to prevent searches from being wrapped around at the end of the file list.
You can disable `ignorecase` option to match cases in the pattern and the filename.
This option is already automatically overridden if the pattern contains upper-case characters.
You can disable `smartcase` option to disable this behavior.
Two similar options `ignoredia` and `smartdia` are provided to control matching diacritics in Latin letters.

# OPENING FILES

You can define an `open` command (default `l` and `<right>`) to configure file opening.
This command is only called when the current file is not a directory, otherwise, the directory is entered instead.
You can define it just as you would define any other command:

	cmd open $vi $fx

It is possible to use different command types:

	cmd open &xdg-open $f

You may want to use either file extensions or mime types from `file` command:

	cmd open ${{
	    case $(file --mime-type -Lb $f) in
	        text/*) vi $fx;;
	        *) for f in $fx; do xdg-open $f > /dev/null 2> /dev/null & done;;
	    esac
	}}

You may want to use `setsid` before your opener command to have persistent processes that continue to run after lf quits.

Regular shell commands (i.e. `$`) drop to the terminal which results in a flicker for commands that finish immediately (e.g. `xdg-open` in the above example).
If you want to use asynchronous shell commands (i.e. `&`) but also want to use the terminal when necessary (e.g. `vi` in the above example), you can use a remote command:

	cmd open &{{
	    case $(file --mime-type -Lb $f) in
	        text/*) lf -remote "send $id \$vi \$fx";;
	        *) for f in $fx; do xdg-open $f > /dev/null 2> /dev/null & done;;
	    esac
	}}

Note that asynchronous shell commands run in their own process group by default so they do not require the manual use of `setsid`.

The following command is provided by default:

	cmd open &$OPENER $f

You may also use any other existing file openers as you like.
Possible options are `libfile-mimeinfo-perl` (executable name is `mimeopen`), `rifle` (ranger's default file opener), or `mimeo` to name a few.

# PREVIEWING FILES

lf previews files on the preview pane by printing the file until the end or until the preview pane is filled.
This output can be enhanced by providing a custom preview script for filtering.
This can be used to highlight source codes, list contents of archive files or view PDF or image files to name a few.
For coloring lf recognizes ANSI escape codes.

To use this feature, you need to set the value of `previewer` option to the path of an executable file.
Five arguments are passed to the file, (1) current file name, (2) width, (3) height, (4) horizontal position, and (5) vertical position of preview pane respectively.
The output of the execution is printed in the preview pane.
You may also want to use the same script in your pager mapping as well:

	set previewer ~/.config/lf/pv.sh
	map i $~/.config/lf/pv.sh $f | less -R

For `less` pager, you may instead utilize `LESSOPEN` mechanism so that useful information about the file such as the full path of the file can still be displayed in the statusline below:

	set previewer ~/.config/lf/pv.sh
	map i $LESSOPEN='| ~/.config/lf/pv.sh %s' less -R $f

Since this script is called for each file selection change it needs to be as efficient as possible and this responsibility is left to the user.
You may use file extensions to determine the type of file more efficiently compared to obtaining mime types from `file` command.
Extensions can then be used to match cleanly within a conditional:

	#!/bin/sh

	case "$1" in
	    *.tar*) tar tf "$1";;
	    *.zip) unzip -l "$1";;
	    *.rar) unrar l "$1";;
	    *.7z) 7z l "$1";;
	    *.pdf) pdftotext "$1" -;;
	    *) highlight -O ansi "$1";;
	esac

Another important consideration for efficiency is the use of programs with short startup times for preview.
For this reason, `highlight` is recommended over `pygmentize` for syntax highlighting.
Besides, it is also important that the application processes the file on the fly rather than first reading it to the memory and then doing the processing afterwards.
This is especially relevant for big files.
lf automatically closes the previewer script output pipe with a SIGPIPE when enough lines are read.
When everything else fails, you can make use of the height argument to only feed the first portion of the file to a program for preview.
Note that some programs may not respond well to SIGPIPE to exit with a non-zero return code and avoid caching.
You may add a trailing `|| true` command to avoid such errors:

	highlight -O ansi "$1" || true

You may also use an existing preview filter as you like.
Your system may already come with a preview filter named `lesspipe`.
These filters may have a mechanism to add user customizations as well.
See the related documentation for more information.

# CHANGING DIRECTORY

lf changes the working directory of the process to the current directory so that shell commands always work in the displayed directory.
After quitting, it returns to the original directory where it is first launched like all shell programs.
If you want to stay in the current directory after quitting, you can use one of the example lfcd wrapper shell scripts provided in the repository at
https://github.com/gokcehan/lf/tree/master/etc

There is a special command `on-cd` that runs a shell command when it is defined and the directory is changed.
You can define it just as you would define any other command:

	cmd on-cd &{{
	    bash -c '
	    # display git repository status in your prompt
	    source /usr/share/git/completion/git-prompt.sh
	    GIT_PS1_SHOWDIRTYSTATE=auto
	    GIT_PS1_SHOWSTASHSTATE=auto
	    GIT_PS1_SHOWUNTRACKEDFILES=auto
	    GIT_PS1_SHOWUPSTREAM=auto
	    git=$(__git_ps1 " (%s)")
	    fmt="\033[32;1m%u@%h\033[0m:\033[34;1m%d\033[0m\033[1m%f$git\033[0m"
	    lf -remote "send $id set promptfmt \"$fmt\""
	    '
	}}

If you want to send escape sequences to the terminal, you can use the `tty-write` command to do so.
The following xterm-specific escape sequence sets the terminal title to the working directory:

	cmd on-cd &{{
	    lf -remote "send $id tty-write \"\033]0;$PWD\007\""
	}}

This command runs whenever you change the directory but not on startup.
You can add an extra call to make it run on startup as well:

	cmd on-cd &{{ ... }}
	on-cd

Note that all shell commands are possible but `%` and `&` are usually more appropriate as `$` and `!` causes flickers and pauses respectively.

There is also a `pre-cd` command, that works like `on-cd`, but is run before the directory is actually changed.
Another related command is `on-load` which gets executed when loading a directory.

# LOADING DIRECTORY

Similar to `on-cd` there also is `on-load` that when defined runs a shell command after loading a directory.
It works well when combined with `addcustominfo`.

The following example can be used to display git indicators in the info column:

	cmd on-load &{{
	    cd "$(dirname "$1")" || exit 1
	    [ "$(git rev-parse --is-inside-git-dir 2>/dev/null)" = false ] || exit 0

	    cmds=""

	    for file in "$@"; do
	        case "$file" in
	            */.git|*/.git/*) continue;;
	        esac

	        status=$(git status --porcelain --ignored -- "$file" | cut -c1-2 | head -n1)

	        if [ -n "$status" ]; then
	            cmds="${cmds}addcustominfo \"${file}\" \"$status\"; "
	        else
	            cmds="${cmds}addcustominfo \"${file}\" ''; "
	        fi
	    done

	    if [ -n "$cmds" ]; then
	        lf -remote "send $id :$cmds"
	    fi
	}}

Another use case could be showing the dimensions of images and videos:

	cmd on-load &{{
	    cmds=""

	    for file in "$@"; do
	        mime=$(file --mime-type -Lb -- "$file")
	        case "$mime" in
	            # vector images cause problems
	            image/svg+xml)
	                ;;
	            image/*|video/*)
	                dimensions=$(exiftool -s3 -imagesize -- "$file")
	                cmds="${cmds}addcustominfo \"${file}\" \"$dimensions\"; "
	                ;;
	        esac
	    done

	    if [ -n "$cmds" ]; then
	        lf -remote "send $id :$cmds"
	    fi
	}}

# COLORS

lf tries to automatically adapt its colors to the environment.
It starts with a default color scheme and updates colors using values of existing environment variables possibly by overwriting its previous values.
Colors are set in the following order:

 1. default
 2. LSCOLORS (Mac/BSD ls)
 3. LS_COLORS (GNU ls)
 4. LF_COLORS (lf specific)
 5. colors file (lf specific)

Please refer to the corresponding man pages for more information about `LSCOLORS` and `LS_COLORS`.
`LF_COLORS` is provided with the same syntax as `LS_COLORS` in case you want to configure colors only for lf but not ls.
This can be useful since there are some differences between ls and lf, though one should expect the same behavior for common cases.
The colors file (refer to the [CONFIGURATION section](https://github.com/gokcehan/lf/blob/master/doc.md#configuration)) is provided for easier configuration without environment variables.
This file should consist of whitespace-separated pairs with a `#` character to start comments until the end of the line.

You can configure lf colors in two different ways.
First, you can only configure 8 basic colors used by your terminal and lf should pick up those colors automatically.
Depending on your terminal, you should be able to select your colors from a 24-bit palette.
This is the recommended approach as colors used by other programs will also match each other.

Second, you can set the values of environment variables or colors file mentioned above for fine-grained customization.
Note that `LS_COLORS/LF_COLORS` are more powerful than `LSCOLORS` and they can be used even when GNU programs are not installed on the system.
You can combine this second method with the first method for the best results.

Lastly, you may also want to configure the colors of the prompt line to match the rest of the colors.
Colors of the prompt line can be configured using the `promptfmt` option which can include hardcoded colors as ANSI escapes.
See the default value of this option to have an idea about how to color this line.

It is worth noting that lf uses as many colors advertised by your terminal's entry in terminfo or infocmp databases on your system.
If an entry is not present, it falls back to an internal database.
If your terminal supports 24-bit colors but either does not have a database entry or does not advertise all capabilities, you can enable support by setting the `$COLORTERM` variable to `truecolor` or ensuring `$TERM` is set to a value that ends with `-truecolor`.

Default lf colors are mostly taken from GNU dircolors defaults.
These defaults use 8 basic colors and bold attribute.
Default dircolors entries with background colors are simplified to avoid confusion with current file selection in lf.
Similarly, there are only file type matchings and extension matchings are left out for simplicity.
Default values are as follows given with their matching order in lf:

	ln  01;36
	or  31;01
	tw  01;34
	ow  01;34
	st  01;34
	di  01;34
	pi  33
	so  01;35
	bd  33;01
	cd  33;01
	su  01;32
	sg  01;32
	ex  01;32
	fi  00

Note that lf first tries matching file names and then falls back to file types.
The full order of matchings from most specific to least are as follows:

 1. Full Path (e.g. `~/.config/lf/lfrc`)
 2. Dir Name  (e.g. `.git/`) (only matches dirs with a trailing slash at the end)
 3. File Type (e.g. `ln`) (except `fi`)
 4. File Name (e.g. `README*`)
 5. File Name (e.g. `*README`)
 6. Base Name (e.g. `README.*`)
 7. Extension (e.g. `*.txt`)
 8. Default   (i.e. `fi`)

For example, given a regular text file `/path/to/README.txt`, the following entries are checked in the configuration and the first one to match is used:

 1. `/path/to/README.txt`
 2. (skipped since the file is not a directory)
 3. (skipped since the file is of type `fi`)
 4. `README.txt*`
 5. `*README.txt`
 6. `README.*`
 7. `*.txt`
 8. `fi`

Given a regular directory `/path/to/example.d`, the following entries are checked in the configuration and the first one to match is used:

 1. `/path/to/example.d`
 2. `example.d/`
 3. `di`
 4. `example.d*`
 5. `*example.d`
 6. `example.*`
 7. `*.d`
 8. `fi`

Note that glob-like patterns do not perform glob matching for performance reasons.

For example, you can set a variable as follows:

	export LF_COLORS="~/Documents=01;31:~/Downloads=01;31:~/.local/share=01;31:~/.config/lf/lfrc=31:.git/=01;32:.git*=32:*.gitignore=32:*Makefile=32:README.*=33:*.txt=34:*.md=34:ln=01;36:di=01;34:ex=01;32:"

Having all entries on a single line can make it hard to read.
You may instead divide it into multiple lines in between double quotes by escaping newlines with backslashes as follows:

	export LF_COLORS="\
	~/Documents=01;31:\
	~/Downloads=01;31:\
	~/.local/share=01;31:\
	~/.config/lf/lfrc=31:\
	.git/=01;32:\
	.git*=32:\
	*.gitignore=32:\
	*Makefile=32:\
	README.*=33:\
	*.txt=34:\
	*.md=34:\
	ln=01;36:\
	di=01;34:\
	ex=01;32:\
	"

The `ln` entry supports the special value `target`, which will use the link target to select a style. File name rules will still apply based on the link's name -- this mirrors GNU's `ls` and `dircolors` behavior.
Having such a long variable definition in a shell configuration file might be undesirable.
You may instead use the colors file (refer to the [CONFIGURATION section](https://github.com/gokcehan/lf/blob/master/doc.md#configuration)) for configuration.
A sample colors file can be found at
https://github.com/gokcehan/lf/blob/master/etc/colors.example
You may also see the wiki page for ANSI escape codes
https://en.wikipedia.org/wiki/ANSI_escape_code

# ICONS

Icons are configured using `LF_ICONS` environment variable or an icons file (refer to the [CONFIGURATION section](https://github.com/gokcehan/lf/blob/master/doc.md#configuration)).
The variable uses the same syntax as `LS_COLORS/LF_COLORS`.
Instead of colors, you should put a single characters as values of entries.
The `ln` entry supports the special value `target`, which will use the link target to select a icon. File name rules will still apply based on the link's name -- this mirrors GNU's `ls` and `dircolors` behavior.
The icons file (refer to the [CONFIGURATION section](https://github.com/gokcehan/lf/blob/master/doc.md#configuration)) should consist of whitespace-separated arrays with a `#` character to start comments until the end of the line.
Each line should contain 1-3 columns, first column is filetype or filename pattern, second column is the icon, third column is an optional icon color. If there is only one column, means to disable rule for this filetype or pattern.
Do not forget to add `set icons true` to your `lfrc` to see the icons.
Default values are as follows given with their matching order in lf:

	ln  l
	or  l
	tw  t
	ow  d
	st  t
	di  d
	pi  p
	so  s
	bd  b
	cd  c
	su  u
	sg  g
	ex  x
	fi  -

A sample icons file can be found at
https://github.com/gokcehan/lf/blob/master/etc/icons.example

A sample colored icons file can be found at
https://github.com/gokcehan/lf/blob/master/etc/icons_colored.example
