package gapp

type ContainerServer interface {
	GetServerName() string
	ServerLoad(bootstrap *Container)
	ServerBoot(bootstrap *Container)
	ServerDestroy(bootstrap *Container)
}