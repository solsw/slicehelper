package slicehelper

import (
	"reflect"
	"testing"
)

func TestReverseNew_int(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "nil",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{name: "0",
			args: args{
				s: []int{},
			},
			want: []int{},
		},
		{name: "1",
			args: args{
				s: []int{1},
			},
			want: []int{1},
		},
		{name: "12",
			args: args{
				s: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{name: "123",
			args: args{
				s: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseNew(tt.args.s)
			if len(got) > 0 {
				if &got[0] == &tt.args.s[0] {
					t.Errorf("same slice: initial: %p, got: %p", tt.args.s, got)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortDesc_int(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "nil",
			args: args{
				x: nil,
			},
			want: nil,
		},
		{name: "0",
			args: args{
				x: []int{},
			},
			want: []int{},
		},
		{name: "1",
			args: args{
				x: []int{},
			},
			want: []int{},
		},
		{name: "8",
			args: args{
				x: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: []int{8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortDesc(tt.args.x)
			if !reflect.DeepEqual(tt.args.x, tt.want) {
				t.Errorf("SortDesc() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		len int
		n   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "01",
			args: args{
				len: 1,
				n:   1,
			},
			wantErr: true,
		},
		{name: "02",
			args: args{
				len: 2,
				n:   1,
			},
			wantErr: true,
		},
		{name: "03",
			args: args{
				len: 2,
				n:   3,
			},
			wantErr: true,
		},
		{name: "2",
			args: args{
				len: 8,
				n:   2,
			},
			want: []int{0, 4, 8},
		},
		{name: "3",
			args: args{
				len: 8,
				n:   3,
			},
			want: []int{0, 3, 6, 8},
		},
		{name: "4",
			args: args{
				len: 8,
				n:   4,
			},
			want: []int{0, 2, 4, 6, 8},
		},
		{name: "5",
			args: args{
				len: 8,
				n:   5,
			},
			want: []int{0, 2, 4, 6, 7, 8},
		},
		{name: "6",
			args: args{
				len: 8,
				n:   6,
			},
			want: []int{0, 2, 4, 5, 6, 7, 8},
		},
		{name: "7",
			args: args{
				len: 8,
				n:   7,
			},
			want: []int{0, 2, 3, 4, 5, 6, 7, 8},
		},
		{name: "8",
			args: args{
				len: 8,
				n:   8,
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Split(tt.args.len, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveInPlace_int(t *testing.T) {
	type args struct {
		s   []int
		idx int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "1e",
			args: args{
				s: nil,
			},
			wantErr: true,
		},
		{name: "2e",
			args: args{
				s: []int{},
			},
			wantErr: true,
		},
		{name: "3e",
			args: args{
				s:   []int{1, 2},
				idx: 2,
			},
			wantErr: true,
		},
		{name: "1",
			args: args{
				s: []int{1},
			},
			want: []int{},
		},
		{name: "2",
			args: args{
				s: []int{1, 2},
			},
			want: []int{2},
		},
		{name: "3",
			args: args{
				s:   []int{1, 2},
				idx: 1,
			},
			want: []int{1},
		},
		{name: "4",
			args: args{
				s:   []int{1, 2, 3},
				idx: 1,
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveInPlace(tt.args.s, tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveInPlace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) > 0 {
				if &got[0] != &tt.args.s[0] {
					t.Errorf("not same slice: initial: %p, got: %p", tt.args.s, got)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
