# clockify-sync

`clockify-sync` is a command-line tool that generates `clockify-cli`
and other tool commands for synchronizing time entries between the systems. 
The tool reads a JSON-formatted Clockify time entry from standard input 
and generates commands:

- A `clockify-cli edit` command to add a tag to the time entry indicating 
that it has been synced with the target system.
- A command for the specified target system to add a time entry to the specified issue.

The tool is not designed to handle the actual syncing of the entries itself, 
as this would require the user to provide two API keys: one for Clockify
and one for the target work platform, i.e. YouTrack, Jira, etc.

Instead, `clockify-sync` generates commands that the user can 
review and execute manually, or pipe directly to `sh` to execute automatically. 
This approach allows the user to review and verify the commands before they are executed, 
which can help prevent accidental or unintended changes to the time entries.

## Installation

You can install `clockify-sync` using the `go install` command:

`go install github.com/biecho/clockify-sync`

## Usage

To use `clockify-sync`, you must provide a subcommand for the target system 
to sync with, along with any necessary flags or parameters.
The currently supported subcommand is `youtrack`.

```
clockify-sync youtrack \
    --project-mapping <mapping-file> 
    --issue-id-regex <regex-pattern>
    [--ignore-tag <tag>]... 
    [--sync-tag <tag>]...
```

Here's what each flag does:

- `--project-mapping`: Specifies a file containing a mapping of Clockify project 
names to YouTrack work item types. The mapping file should be a JSON file with 
the following structure:

```json
{
    "project1": "workItemType1",
    "project2": "workItemType2"
}
```

- `--issue-id-regex`: Specifies a regex pattern to match against the Clockify 
entry description. The first matches defines the YouTrack issue ID. For example, 
the pattern YT-\d+ would match any string that starts with "YT-" followed by one 
or more digits, and would capture the digits as the issue ID.

- `--ignore-tag`: Specifies a tag to ignore. Clockify entries that have 
this tag will not be synced with YouTrack. This flag can be specified multiple
times to ignore multiple tags.

- `--sync-tag`: Specifies a tag to add to all synced Clockify entries.
This flag can be specified multiple times to add multiple tags.

Here's an example of how to use `clockify-youtrack-sync`:

```bash
clockify-cli report today -j \
   | clockify-sync youtrack \ 
   --project-mapping mapping.json \
   --issue-id-regex "YT-\d+" \
   --ignore-tag "ignore-this-tag" \
   --sync-tag "synced" 
```

This would read the Clockify entries (in JSON format) from STDIN, map its project name 
to the corresponding YouTrack work item type using the mapping in `mapping.json`,
ignore any entries with the tag "ignore-this-tag", add the tag "synced" 
to all synced entries, and match the YouTrack issue ID using the pattern `YT-\d+`.

## Supported Systems

The following systems are currently supported:

### YouTrack

The `youtrack` subcommand generates two commands:

- A `clockify-cli edit` command to add a tag to the time entry indicating 
that it has been synced with YouTrack. The command is written to standard output 
and can be executed by piping it to `sh`, for example.

For example:

`clockify-cli edit 642c00b3fcde1e287f9af847 --tag "AddedToYouTrack"`

- A `yt add work` command to add a work item to the specified YouTrack issue. 

For example:

`yt add work --issue-id YT-725 --type Development --date 2023-04-14 --duration 1h10m`

### Jira (planned)

Jira support is planned for a future release of `clockify-sync`.