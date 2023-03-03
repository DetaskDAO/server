package utils

import (
	"testing"
)

//func Test_ipfsRequest(t *testing.T) {
//
//	gotSpent, err := ipfsRequest("http://ipfs.learnblockchain.cn", "http://192.168.1.10:3022/v1")
//	fmt.Println(err)
//	fmt.Println(gotSpent)
//}

func Test_ipfsRequest1(t *testing.T) {
	type args struct {
		api       string
		uploadAPI string
	}
	tests := []struct {
		name      string
		args      args
		wantSpent int64
		wantErr   bool
	}{
		{name: "#0", args: args{api: "http://ipfs.learnblockchain.cn", uploadAPI: "http://47.242.152.213:3022/v1"}},
		{name: "#1", args: args{api: "http://ipfs.learnblockchain.cn", uploadAPI: "http://47.242.152.213:3022/v1"}},
		{name: "#2", args: args{api: "http://ipfs.learnblockchain.cn", uploadAPI: "http://47.242.152.213:3022/v1"}},
		{name: "#3", args: args{api: "http://ipfs.learnblockchain.cn", uploadAPI: "http://47.242.152.213:3022/v1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSpent, err := ipfsRequest(tt.args.api, tt.args.uploadAPI)
			if (err != nil) != tt.wantErr {
				t.Errorf("ipfsRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSpent != tt.wantSpent {
				t.Errorf("ipfsRequest() gotSpent = %v, want %v", gotSpent, tt.wantSpent)
			}
		})
	}
}
