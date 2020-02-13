package yrpcmsg

import "google.golang.org/grpc/metadata"

//the api ver in ymsg meta
func (this *Ymsg) MetaVer() string {
	for i := 0; i < len(this.Meta)-1; i += 2 {
		if this.Meta[i] == "v" {
			return this.Meta[i+1]
		}
	}

	return ""
}

//copy meta data from this:ymsg to md:metadata
func (this *Ymsg) CopyMeta(md *metadata.MD) {
	for i := 0; i < len(this.Meta)-1; i += 2 {
		md.Set(this.Meta[i], this.Meta[i+1])
	}
}
