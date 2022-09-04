# Jira CLI Helper Tool

This CLI app is intended to help me to organize my work jira tasks. It
is going to be a pretty personalized app so it may not fit your
requirements.

## Install

This command is using rwxrob/bonzai library to 

It can be installed as a standalone program or composed into a Bonzai
command tree.

Standalone

```
go install github.com/espinosajuanma/jira/cmd/jira@latest
```

Composed

```
package cmds

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/twitch"
)

var Cmd = &bonzai.Cmd{
	Name:     `cmds`,
	Commands: []*bonzai.Cmd{help.Cmd, twitch.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C twitch twitch
```

If you don't have bash or tab completion check use the shortcut commands
instead.
