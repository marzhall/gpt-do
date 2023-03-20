# GPT-do

GPT-do is a CLI utility to be used like grep or sed, except the instructions are
plain English instructions, and those instructions are sent to the gpt-3 server
along with your data, which you pipe in via the CLI.

Who needs regular expressions when you can just ask what you want in english?*

\* Other than people who need speed, accuracy, cost-effectiveness, security,
etc.

## Example:

GPT-do is invoked via the `gpt` command.

```
$ cat test/dates.txt 
Oct
January 1st, 1989
foobarr
bazinga, feb 10 2010, yes
10/16/2020
January 1st, 2001

$ cat test/dates.txt | ./gpt "increment every date that's after the year 2000 by one month" 


Oct
January 1st, 1989
foobarr
bazinga, mar 10 2010, yes
11/16/2020
February 1st, 2001
```

Note that only those dates after 2000 are modified.

## Requirements

`gpt` expects a file named `.env` containing your openapi api key to be either in your home
directory or in the current working directly; alternatively, you can have an environment variable
containing your API key called `API_KEY`. More info is defined in the documentation for the [gpt-3 golang library](https://github.com/PullRequestInc/go-gpt3).

If you need an openapi api key, you'll want to create an account at https://platform.openai.com.
After logging in, you can get your API key at [this page](https://platform.openai.com/account/api-keys)

All commands will use your account via the API key, and costs will be accrued as expected.
The maximum number of tokens you'll receive in a reply is slightly more than twice the total
length of your instructions + the amount of data you sent. In the example, it's 2x the length
of dates.txt and the instruction " increment every date that's after the year 2000 by one month"
combined.

## Notes

Is this:
- safe to use with my data?
- cost-effecient?
- better than writing a regex, using sed, etc.?
- gauranteed to work correctly, predictably, and quickly?

Hell no. But it is neat.

## Far-out notes, man

I made this because people are asking gpt to write code to do things, instead of just asking
gpt to do the things that need done.

The future will be you asking your computer to do things in natural english. GPT-do is meant to
be a first exploration of that.