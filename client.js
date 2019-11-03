const { LookAroundAnswer, LookAroundRequest, Object, ObjectAttribute, SessionGrant, SessionRequest } = require('./js/game_pb');
const { SessionClient, SessionPromiseClient, } = require('./js/game_grpc_web_pb.js');

var echoService = new SessionClient('http://localhost:8080');

var request = new SessionRequest();
request.SetUsername("kompadre");
request.SetPassword("Unlikely");


echoService.echo(request, {}, function(err, response) {
    // ...
});