package gapp

type Server interface {
	GetServerName() string
	ServerInit(bootstrap *Container)
	ServerBoot(bootstrap *Container)
	ServerDestroy(bootstrap *Container)
}