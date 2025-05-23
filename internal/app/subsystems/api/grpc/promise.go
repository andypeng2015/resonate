package grpc

import (
	"context"
	"encoding/json"

	"github.com/resonatehq/resonate/internal/kernel/t_api"
	"github.com/resonatehq/resonate/internal/util"

	"github.com/resonatehq/resonate/internal/app/subsystems/api/grpc/pb"
	"github.com/resonatehq/resonate/pkg/callback"
	"github.com/resonatehq/resonate/pkg/idempotency"
	"github.com/resonatehq/resonate/pkg/message"
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/resonatehq/resonate/pkg/receiver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) ReadPromise(c context.Context, r *pb.ReadPromiseRequest) (*pb.ReadPromiseResponse, error) {
	res, err := s.api.Process(r.RequestId, &t_api.Request{
		Payload: &t_api.ReadPromiseRequest{
			Id: r.Id,
		},
	})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.ReadPromise != nil, "result must not be nil")
	return &pb.ReadPromiseResponse{
		Promise: protoPromise(res.ReadPromise.Promise),
	}, nil
}

func (s *server) SearchPromises(c context.Context, r *pb.SearchPromisesRequest) (*pb.SearchPromisesResponse, error) {
	req, err := s.api.SearchPromises(r.Id, searchState(r.State), r.Tags, int(r.Limit), r.Cursor)
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	res, err := s.api.Process(r.RequestId, &t_api.Request{Payload: req})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.SearchPromises != nil, "result must not be nil")

	promises := make([]*pb.Promise, len(res.SearchPromises.Promises))
	for i, promise := range res.SearchPromises.Promises {
		promises[i] = protoPromise(promise)
	}

	var cursor string
	if res.SearchPromises.Cursor != nil {
		var err error
		cursor, err = res.SearchPromises.Cursor.Encode()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.SearchPromisesResponse{
		Promises: promises,
		Cursor:   cursor,
	}, nil
}

func (s *server) CreatePromise(c context.Context, r *pb.CreatePromiseRequest) (*pb.CreatePromiseResponse, error) {
	var idempotencyKey *idempotency.Key
	if r.IdempotencyKey != "" {
		idempotencyKey = util.ToPointer(idempotency.Key(r.IdempotencyKey))
	}

	var headers map[string]string
	if r.Param != nil {
		headers = r.Param.Headers
	}

	var data []byte
	if r.Param != nil {
		data = r.Param.Data
	}

	res, err := s.api.Process(r.RequestId, &t_api.Request{
		Payload: &t_api.CreatePromiseRequest{
			Id:             r.Id,
			IdempotencyKey: idempotencyKey,
			Strict:         r.Strict,
			Param:          promise.Value{Headers: headers, Data: data},
			Timeout:        r.Timeout,
			Tags:           r.Tags,
		},
	})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.CreatePromise != nil, "result must not be nil")
	return &pb.CreatePromiseResponse{
		Noop:    res.CreatePromise.Status == t_api.StatusOK,
		Promise: protoPromise(res.CreatePromise.Promise),
	}, nil
}

func (s *server) CreatePromiseAndTask(c context.Context, r *pb.CreatePromiseAndTaskRequest) (*pb.CreatePromiseAndTaskResponse, error) {
	if r.Promise == nil {
		return nil, status.Error(codes.InvalidArgument, "The field task is required.")
	}

	if r.Task == nil {
		return nil, status.Error(codes.InvalidArgument, "The field task is required.")
	}
	if r.Task.Ttl < 0 {
		return nil, status.Error(codes.InvalidArgument, "The field ttl must be greater than or equal to zero")
	}

	var idempotencyKey *idempotency.Key
	if r.Promise.IdempotencyKey != "" {
		idempotencyKey = util.ToPointer(idempotency.Key(r.Promise.IdempotencyKey))
	}

	var headers map[string]string
	if r.Promise.Param != nil {
		headers = r.Promise.Param.Headers
	}

	var data []byte
	if r.Promise.Param != nil {
		data = r.Promise.Param.Data
	}

	res, err := s.api.Process(r.Promise.RequestId, &t_api.Request{
		Payload: &t_api.CreatePromiseAndTaskRequest{
			Promise: &t_api.CreatePromiseRequest{
				Id:             r.Promise.Id,
				IdempotencyKey: idempotencyKey,
				Strict:         r.Promise.Strict,
				Param:          promise.Value{Headers: headers, Data: data},
				Timeout:        r.Promise.Timeout,
				Tags:           r.Promise.Tags,
			},
			Task: &t_api.CreateTaskRequest{
				PromiseId: r.Promise.Id,
				ProcessId: r.Task.ProcessId,
				Ttl:       int(r.Task.Ttl),
				Timeout:   r.Promise.Timeout,
			},
		},
	})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.CreatePromiseAndTask != nil, "result must not be nil")
	return &pb.CreatePromiseAndTaskResponse{
		Noop:    res.CreatePromiseAndTask.Status == t_api.StatusOK,
		Promise: protoPromise(res.CreatePromiseAndTask.Promise),
	}, nil
}

