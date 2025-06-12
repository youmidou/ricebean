package sys_net

// interface
const (
	state_read_init = 1
	state_read_head = 2
	state_read_body = 3
	P_HEAD          = 4
	p_buffer_len    = 4 * 1024
)
const (
	ConnectIdle       = 0 //闲置
	ConnectStart      = 1 //开始连接
	ConnectSuccess    = 2 //连接成功
	ConnectFail       = 3 //连接失败
	ConnectDisconnect = 4 //断开
)
