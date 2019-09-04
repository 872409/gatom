package gc

const tokenUIDKey = "UID"

func (g *gContent) SetUID(uid int64) {
	g.Set(tokenUIDKey, uid)
}

func (g *gContent) GetUID() (uid int64, exists bool) {
	v, exists := g.Get(tokenUIDKey)
	if exists {
		uid = v.(int64)
		return
	}

	return
}
