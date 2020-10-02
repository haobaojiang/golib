package gtcpforward

type Decryption interface {
	Decrypt([]byte)([]byte,error)
}

type Encryption interface {
	Encrypt([]byte)([]byte,error)
}

type Server struct {
	addr        string
	forwardAddr string
	enc Encryption
	dec Decryption
}

//TODO: implement pkgOption
func New(addr string, forwardAddr string,enc Encryption,dec Decryption) *Server {
	return &Server{
		addr:        addr,
		forwardAddr: forwardAddr,
		enc: enc,
		dec: dec,
	}
}