package cmdapi

import (
	"bytes"
	"strings"

	"github.com/13k/night-stalker/cmd/modelgen/internal/common"
	mg "github.com/13k/night-stalker/cmd/modelgen/internal/parser"
)

type genOpts struct {
	*Options

	p     *mg.Package
	s     *mg.Struct
	fname string
}

func gen(opts *genOpts) error {
	s, p, fname := opts.s, opts.p, opts.fname

	recv := s.GetFirstRecvName(opts.DefaultRecvName)
	rels, err := parseModelAssociations(s)

	if err != nil {
		return err
	}

	b := &bytes.Buffer{}

	if err = execTplAPI(b, p, s, recv, rels); err != nil {
		return err
	}

	formatted, err := common.FormatCode(b.Bytes())

	if err != nil {
		return err
	}

	if opts.Output == "-" {
		opts.Cmd.PrintErrln(strings.Repeat("-", 80))
	}

	opts.Cmd.PrintErrln(fname)

	if opts.Output == "-" {
		opts.Cmd.PrintErrln(strings.Repeat("-", 80))
		opts.Cmd.Println(string(formatted))
		return nil
	}

	return common.CreateFile(fname, formatted)
}
