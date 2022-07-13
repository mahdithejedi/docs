import route_pb2_grpc
import route_pb2


def _get_msgs():
    from datetime import datetime
    maps = [
        {'a': '1',
         'b': '2',
         'time': f'{datetime.now()}'} for _ in range(20)
    ]
    for _map in maps:
        yield route_pb2.pure_data(
            data=_map
        )


if __name__ == '__main__':
    import grpc
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = route_pb2_grpc.RouteGuideStub(channel)
        tester = stub.GetTimings(
            _get_msgs()
        )
        for test in tester:
            print('TESTTTT', test)
