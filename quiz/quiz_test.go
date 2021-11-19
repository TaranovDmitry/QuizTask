package quiz

import (
	"testing"

	"QuizTask/entity"

	"github.com/stretchr/testify/assert"
)

func Test_readQuizFromFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Quiz
		err     string
		wantErr bool
	}{
		{
			name: "test #1 Success",
			args: args{fileName: "./testdata/test_problems.json"},
			want: []entity.Quiz{
				{
					Question: "5+5",
					Answer:   "10",
				},
				{
					Question: "7+5",
					Answer:   "12",
				},
			},
			err:     "",
			wantErr: false,
		},
		{
			name:    "test #2 Fail to open file",
			args:    args{fileName: "test_problems.error"},
			want:    nil,
			err:     "failed to open file: open test_problems.error: The system cannot find the file specified.",
			wantErr: true,
		},
		{
			name:    "test #3 Fail to unmarshal",
			args:    args{fileName: "./testdata/problems_bad.json"},
			want:    nil,
			err:     "failed to unmarshal the file content ./testdata/problems_bad.json: invalid character '{' after array element",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadQuizFromFile(tt.args.fileName)
			if tt.wantErr {
				assert.Nil(t, got)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.Equal(t, tt.want, got)
				assert.NoError(t, err)
			}
		})
	}
}
