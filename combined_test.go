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

	"github.com/orayew2002/json-bench/structures"
)

func BenchmarkGoccyJson(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			jsonData, err := goccy_json.Marshal(UserProfileData)
			if err != nil {
				return errors.Join(err, errors.New("goccy json marshal error"))
			}

			var newData structures.UserProfile
			if err = goccy_json.Unmarshal(jsonData, &newData); err != nil {
				return errors.Join(err, errors.New("goccy json unmarshal error"))
			}
		}

		return nil
	})
}

func BenchmarkEasyjson(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			jsonData, err := easyjson.Marshal(UserProfileData)
			if err != nil {
				return errors.Join(err, errors.New("easyjson marshal error"))
			}

			var newData structures.UserProfile
			if err = easyjson.Unmarshal(jsonData, &newData); err != nil {
				return errors.Join(err, errors.New("easyjson unmarshal error"))
			}
		}

		return nil
	})
}

func BenchmarkJsoniter(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			jsonData, err := jsoniter.Marshal(UserProfileData)
			if err != nil {
				return errors.Join(err, errors.New("jsoniter marshal error"))
			}

			var newData structures.UserProfile
			if err = jsoniter.Unmarshal(jsonData, &newData); err != nil {
				return errors.Join(err, errors.New("jsoniter unmarshal error"))
			}
		}

		return nil
	})
}

func BenchmarkSonicJSON(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			jsonData, err := sonic.Marshal(UserProfileData)
			if err != nil {
				return errors.Join(err, errors.New("sonic marshal error"))
			}

			var newData structures.UserProfile
			if err = sonic.Unmarshal(jsonData, &newData); err != nil {
				return errors.Join(err, errors.New("sonic unmarshal error"))
			}
		}

		return nil
	})
}

func BenchmarkJSON(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			jsonData, err := json.Marshal(UserProfileData)
			if err != nil {
				return errors.Join(err, errors.New("json marshal error"))
			}

			var newData structures.UserProfile
			if err = json.Unmarshal(jsonData, &newData); err != nil {
				return errors.Join(err, errors.New("json unmarshal error"))
			}
		}

		return nil
	})
}

func BenchmarkFFJSON(b *testing.B) {
	RunBenchMark(b, func(b *testing.B) error {
		for i := 0; i < b.N; i++ {
			jsonData, err := ffjson.Marshal(UserProfileData)
			if err != nil {
				return errors.Join(err, errors.New("ffjson marshal error"))
			}

			var newData structures.UserProfile
			if err = ffjson.Unmarshal(jsonData, &newData); err != nil {
				return errors.Join(err, errors.New("ffjson unmarshal error"))
			}
		}

		return nil
	})
}

func RunBenchMark(b *testing.B, runner func(b *testing.B) error) {
	b.ResetTimer()

	if err := runner(b); err != nil {
		b.Fatal(err)
	}
}
