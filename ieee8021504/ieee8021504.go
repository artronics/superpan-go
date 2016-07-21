package ieee8021504

type IEEE8021504 struct {
	MLME mlme
}

type Requester interface {
	request(mlmePrimitive interface{}, mlme *mlme) error
}

type mlme struct {
}

func (m *mlme) Request(req Requester) error {
	return req.request(req, m)
}
