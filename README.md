# Who needs regular expressions when you can say it in plain english?

`gpt-do` is a CLI utility to which you can pass data, tell gpt what you'd like done
with it, and be returned the tweaked data, feeling much like grep or sed.

The instructions are in plain English, and are sent to the openapi server along with
your data, which you pipe in via the CLI.

## Example:

gpt-do is invoked via the `gpt-do` command.

```
$ cat test/dates.txt 
Oct
January 1st, 1989
foobarr
bazinga, feb 10 2010, yes
10/16/2020
January 1st, 2001

$ cat test/dates.txt | ./gpt-do "increment every date that's after the year 2000 by one month" 


Oct
January 1st, 1989
foobarr
bazinga, mar 10 2010, yes
11/16/2020
February 1st, 2001
```

Note that, as asked, only those dates after 2000 are modified.

## Notes

Is this:
- safe to use with my data?
- cost-effecient?
- better than writing a regex, using sed, etc.?
- gauranteed to work correctly, predictably, and quickly?

Hell no. But it is neat.

## Installation

    go install github.com/marzhall/gpt-do/cmd/gpt-do@latest

This should install the binary `gpt-do` into your go bin directory (usually $HOME/go/bin).

## Setup

In order to function, `gpt-do` needs you to have an openapi API key so it can talk to the
gpt 3.5 chat API.

If you need an openapi api key, you'll want to create an account at https://platform.openai.com.
After logging in, you can get your API key at [this page](https://platform.openai.com/account/api-keys).

Once you have your API key, gpt-go needs to be able to find it. This can be done a few ways.
One is to just set your API key as an environment variable in your shell: `export API_KEY=<my_key>`.
From then on, `gpt-do` will pick up that environment variabe when it runs.

Alternatively, you can create a file named `.env` containing your openapi api key, and place it
either in your home directory or in the current directory you're running `gpt-do` in. More info
for using an env file is in the documentation for the [gpt-3 golang library](https://github.com/PullRequestInc/go-gpt3).


## Far-out notes, man

I made this because people are asking gpt to write code that does things instead of just asking
gpt to do things.

The future will be you asking your computer to do things in natural english. GPT-do is meant to
be a first exploration of that.

## Openapi-specific notes

The command uses gpt-3.5, as that's the best api available at the moment.

All commands will use your account via the API key, and costs will be accrued as expected.
The maximum number of tokens you'll receive in a reply is slightly more than twice the total
length of your instructions + the amount of data you sent. In the example, it's 2x the length
of dates.txt and the instruction " increment every date that's after the year 2000 by one month"
combined.
