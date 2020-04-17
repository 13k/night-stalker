package cmddebug

import (
	"encoding/json"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
	d2pb "github.com/paralin/go-dota2/protocol"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
	"google.golang.org/protobuf/encoding/protojson"

	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nsjson "github.com/13k/night-stalker/internal/json"
)

var CmdTransformJSON = &cobra.Command{
	Use:   "transform_json <input.json>",
	Short: "Debug JSON transformation",
	RunE:  debugTransformJSON,
}

func init() {
	Cmd.AddCommand(CmdTransformJSON)
}

func debugTransformJSON(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	if len(args) < 1 {
		return cmd.Usage()
	}

	data, err := ioutil.ReadFile(args[0])

	if err != nil {
		return xerrors.Errorf("error reading inupt file: %w", err)
	}

	t := &transform{data: data}

	msg, err := t.Parse()

	if err != nil {
		if e := (&errGeneratedInvalidJSON{}); xerrors.As(err, &e) {
			var m map[string]interface{}

			if err = json.Unmarshal(e.gen, &m); err != nil {
				log.WithError(err).Error("cannot parse generated JSON")
				return err
			}

			dumpf("gen map: %s\n", m)
		}

		return xerrors.Errorf("error transforming JSON: %w", err)
	}

	dumpf("msg: %s\n", msg)

	return nil
}

type transform struct {
	data []byte
}

func (r *transform) Parse() (*d2pb.CMsgDOTARealtimeGameStatsTerse, error) {
	b, err := r.transform()

	if err != nil {
		return nil, xerrors.Errorf("error transforming data: %w", err)
	}

	if b == nil {
		return nil, nil
	}

	pbmsg := &d2pb.CMsgDOTARealtimeGameStatsTerse{}

	if err := protojson.Unmarshal(b, pbmsg); err != nil {
		return nil, xerrors.Errorf("error unmarshaling transformed response: %w", err)
	}

	return pbmsg, nil
}

func (r *transform) transform() ([]byte, error) {
	t, err := nsjson.NewTransform(r.data)

	if err != nil {
		return nil, xerrors.Errorf("error creating JSON transform: %w", err)
	}

	if t.ReadNil() {
		return nil, nil
	}

	// root
	b, err := t.Object(func(key string) ([]byte, error) {
		if key != "teams" {
			return t.SkipAndReturnBytes(), nil
		}

		// ignore non-array "teams" field
		if typ := t.WhatIsNext(); typ != jsoniter.ArrayValue {
			return nil, nil
		}

		// teams array
		return t.Array(func() ([]byte, error) {
			// ignore non-object teams items
			if typ := t.WhatIsNext(); typ != jsoniter.ObjectValue {
				return nil, nil
			}

			// team object
			return t.Object(func(key string) ([]byte, error) {
				if key != "players" {
					return t.SkipAndReturnBytes(), nil
				}

				// ignore non-array "players" field
				if typ := t.WhatIsNext(); typ != jsoniter.ArrayValue {
					return nil, nil
				}

				// players array
				return t.Array(func() ([]byte, error) {
					// skip bogus player (usually an empty array)
					if typ := t.WhatIsNext(); typ != jsoniter.ObjectValue {
						return nil, nil
					}

					return t.SkipAndReturnBytes(), nil
				})
			})
		})
	})

	if err != nil {
		return nil, xerrors.Errorf("error transforming JSON: %w", err)
	}

	if !jsoniter.Valid(b) {
		return nil, xerrors.Errorf("error transforming JSON: %w", &errGeneratedInvalidJSON{gen: b})
	}

	return b, nil
}

type errGeneratedInvalidJSON struct {
	gen []byte
}

func (*errGeneratedInvalidJSON) Error() string {
	return "generated invalid JSON"
}
