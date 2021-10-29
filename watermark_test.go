package go_watermark

import (
	"image"
	"testing"
)

func Test_mergeImg(t *testing.T) {
	type args struct {
		watermark image.Image
		original  image.Image
		position  image.Point
	}
	tests := []struct {
		name    string
		args    args
		want    image.Image
		wantErr bool
	}{
		{
			name: "no error expected",
			args: args{
				watermark: nil,
				original:  nil,
				position:  image.Pt(50, 50),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// import watermark & original image
			tt.args.watermark, _, _ = importImg("./test_image/spongebob_watermark.png")
			tt.args.original, _, _ = importImg("./test_image/spongebob.png")

			// setup watermark opacity
			transparentWatermark, _ := Opacity(tt.args.watermark, 0.5)

			// resizing watermark image
			b := transparentWatermark.Bounds()
			w := uint(b.Max.X / 4)
			h := uint(b.Max.Y / 4)

			resizeWatermark, _ := Resize(transparentWatermark, w, h)

			got, err := mergeImg(resizeWatermark, tt.args.original, &tt.args.position)
			if (err != nil) != tt.wantErr {
				t.Errorf("Watermark() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// export image result
			_ = exportImg(got, "", "./test_image/watermarkRes")
		})
	}
}
