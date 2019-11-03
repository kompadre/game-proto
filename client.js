const { EchoRequest, EchoResponse } = require('./js/game_pb');
const { SessionClient, SessionPromiseClient } = require('./js/game_grpc_web_pb.js');

var echoService = new EchoServiceClient('http://localhost:8080');

var request = new EchoRequest();
request.setMessage('Hello World!');

echoService.echo(request, {}, function(err, response) {
    // ...
});