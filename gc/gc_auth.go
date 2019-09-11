package gc

const authUIDKey = "_AUTH_UID_"

func (g *gContent) SetAuthID(id int64) {
	g.Set(authUIDKey, id)
}

func (g *gContent) AuthID() int64 {
	v, exists := g.Get(authUIDKey)
	if exists {
		return v.(int64)
	}

	return 0
}
