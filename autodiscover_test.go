package muacs

import (
	"reflect"
	"testing"
)

func TestAutodiscoverResponseXML(t *testing.T) {
	type args struct {
		r ADResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"Empty Response",
			args{ADResponse{}},
			[]byte(`<Autodiscover xmlns="http://schemas.microsoft.com/exchange/autodiscover/responseschema/2006"><Response xmlns="http://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a"></Response></Autodiscover>`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AutodiscoverResponseXML(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Response() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response() = %s, want %s", got, tt.want)
			}
		})
	}
}