func (s *server) ResolvePromise(c context.Context, r *pb.ResolvePromiseRequest) (*pb.ResolvePromiseResponse, error) {
	var idempotencyKey *idempotency.Key
	if r.IdempotencyKey != "" {
		idempotencyKey = util.ToPointer(idempotency.Key(r.IdempotencyKey))
	}

	var headers map[string]string
	if r.Value != nil {
		headers = r.Value.Headers
	}

	var data []byte
	if r.Value != nil {
		data = r.Value.Data
	}

	res, err := s.api.Process(r.RequestId, &t_api.Request{
		Payload: &t_api.CompletePromiseRequest{
			Id:             r.Id,
			IdempotencyKey: idempotencyKey,
			Strict:         r.Strict,
			State:          promise.Resolved,
			Value:          promise.Value{Headers: headers, Data: data},
		},
	})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.CompletePromise != nil, "result must not be nil")
	return &pb.ResolvePromiseResponse{
		Noop:    res.CompletePromise.Status == t_api.StatusOK,
		Promise: protoPromise(res.CompletePromise.Promise),
	}, nil
}

func (s *server) RejectPromise(c context.Context, r *pb.RejectPromiseRequest) (*pb.RejectPromiseResponse, error) {
	var idempotencyKey *idempotency.Key
	if r.IdempotencyKey != "" {
		idempotencyKey = util.ToPointer(idempotency.Key(r.IdempotencyKey))
	}

	var headers map[string]string
	if r.Value != nil {
		headers = r.Value.Headers
	}

	var data []byte
	if r.Value != nil {
		data = r.Value.Data
	}

	res, err := s.api.Process(r.RequestId, &t_api.Request{
		Payload: &t_api.CompletePromiseRequest{
			Id:             r.Id,
			IdempotencyKey: idempotencyKey,
			Strict:         r.Strict,
			State:          promise.Rejected,
			Value:          promise.Value{Headers: headers, Data: data},
		},
	})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.CompletePromise != nil, "result must not be nil")
	return &pb.RejectPromiseResponse{
		Noop:    res.CompletePromise.Status == t_api.StatusOK,
		Promise: protoPromise(res.CompletePromise.Promise),
	}, nil
}

func (s *server) CancelPromise(c context.Context, r *pb.CancelPromiseRequest) (*pb.CancelPromiseResponse, error) {
	var idempotencyKey *idempotency.Key
	if r.IdempotencyKey != "" {
		idempotencyKey = util.ToPointer(idempotency.Key(r.IdempotencyKey))
	}

	var headers map[string]string
	if r.Value != nil {
		headers = r.Value.Headers
	}

	var data []byte
	if r.Value != nil {
		data = r.Value.Data
	}

	res, err := s.api.Process(r.RequestId, &t_api.Request{
		Payload: &t_api.CompletePromiseRequest{
			Id:             r.Id,
			IdempotencyKey: idempotencyKey,
			Strict:         r.Strict,
			State:          promise.Canceled,
			Value:          promise.Value{Headers: headers, Data: data},
		},
	})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.CompletePromise != nil, "result must not be nil")
	return &pb.CancelPromiseResponse{
		Noop:    res.CompletePromise.Status == t_api.StatusOK,
		Promise: protoPromise(res.CompletePromise.Promise),
	}, nil
}

