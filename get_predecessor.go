package chord

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	pb "github.com/wang502/chord/protobuf"
)

// GetPredecessorResponse represents a response to request of getting predecessor
type GetPredecessorResponse struct {
	ID   string
	host string
}

// NewGetPredecessorResponse initializes a GetPredecessorResponse object
func NewGetPredecessorResponse(id []byte, host string) *GetPredecessorResponse {
	return &GetPredecessorResponse{
		ID:   string(id),
		host: host,
	}
}

// Invalid check whether the GetPredecessorResponse is invalid (containing empty values)
func (resp *GetPredecessorResponse) Invalid() bool {
	return resp.ID == "" || resp.host == ""
}

// Encode encodes GetPredecessorResponse into data buffer
func (resp *GetPredecessorResponse) Encode(w io.Writer) (int, error) {
	pb := &pb.GetPredecessorResponse{
		ID:   resp.ID,
		Host: resp.host,
	}
	data, err := proto.Marshal(pb)
	if err != nil {
		//log.Println("chord.GetPredecessorResponse.encode.error")
		return -1, fmt.Errorf("encode GetPredecessorResponse failed: %s", err)
	}

	return w.Write(data)
}

// Decode decodes data from buffer and stores it in GetPredecessorResponse
func (resp *GetPredecessorResponse) Decode(r io.Reader) (int, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		//log.Println("chord.GetPredecessorResponse.decode.error")
		return -1, fmt.Errorf("decode GetPredecessorResponse failed: %s", err)
	}

	pb := &pb.GetPredecessorResponse{}
	if err = proto.Unmarshal(data, pb); err != nil {
		return -1, err
	}

	resp.ID = pb.ID
	resp.host = pb.Host
	return len(data), nil
}
