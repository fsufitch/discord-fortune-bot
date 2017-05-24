package bot

import (
	"bytes"
	"errors"
	"flag"

	"github.com/fsufitch/discord-fortune-bot/fortune"
)

const helpFooter = `
Anything after "--" will be passed through to the fortune command.

Check the Github repo for more info: https://github.com/fsufitch/discord-fortune-bot
`

type botOptions struct {
	Offensive    bool
	Length       fortune.Length
	Passthrough  []string
	TextOverride string
}

func parseFlags(flags []string) (options botOptions, err error) {
	var (
		longFlag       bool
		allLengthsFlag bool
		offensiveFlag  bool
		threshold      int
		passthrough    = []string{}
	)

	for index, f := range flags {
		if f == "--" {
			passthrough = flags[index+1:]
			flags = flags[:index]
			break
		}
	}

	flagBuffer := new(bytes.Buffer)
	flagSet := flag.NewFlagSet("fortune", flag.ContinueOnError)
	flagSet.SetOutput(flagBuffer)
	flagSet.BoolVar(&longFlag, "long", false, "long fortunes only (default: short only)")
	flagSet.BoolVar(&allLengthsFlag, "allLengths", false, "all fortune lengths (default: short only)")
	flagSet.BoolVar(&offensiveFlag, "offensive", false, "offensive fortunes only")
	flagSet.IntVar(&threshold, "length", 160, "length threshold for short/long decision")

	err = flagSet.Parse(flags)
	if err != nil {
		if err == flag.ErrHelp {
			options.TextOverride = string(flagBuffer.Bytes())
			options.TextOverride += helpFooter
		}
		return
	}

	if longFlag && allLengthsFlag {
		err = errors.New("Cannot use both -long and -allLengths")
		return
	}
	if longFlag {
		options.Length = fortune.Long
	} else if allLengthsFlag {
		options.Length = fortune.All
	}

	options.Offensive = offensiveFlag
	options.Passthrough = passthrough

	return
}
