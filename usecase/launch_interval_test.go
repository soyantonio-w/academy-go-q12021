package usecase

import "testing"

func Test_interval_getIntervals(t *testing.T) {
	type fields struct {
		iteration int
		items     int
		length    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  int
	}{
		{
			name: "Should group correctly first group",
			fields: fields{
				iteration: 0,
				items:     3,
				length:    6,
			},
			want:  0,
			want1: 3,
		},
		{
			name: "Should group correctly second group",
			fields: fields{
				iteration: 1,
				items:     4,
				length:    8,
			},
			want:  4,
			want1: 8,
		},
		{
			name: "Should group correctly last group",
			fields: fields{
				iteration: 2,
				items:     3,
				length:    3,
			},
			want:  3,
			want1: 3,
		},
		{
			name: "Should truncate correctly last group",
			fields: fields{
				iteration: 2,
				items:     3,
				length:    8,
			},
			want:  6,
			want1: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &interval{
				iteration: tt.fields.iteration,
				items:     tt.fields.items,
				length:    tt.fields.length,
			}
			got, got1 := i.getIntervals()
			if got != tt.want {
				t.Errorf("getIntervals() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getIntervals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
