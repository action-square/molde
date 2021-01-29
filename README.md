![Molde](https://live.staticflickr.com/65535/50878282997_a819a0a40f_b.jpg)
[![Go](https://github.com/action-square/molde/workflows/Go/badge.svg)](https://github.com/action-square/molde/actions)
[![Docker](https://github.com/action-square/molde/workflows/Docker/badge.svg)](https://github.com/action-square/molde/actions)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/action-square/molde)](https://github.com/action-square/molde/releases)
[![License](https://img.shields.io/badge/license-MIT-informational.svg)](https://opensource.org/licenses/MIT)

A simple command line mail template generator and sender using **Sass** and **Markdown**.
This tool can send customized mails using data from a single *JSON* file throught any mail provider.

## Usage

To use this tool you can either run it or build it with `go`. To build it, use:

```sh
go get -d -v ./...
go build -o ./dist/molde -v ./cmd/molde/
```

Type `-help` on the terminal for a list of options.
You can take a look at the `sample` directory or look at the test files for examples.

To send the mails you'll need to make a `layout` file with the base **HTML** for the template and two tags `{% css %}`, for the generated **CSS**, using **Sass**; and `{% content %}`, for the generated **HTML**, using **Markdown**.
The `content` file can include any number of tags, formatted as `{{ flag }}`, as long as there's a correponding field in the `data` file.
Each mail's tags can have different values in the `data` file, but there must always be the `to` and `subject` fields.

## Contributions

Feel free to leave your contribution here, I would really appreciate it!
Also, if you have any doubts or troubles using this tool just contact me or leave an issue.