func (s *server) CreateCallback(c context.Context, r *pb.CreateCallbackRequest) (*pb.CreateCallbackResponse, error) {
	recv, rErr := protoRecv(r.Recv)
	if rErr != nil {
		return nil, rErr
	}

	res, err := s.api.Process(r.RequestId, &t_api.Request{
		Payload: &t_api.CreateCallbackRequest{
			Id:        util.ResumeId(r.RootPromiseId, r.PromiseId),
			PromiseId: r.PromiseId,
			Recv:      recv,
			Mesg:      &message.Mesg{Type: "resume", Root: r.RootPromiseId, Leaf: r.PromiseId},
			Timeout:   r.Timeout,
		},
	})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.CreateCallback != nil, "result must not be nil")
	return &pb.CreateCallbackResponse{
		Noop:     res.CreateCallback.Status == t_api.StatusOK,
		Callback: protoCallback(res.CreateCallback.Callback),
		Promise:  protoPromise(res.CreateCallback.Promise),
	}, nil
}

func (s *server) CreateSubscription(c context.Context, r *pb.CreateSubscriptionRequest) (*pb.CreateSubscriptionResponse, error) {
	recv, rErr := protoRecv(r.Recv)
	if rErr != nil {
		return nil, rErr
	}

	res, err := s.api.Process(r.RequestId, &t_api.Request{
		Payload: &t_api.CreateCallbackRequest{
			Id:        util.NotifyId(r.PromiseId, r.Id),
			PromiseId: r.PromiseId,
			Recv:      recv,
			Mesg:      &message.Mesg{Type: "notify", Root: r.PromiseId},
			Timeout:   r.Timeout,
		},
	})
	if err != nil {
		return nil, status.Error(s.code(err.Code), err.Error())
	}

	util.Assert(res.CreateCallback != nil, "result must not be nil")
	return &pb.CreateSubscriptionResponse{
		Noop:     res.CreateCallback.Status == t_api.StatusOK,
		Callback: protoCallback(res.CreateCallback.Callback),
		Promise:  protoPromise(res.CreateCallback.Promise),
	}, nil
}

// Helper functions

func protoPromise(promise *promise.Promise) *pb.Promise {
	if promise == nil {
		return nil
	}

	return &pb.Promise{
		Id:                        promise.Id,
		State:                     protoState(promise.State),
		Param:                     &pb.Value{Headers: promise.Param.Headers, Data: promise.Param.Data},
		Value:                     &pb.Value{Headers: promise.Value.Headers, Data: promise.Value.Data},
		Timeout:                   promise.Timeout,
		IdempotencyKeyForCreate:   string(util.SafeDeref(promise.IdempotencyKeyForCreate)),
		IdempotencyKeyForComplete: string(util.SafeDeref(promise.IdempotencyKeyForComplete)),
		CreatedOn:                 util.SafeDeref(promise.CreatedOn),
		CompletedOn:               util.SafeDeref(promise.CompletedOn),
		Tags:                      promise.Tags,
	}
}

func protoState(state promise.State) pb.State {
	switch state {
	case promise.Pending:
		return pb.State_PENDING
	case promise.Resolved:
		return pb.State_RESOLVED
	case promise.Rejected:
		return pb.State_REJECTED
	case promise.Timedout:
		return pb.State_REJECTED_TIMEDOUT
	case promise.Canceled:
		return pb.State_REJECTED_CANCELED
	default:
		panic("invalid state")
	}
}

func searchState(searchState pb.SearchState) string {
	switch searchState {
	case pb.SearchState_SEARCH_ALL:
		return ""
	case pb.SearchState_SEARCH_PENDING:
		return "pending"
	case pb.SearchState_SEARCH_RESOLVED:
		return "resolved"
	case pb.SearchState_SEARCH_REJECTED:
		return "rejected"
	default:
		return "INVALID" // will be validated in the helper function
	}
}

func protoCallback(callback *callback.Callback) *pb.Callback {
	if callback == nil {
		return nil
	}

	return &pb.Callback{
		Id:        callback.Id,
		PromiseId: callback.PromiseId,
		Timeout:   callback.Timeout,
		CreatedOn: callback.CreatedOn,
	}
}

func protoRecv(recv *pb.Recv) ([]byte, error) {
	switch r := recv.Recv.(type) {
	case *pb.Recv_Logical:
		return json.Marshal(&r.Logical)
	case *pb.Recv_Physical:
		return json.Marshal(&receiver.Recv{Type: r.Physical.Type, Data: r.Physical.Data})
	default:
		return nil, status.Error(codes.InvalidArgument, "The field recv is required.")
	}
}
