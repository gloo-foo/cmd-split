package command

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	yup "github.com/gloo-foo/framework"
)

type command yup.Inputs[string, flags]

func Split(parameters ...any) yup.Command {
	cmd := command(yup.Initialize[string, flags](parameters...))
	if cmd.Flags.Lines == 0 && cmd.Flags.Bytes == 0 && cmd.Flags.Size == "" {
		cmd.Flags.Lines = 1000
	}
	if cmd.Flags.Prefix == "" {
		cmd.Flags.Prefix = "x"
	}
	if cmd.Flags.SuffixLength == 0 {
		cmd.Flags.SuffixLength = 2
	}
	return cmd
}

func (p command) Executor() yup.CommandExecutor {
	return func(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
		// Get input source (file or stdin)
		var input io.Reader = stdin
		if len(p.Positional) > 0 {
			file, err := os.Open(p.Positional[0])
			if err != nil {
				_, _ = fmt.Fprintf(stderr, "split: %s: %v\n", p.Positional[0], err)
				return err
			}
			defer file.Close()
			input = file
		}

		// Get prefix (default "x" or from second positional arg)
		prefix := string(p.Flags.Prefix)
		if len(p.Positional) > 1 {
			prefix = p.Positional[1]
		}

		scanner := bufio.NewScanner(input)

		var currentFile *os.File
		var currentWriter *bufio.Writer
		fileIndex := 0
		linesInFile := 0
		linesPerFile := int(p.Flags.Lines)

		// Helper to get next filename
		nextFilename := func() string {
			// Generate suffix (aa, ab, ac, ..., az, ba, bb, ...)
			suffixLen := int(p.Flags.SuffixLength)
			suffix := make([]byte, suffixLen)

			idx := fileIndex
			for i := suffixLen - 1; i >= 0; i-- {
				suffix[i] = byte('a' + (idx % 26))
				idx /= 26
			}

			return prefix + string(suffix)
		}

		// Helper to open next file
		openNextFile := func() error {
			// Close previous file
			if currentFile != nil {
				if err := currentWriter.Flush(); err != nil {
					return err
				}
				if err := currentFile.Close(); err != nil {
					return err
				}
			}

			// Open new file
			filename := nextFilename()
			var err error
			currentFile, err = os.Create(filename)
			if err != nil {
				return err
			}
			currentWriter = bufio.NewWriter(currentFile)

			if bool(p.Flags.Verbose) {
				_, _ = fmt.Fprintf(stdout, "creating file '%s'\n", filename)
			}

			fileIndex++
			linesInFile = 0

			return nil
		}

		// Open first file
		if err := openNextFile(); err != nil {
			_, _ = fmt.Fprintf(stderr, "split: %v\n", err)
			return err
		}

		// Process input line by line
		for scanner.Scan() {
			line := scanner.Text()

			// Check if we need to open a new file
			if linesPerFile > 0 && linesInFile >= linesPerFile {
				if err := openNextFile(); err != nil {
					_, _ = fmt.Fprintf(stderr, "split: %v\n", err)
					return err
				}
			}

			// Write line to current file
			_, err := fmt.Fprintln(currentWriter, line)
			if err != nil {
				_, _ = fmt.Fprintf(stderr, "split: %v\n", err)
				return err
			}

			linesInFile++
		}

		// Close final file
		if currentFile != nil {
			if err := currentWriter.Flush(); err != nil {
				return err
			}
			if err := currentFile.Close(); err != nil {
				return err
			}
		}

		return scanner.Err()
	}
}
