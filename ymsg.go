package yrpcmsg

import "google.golang.org/grpc/metadata"

//the api ver in ymsg meta
func (this *Ymsg) MetaVer() string {
	if this.MetaInfo == nil {
		return ""
	}

	for _, item := range this.MetaInfo.Val {
		if item.Key == "v" {
			if len(item.Vals) > 0 {
				return item.Vals[0]
			} else {
				return ""
			}
		}
	}
	return ""
}

//copy meta data from this:ymsg to md:metadata
func (this *Ymsg) CopyMeta(md *metadata.MD) {
	if this.MetaInfo == nil {
		return
	}

	for _, metaItem := range this.MetaInfo.Val {
		md.Set(metaItem.Key, metaItem.Vals...)
	}

}

//ymsg metainfo to metadata MD
func (this *Ymsg) GrpcMeta() metadata.MD {
	md := metadata.MD{}

	if this.MetaInfo == nil {
		return md
	}

	for _, metaItem := range this.MetaInfo.Val {
		md.Set(metaItem.Key, metaItem.Vals...)
	}

	return md
}

//metadata.MD to yrpcmsg.meta
func GrpcMD2YrpcMeta(md metadata.MD) Meta {
	m := Meta{}
	for key, vals := range md {
		if key == "content-type" {
			//no need this item
			continue
		}
		mi := &MetaItem{
			Key:  key,
			Vals: vals,
		}
		m.Val = append(m.Val, mi)
	}

	return m
}

//yrpcmsg.meta to metadata.MD
func YrpcMeta2GrpcMD(meta *Meta) metadata.MD {
	md := metadata.MD{}

	for _, metaItem := range meta.Val {
		md.Set(metaItem.Key, metaItem.Vals...)
	}

	return md
}
