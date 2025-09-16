package main

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/bytedance/sonic"
	goccy_json "github.com/goccy/go-json"
	"github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/pquerna/ffjson/ffjson"
)

func BenchmarkMarshalGoccyJson(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			if _, err := goccy_json.Marshal(UserProfileData); err != nil {
				return errors.Join(err, errors.New("goccy json marshal error"))
			}
		}

		return nil
	})
}

func BenchmarkMarshalEasyjson(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			if _, err := easyjson.Marshal(UserProfileData); err != nil {
				return errors.Join(err, errors.New("easyjson marshal error"))
			}
		}

		return nil
	})
}

func BenchmarkMarshalJsoniter(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			if _, err := jsoniter.Marshal(UserProfileData); err != nil {
				return errors.Join(err, errors.New("jsoniter marshal error"))
			}
		}

		return nil
	})
}

func BenchmarkMarshalSonicJSON(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			if _, err := sonic.Marshal(UserProfileData); err != nil {
				return errors.Join(err, errors.New("sonic marshal error"))
			}
		}

		return nil
	})
}

func BenchmarkMarshalJSON(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			if _, err := json.Marshal(UserProfileData); err != nil {
				return errors.Join(err, errors.New("json marshal error"))
			}
		}

		return nil
	})
}

func BenchmarkMarshalFFJSON(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			if _, err := ffjson.Marshal(UserProfileData); err != nil {
				return errors.Join(err, errors.New("ffjson marshal error"))
			}
		}

		return nil
	})
}
