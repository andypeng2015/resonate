package callback

import (
	"encoding/json"
	"fmt"

	"github.com/resonatehq/resonate/pkg/message"
)

type Callback struct {
	Id            string          `json:"id"`
	PromiseId     string          `json:"promiseId"`
	RootPromiseId string          `json:"rootPromiseId"`
	Recv          json.RawMessage `json:"-"`
	Mesg          *message.Mesg   `json:"-"`
	Timeout       int64           `json:"timeout"`
	CreatedOn     int64           `json:"createdOn"`
}

func (c *Callback) String() string {
	return fmt.Sprintf(
		"Callback(id=%s, promiseId=%s, recv=%s, mesg=%v, timeout=%d, createdOn=%d)",
		c.Id,
		c.PromiseId,
		c.Recv,
		c.Mesg,
		c.Timeout,
		c.CreatedOn,
	)
}

func (c1 *Callback) Equals(c2 *Callback) bool {
	// for dst only
	return c1.Id == c2.Id && c1.PromiseId == c2.PromiseId
}
