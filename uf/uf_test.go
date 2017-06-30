package uf

import "testing"

func TestUF(t *testing.T) {
	uf := NewUF(6)
	if uf.Count() != 6 {
		t.Error("count not ok")
	}
}
func TestUF_Union(t *testing.T) {
	type fields struct {
		id []int
		sz []int
	}
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "7 case",
			fields: fields{
				id: []int{0, 1, 2, 3, 4, 5, 6},
				sz: []int{1, 1, 1, 1, 1, 1, 1},
			},
			args: args{
				p: 3,
				q: 4,
			},
		},
		{
			name: "1 case",
			fields: fields{
				id: []int{0},
				sz: []int{1},
			},
			args: args{
				p: 0,
				q: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := UF{
				id: tt.fields.id,
				sz: tt.fields.sz,
			}
			uf.Union(tt.args.p, tt.args.q)
			if uf.Connected(tt.args.p, tt.args.q) {
				t.Log("ok!!")
			} else {
				t.Error("fuck")
			}
		})
	}
}
