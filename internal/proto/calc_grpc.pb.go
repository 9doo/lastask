package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)


const _ = grpc.SupportPackageIsVersion9

const (
	Calculator_GetTask_FullMethodName      = "/calc_service.Calculator/GetTask"
	Calculator_SubmitResult_FullMethodName = "/calc_service.Calculator/SubmitResult"
)


type CalculatorClient interface {
	GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	SubmitResult(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, Calculator_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) SubmitResult(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, Calculator_SubmitResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type CalculatorServer interface {
	GetTask(context.Context, *TaskRequest) (*TaskResponse, error)
	SubmitResult(context.Context, *ResultRequest) (*ResultResponse, error)
	mustEmbedUnimplementedCalculatorServer()
}


type UnimplementedCalculatorServer struct{}

func (UnimplementedCalculatorServer) GetTask(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedCalculatorServer) SubmitResult(context.Context, *ResultRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitResult not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}
func (UnimplementedCalculatorServer) testEmbeddedByValue()                    {}


type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).GetTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_SubmitResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).SubmitResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_SubmitResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).SubmitResult(ctx, req.(*ResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}


var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calc_service.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _Calculator_GetTask_Handler,
		},
		{
			MethodName: "SubmitResult",
			Handler:    _Calculator_SubmitResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/calc.proto",
}
