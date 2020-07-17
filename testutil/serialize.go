package testutil

import (
	"encoding"
	"github.com/stretchr/testify/require"
	"testing"
)

type Message interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	Len() int
	Append([]byte) ([]byte, error)
}

func Create(t testing.TB, f func() (Message, []byte)) {
	testM, testData := f()

	TestAll(t, func(zero bool) (m Message) {
		return testM
	}, testData)
}

func TestAll(t testing.TB, constructor func(zero bool) Message, testData []byte) {
	t.Helper()

	switch v := t.(type) {
	case *testing.T:
		TestSerialize(v, constructor, testData)
	case *testing.B:
		BenchSerialize(v, constructor, testData)
	}
}

func TestSerialize(t *testing.T, constructor func(zero bool) Message, testData []byte) {
	t.Helper()
	m := constructor(false)

	t.Run("marshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := m.MarshalBinary()
		assertions.NoError(err)
		assertions.Equal(testData, data)
	})

	t.Run("unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		m2 := constructor(true)
		err := m2.UnmarshalBinary(testData)
		assertions.NoError(err)
		assertions.Equal(m, m2)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := m.MarshalBinary()
		assertions.NoError(err)

		m2 := constructor(true)
		err = m2.UnmarshalBinary(data)
		assertions.NoError(err)

		assertions.Equal(m, m2)
	})
}