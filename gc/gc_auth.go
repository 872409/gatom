package gc

const authUIDKey = "_AUTH_UID_"

func (g *GContext) SetAuthID(id int64) {
	g.Set(authUIDKey, id)
}
func (g *GContext) GetGRequestId(gRequestIdKey ...string) string {
	key := "g-request-id"
	if len(gRequestIdKey) == 1 {
		key = gRequestIdKey[0]
	}
	return g.GetHeader(key)
}

func (g *GContext) AuthID() int64 {
	v, exists := g.Get(authUIDKey)
	if exists {
		return v.(int64)
	}

	return 0
}
