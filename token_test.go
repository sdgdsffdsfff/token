package token

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want *Token
	}{
		{
			name: "TestNew",
			args: args{
				key: []byte("key"),
			},
			want: &Token{version: CurrentVesion, key: []byte("key")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenSign(t *testing.T) {
	token := New([]byte("key"))
	assert.NotNil(t, token)
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestSignCase1",
			args: args{
				data: []byte("TestSignCase1"),
			},
		},
		{
			name: "TestSignCase2",
			args: args{
				data: []byte("TestSignCase2"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := token.Sign(tt.args.data)
			assert.NotNil(t, got)
			assert.NoError(t, err)

			err = token.Verify(got)
			assert.NoError(t, err)
		})
	}
}

func TestTokenVerify(t *testing.T) {
	token := New([]byte("key"))
	assert.NotNil(t, token)
	type args struct {
		sign []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestVerifyCase1",
			args: args{
				sign: []byte("TestVerifyCase1"),
			},
		},
		{
			name: "TestVerifyCase2",
			args: args{
				sign: []byte("TestVerifyCase2"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sign, err := token.Sign(tt.args.sign)
			assert.NotNil(t, sign)
			assert.NoError(t, err)
			err = token.Verify(sign)
			assert.NoError(t, err)
		})
	}
}

func Test_message_MarshalBinary(t *testing.T) {
	token := New([]byte("key"))
	assert.NotNil(t, token)
	tests := []struct {
		name     string
		payload  []byte
		wantData string
		wantErr  bool
	}{
		{
			name:     "TestMarshalBinary",
			payload:  []byte("TestMarshalBinary"),
			wantData: "TestMarshalBinary-1545205200-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				version:  CurrentVesion,
				createAt: 1545205200,
				payload:  tt.payload,
			}
			gotData, err := m.MarshalBinary()
			assert.NotNil(t, gotData)
			assert.NoError(t, err)

			assert.Equal(t, string(gotData), tt.wantData)
		})
	}

}