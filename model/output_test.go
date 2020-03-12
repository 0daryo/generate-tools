package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Model_NewAIImage(t *testing.T) {
	type args struct{
	productID	int64
	imageURL	string
	order	int64
	category	string
	scaleAB	int64
	scaleBD	int64
	scaleDE	int64
	scaleEA	int64
	}
	tests := []struct {
		name   string
		args  args
		want  want
		wantErr  wantErr
	}{
		{
			name: "success",
						args: args{
				productID:	1,
				imageURL:	"testurl/jpeg",
			},
			want: &%s{ AIImage
				productID:	1,
				imageURL:	"testurl/jpeg",
			},
		},
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			got, err := New%s( AIImage
				tt.args.productID
				tt.args.imageURL
				tt.args.order
				tt.args.category
				tt.args.scaleAB
				tt.args.scaleBD
				tt.args.scaleDE
				tt.args.scaleEA
			)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
