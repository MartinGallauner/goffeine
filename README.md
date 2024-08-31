# goffeine

## About
Goffeine is a simple caffeine tracking tool, running in your terminal.
It allows you to add your caffeine consumption in natural language. 
Also, you can prompt your estimated caffeine levels including the expected breakdown of caffeine in your body.
It offers you 3 different commands:
- `help` will show you all available commands with a little description
- `status` shows you your Caffeine level for right now.
- `add {what you consumed when in simple language}` passes your input to openai to find a reasonable timestamp and amount of caffeine consumed.

## Installation

Warning: So far, I put very limited effort into this project and therefore, usability is not where it may be one day.

1. [Install Go](https://go.dev/doc/install)
2. Run `go install github.com/MartinGallauner/goffeine`
3. Run `export OPENAI_API_KEY={your openai key}` 
4. Add file `$GO_PATH/bin/data.csv` 


## Ideas for further improvement
- Deploy as webservice
- Allow deletion of entries
