# thresher

A dumb cli utility that you give a theshold and a number and it tells if your number is bigger.

Don't ask why, it just became necessary before I lost my mind with Bash inconcistencies.

## Installing

```
go install github.com:cedrus-and-thuja/thresher/cmd/thresher@v1
```

## Usage

It's simple, you give it a theshold value that the input value should be greater than. if it is greater than the theshold it prints a message saying so and returns 0. If the value is not greater than the threshold then it prints a message saying so and returns non-zero (5). Parsing errors for the threshold or value return 1 or 2 respectively.

```
thresher -threshold <threshold> -value <value>
```

example:

```
thresher -threshold 8.98 -value 70.0
> value 70.00 is greater than threshold 8.98

thresher  -threshold 8.98 -value 7.5
> value 7.50 is less than threshold 8.98
> exit status 5

```
