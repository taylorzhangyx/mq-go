package main

import (
	"testing"
)

func TestMsgQue_AppendOne(t *testing.T) {
	type fields struct {
		count int32
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantVal interface{}
		wantErr bool
	}{
		{
			name:    "Append one element should add one element to the queue",
			fields:  fields{count: 1},
			args:    args{v: "success"},
			wantVal: "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MsgQue{
				count: tt.fields.count,
			}
			if err := q.AppendOne(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("MsgQue.AppendOne() error = %v, wantErr %v", err, tt.wantErr)
			}
			if val, err := q.ConsumeOne(); val != tt.wantVal || err != nil {
				t.Errorf("MsgQue.AppendOne() val = %v, wantVal %v", val, tt.wantVal)
			}

		})
	}
}
