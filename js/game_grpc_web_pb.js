/**
 * @fileoverview gRPC-Web generated client stub for gameproto.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.gameproto = {};
proto.gameproto.v1 = require('./game_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.gameproto.v1.SessionClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.gameproto.v1.SessionPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.gameproto.v1.SessionRequest,
 *   !proto.gameproto.v1.SessionGrant>}
 */
const methodDescriptor_Session_NewSession = new grpc.web.MethodDescriptor(
  '/gameproto.v1.Session/NewSession',
  grpc.web.MethodType.UNARY,
  proto.gameproto.v1.SessionRequest,
  proto.gameproto.v1.SessionGrant,
  /** @param {!proto.gameproto.v1.SessionRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.gameproto.v1.SessionGrant.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.gameproto.v1.SessionRequest,
 *   !proto.gameproto.v1.SessionGrant>}
 */
const methodInfo_Session_NewSession = new grpc.web.AbstractClientBase.MethodInfo(
  proto.gameproto.v1.SessionGrant,
  /** @param {!proto.gameproto.v1.SessionRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.gameproto.v1.SessionGrant.deserializeBinary
);


/**
 * @param {!proto.gameproto.v1.SessionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.gameproto.v1.SessionGrant)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.gameproto.v1.SessionGrant>|undefined}
 *     The XHR Node Readable Stream
 */
proto.gameproto.v1.SessionClient.prototype.newSession =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/gameproto.v1.Session/NewSession',
      request,
      metadata || {},
      methodDescriptor_Session_NewSession,
      callback);
};


/**
 * @param {!proto.gameproto.v1.SessionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.gameproto.v1.SessionGrant>}
 *     A native promise that resolves to the response
 */
proto.gameproto.v1.SessionPromiseClient.prototype.newSession =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/gameproto.v1.Session/NewSession',
      request,
      metadata || {},
      methodDescriptor_Session_NewSession);
};


module.exports = proto.gameproto.v1;

