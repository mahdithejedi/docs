from route_pb2_grpc import RouteGuideServicer as _RouteGuideService, add_RouteGuideServicer_to_server
from route_pb2 import structured_data
from random import uniform


class RouteGuideServicer(_RouteGuideService):
    def __init__(self):
        super().__init__()

    def GetTimings(self, request_iterator, context):
        for iteration in request_iterator:
            print('fucking iteration', iteration)
            yield structured_data(
                service_name='Tester',
                uptime_duration=uniform(1.0, 10000.00),
                request_duration=uniform(1.0, 10000.00)
            )


if __name__ == '__main__':
    def _start_server(server, port='[::]:50051'):
        server.add_insecure_port(port)
        print('server started!!!!!')
        server.start()
        server.wait_for_termination()

    import grpc
    from concurrent import futures
    _server = grpc.server(futures.ThreadPoolExecutor(max_workers=1))
    add_RouteGuideServicer_to_server(
        RouteGuideServicer(), _server
    )
    _start_server(_server)



