package go_watermark

import (
	"image"
	"testing"
)

func TestOpacity(t *testing.T) {
	type args struct {
		imageObject image.Image
		alpha       float64
	}
	tests := []struct {
		name    string
		args    args
		wantImg image.Image
		wantErr bool
	}{
		{
			name: "no error expected",
			args: args{
				imageObject: nil,
				alpha:       0.5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// import watermark & original image
			tt.args.imageObject, _, _ = Import("./test_image/spongebob_watermark.png")

			gotImg, err := Opacity(tt.args.imageObject, tt.args.alpha)
			if (err != nil) != tt.wantErr {
				t.Errorf("Opacity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// export image result
			_ = Export(gotImg, "", "./test_image/opacityRes")
		})
	}
}
