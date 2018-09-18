package cmd

import (
	"strings"
	"testing"

	"github.com/qri-io/ioes"
	"github.com/qri-io/qri/lib"
)

func TestUseComplete(t *testing.T) {
	streams, in, out, errs := ioes.NewTestIOStreams()
	setNoColor(true)

	f, err := NewTestFactory(nil)
	if err != nil {
		t.Errorf("error creating new test factory: %s", err)
		return
	}

	cases := []struct {
		args   []string
		expect string
		err    string
	}{
		{[]string{}, "", ""},
		{[]string{"me/test"}, "me/test", ""},
		{[]string{"me/test", "me/test2"}, "me/test, me/test2", ""},
	}

	for i, c := range cases {
		opt := &UseOptions{
			IOStreams: streams,
		}

		opt.Complete(f, c.args)

		if c.err != errs.String() {
			t.Errorf("case %d, error mismatch. Expected: '%s', Got: '%s'", i, c.err, errs.String())
			ioReset(in, out, errs)
			continue
		}

		optRefs := strings.Join(opt.Refs, ", ")

		if c.expect != optRefs {
			t.Errorf("case %d, opt.Refs not set correctly. Expected: [%s], Got: [%s]", i, c.expect, optRefs)
			ioReset(in, out, errs)
			continue
		}

		if opt.SelectionRequests == nil {
			t.Errorf("case %d, opt.SelectionRequests not set.", i)
			ioReset(in, out, errs)
			continue
		}
		ioReset(in, out, errs)
	}
}

func TestUseValidate(t *testing.T) {
	cases := []struct {
		refs  []string
		list  bool
		clear bool
		err   string
		msg   string
	}{
		{[]string{}, false, false, lib.ErrBadArgs.Error(), "please provide dataset name, or --clear flag, or --list flag\nsee `qri use --help` for more info"},
		{[]string{"me/test"}, false, false, "", ""},
		{[]string{}, true, false, "", ""},
		{[]string{}, false, true, "", ""},
		{[]string{"me/test"}, true, true, "bad arguments provided", "please only give a dataset name, or a --clear flag, or  a --list flag"},
	}
	for i, c := range cases {
		opt := &UseOptions{
			Refs:  c.refs,
			List:  c.list,
			Clear: c.clear,
		}

		err := opt.Validate()
		if (err == nil && c.err != "") || (err != nil && c.err != err.Error()) {
			t.Errorf("case %d, mismatched error. Expected: '%s', Got: '%s'", i, c.err, err)
			continue
		}
		if libErr, ok := err.(lib.Error); ok {
			if libErr.Message() != c.msg {
				t.Errorf("case %d, mismatched user-friendly message. Expected: '%s', Got: '%s'", i, c.msg, libErr.Message())
				continue
			}
		} else if c.msg != "" {
			t.Errorf("case %d, mismatched user-friendly message. Expected: '%s', Got: ''", i, c.msg)
			continue
		}
	}
}

func TestUseRun(t *testing.T) {
	streams, in, out, errs := ioes.NewTestIOStreams()
	setNoColor(true)

	f, err := NewTestFactory(nil)
	if err != nil {
		t.Errorf("error creating new test factory: %s", err)
		return
	}

	cases := []struct {
		refs     []string
		list     bool
		clear    bool
		err      string
		expected string
	}{
		{[]string{"me/test1"}, false, false, "", "me/test1\n"},
		{[]string{"me/test2", "me/test3"}, false, false, "", "me/test2\nme/test3\n"},
		{[]string{}, true, false, "", "me/test2\nme/test3\n"},
		{[]string{}, false, true, "", "cleared selected datasets\n"},
		{[]string{}, true, false, "", ""},
	}

	for i, c := range cases {
		slr, err := f.SelectionRequests()
		if err != nil {
			t.Errorf("case %d, error creating dataset request: %s", i, err)
			continue
		}

		opt := &UseOptions{
			IOStreams:         streams,
			Refs:              c.refs,
			List:              c.list,
			Clear:             c.clear,
			SelectionRequests: slr,
		}

		err = opt.Run()
		if (err == nil && c.err != "") || (err != nil && c.err != err.Error()) {
			t.Errorf("case %d, mismatched error. Expected: '%s', Got: '%v'", i, c.err, err)
			ioReset(in, out, errs)
			continue
		}

		if c.expected != out.String() {
			t.Errorf("case %d, output mismatch. Expected: '%s', Got: '%s'", i, c.expected, out.String())
			ioReset(in, out, errs)
			continue
		}
		ioReset(in, out, errs)
	}
}
