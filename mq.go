package main

type MsgQue struct {
	count int32         // number of queues of this message queue
	bank  []interface{} // the place that saves the data
}

func (q *MsgQue) AppendOne(v interface{}) error {
	q.bank = append(q.bank, v)
	return nil
}

// Peek Show the latest queue element.
//   Return error when the element is not found.
//   Return the element with nil error when element is found and retrived successfully.
func (q *MsgQue) Peek() (interface{}, error) {

	return nil, nil
}

// Len Show the number of element in this MsgQue object.
func (q *MsgQue) Len() int {
	return len(q.bank)
}

// ConsumeOne Get and remove the first element from the queue.
//   Return the element with nil error when element not.
//   Return nil with error when element is not found or Error occurred.
func (q *MsgQue) ConsumeOne() (interface{}, error) {
	e := q.bank[0]
	q.bank = q.bank[1:]
	return e, nil
}
