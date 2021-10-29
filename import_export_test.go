package go_watermark

import (
	"image"
	"os"
	"testing"
)

func TestImport(t *testing.T) {
	type args struct {
		imageFile string
	}
	tests := []struct {
		name            string
		args            args
		wantImageObject image.Image
		wantImageType   string
		wantErr         bool
	}{
		{
			name:          "no error expected",
			args:          args{imageFile: "./test_image/spongebob.png"},
			wantImageType: "png",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//gotImageObject, gotFilename, err := Import(tt.args.imageFile)
			_, gotImageType, err := importImg(tt.args.imageFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Import() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotImageType != tt.wantImageType {
				t.Errorf("Import() gotImageType = %v, want %v", gotImageType, tt.wantImageType)
			}
		})
	}
}

func TestExport(t *testing.T) {
	type args struct {
		imageObject image.Image
		imageType   string
		fileName    string
	}
	tests := []struct {
		name    string
		args    args
		wantOut *os.File
		wantErr bool
	}{
		{
			name: "no error expected",
			args: args{
				imageObject: nil,
				imageType:   "",
				fileName:    "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			tt.args.imageObject, _, err = importImg("./test_image/spongebob.png")
			if err != nil {
				t.Error(err)
			}

			err = exportImg(tt.args.imageObject, tt.args.imageType, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Export() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
