# git-email-collector

Collects emails from public git repo. Emails will be printed to stdout along with their names and commit messages to confirm if they are good target for cold calling.

It accepts 2 arguments:

1. Git https address
2. Number of prints per email. Multiple prints are recommended for guessing if the email author is actually doing something, just merging, or else.

## Example

```bash
go run . https://github.com/gatsu420/git-email-collector.git 5
```
