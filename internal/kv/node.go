package kv

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	kv "github.com/galaco/KeyValues"
)

type Node struct {
	*kv.KeyValue
}

func NewNode(kv *kv.KeyValue) *Node {
	return &Node{KeyValue: kv}
}

func ParseFile(filename string) (*Node, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := kv.NewReader(f)
	root, err := reader.Read()

	if err != nil {
		return nil, err
	}

	return NewNode(&root), nil
}

func (n *Node) Children() ([]*Node, error) {
	kvChildren, err := n.KeyValue.Children()

	if err != nil {
		return nil, err
	}

	children := make([]*Node, len(kvChildren))

	for i, child := range kvChildren {
		children[i] = NewNode(child)
	}

	return children, nil
}

func (n *Node) Child(key string, optional bool) (*Node, error) {
	child, err := n.Find(key)

	if err != nil {
		if IsKeyMissingError(err) && optional {
			return nil, nil
		}

		return nil, err
	}

	return NewNode(child), nil
}

func (n *Node) TypedChild(key string, typ kv.ValueType, optional bool) (*Node, error) {
	child, err := n.Child(key, optional)

	if err != nil {
		return nil, err
	}

	if child == nil {
		return nil, nil
	}

	if child.Type() != typ {
		return nil, &InvalidTypeError{Key: key, Type: child.Type(), ExpectedType: typ}
	}

	return child, nil
}

func (n *Node) ChildAsInt32(key string, optional bool) (int32, error) {
	child, err := n.TypedChild(key, kv.ValueInt, optional)

	if err != nil {
		return 0, err
	}

	if child == nil {
		return 0, nil
	}

	num, err := child.AsInt()

	if err != nil {
		return 0, err
	}

	return num, nil
}

func (n *Node) ChildAsInt64(key string, optional bool) (int64, error) {
	num, err := n.ChildAsInt32(key, optional)

	if err != nil {
		return 0, err
	}

	return int64(num), nil
}

func (n *Node) ChildAsInt(key string, optional bool) (int, error) {
	num, err := n.ChildAsInt32(key, optional)

	if err != nil {
		return 0, err
	}

	return int(num), nil
}

func (n *Node) ChildAsString(key string, optional, trim bool) (string, error) {
	child, err := n.TypedChild(key, kv.ValueString, optional)

	if err != nil {
		return "", err
	}

	if child == nil {
		return "", nil
	}

	s, err := child.AsString()

	if err != nil {
		return "", err
	}

	if trim {
		s = strings.TrimSpace(s)
	}

	return s, nil
}

func (n *Node) ChildAsStringArray(key string, sep *regexp.Regexp, optional, trim, toLower bool) ([]string, error) {
	s, err := n.ChildAsString(key, optional, trim)

	if err != nil {
		return nil, err
	}

	if s == "" {
		return nil, nil
	}

	ss := sep.Split(s, -1)

	if trim || toLower {
		for i, s := range ss {
			if trim {
				s = strings.TrimSpace(s)
			}

			if toLower {
				s = strings.ToLower(s)
			}

			ss[i] = s
		}
	}

	return ss, nil
}

func (n *Node) ChildAsInt64Array(key string, sep *regexp.Regexp, optional bool) ([]int64, error) {
	child, err := n.Child(key, optional)

	if err != nil {
		return nil, err
	}

	switch t := child.Type(); t {
	case kv.ValueInt:
		num, err := child.AsInt()

		if err != nil {
			return nil, err
		}

		return []int64{int64(num)}, nil
	case kv.ValueString:
		s, err := child.AsString()

		if err != nil {
			return nil, err
		}

		if s == "" {
			return nil, nil
		}

		ss := sep.Split(s, -1)
		nn := make([]int64, len(ss))

		for i, s := range ss {
			s = strings.TrimSpace(s)
			nn[i], err = strconv.ParseInt(s, 10, 64)

			if err != nil {
				return nil, err
			}
		}

		return nn, nil
	default:
		return nil, &InvalidTypeError{Key: key, Type: t, ExpectedType: kv.ValueString}
	}
}
