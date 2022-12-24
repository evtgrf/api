"use strict";
// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file domainconfig/v1/archetypes.proto (package domainconfig.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck
Object.defineProperty(exports, "__esModule", { value: true });
exports.ArchetypeDefinition = exports.QueryTerm_LeafConnection = exports.QueryTerm_EdgeConnection = exports.QueryTerm_Connection = exports.QueryTerm = void 0;
const protobuf_1 = require("@bufbuild/protobuf");
const shared_pb_js_1 = require("./shared_pb.js");
/**
 * @generated from message domainconfig.v1.QueryTerm
 */
class QueryTerm extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: int64 component_id = 1;
         */
        this.componentId = protobuf_1.protoInt64.zero;
        /**
         * @generated from oneof domainconfig.v1.QueryTerm.connection
         */
        this.connection = { case: undefined };
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new QueryTerm().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new QueryTerm().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new QueryTerm().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(QueryTerm, a, b);
    }
}
exports.QueryTerm = QueryTerm;
QueryTerm.runtime = protobuf_1.proto3;
QueryTerm.typeName = "domainconfig.v1.QueryTerm";
QueryTerm.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "component_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "edge", kind: "message", T: QueryTerm_EdgeConnection, oneof: "connection" },
    { no: 3, name: "leaf", kind: "message", T: QueryTerm_LeafConnection, oneof: "connection" },
]);
/**
 * @generated from message domainconfig.v1.QueryTerm.Connection
 */
class QueryTerm_Connection extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from oneof domainconfig.v1.QueryTerm.Connection.mode
         */
        this.mode = { case: undefined };
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new QueryTerm_Connection().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new QueryTerm_Connection().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new QueryTerm_Connection().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(QueryTerm_Connection, a, b);
    }
}
exports.QueryTerm_Connection = QueryTerm_Connection;
QueryTerm_Connection.runtime = protobuf_1.proto3;
QueryTerm_Connection.typeName = "domainconfig.v1.QueryTerm.Connection";
QueryTerm_Connection.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "instance", kind: "scalar", T: 3 /* ScalarType.INT64 */, oneof: "mode" },
    { no: 2, name: "init_variable", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "mode" },
    { no: 3, name: "variable", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "mode" },
]);
/**
 * @generated from message domainconfig.v1.QueryTerm.EdgeConnection
 */
class QueryTerm_EdgeConnection extends protobuf_1.Message {
    constructor(data) {
        super();
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new QueryTerm_EdgeConnection().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new QueryTerm_EdgeConnection().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new QueryTerm_EdgeConnection().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(QueryTerm_EdgeConnection, a, b);
    }
}
exports.QueryTerm_EdgeConnection = QueryTerm_EdgeConnection;
QueryTerm_EdgeConnection.runtime = protobuf_1.proto3;
QueryTerm_EdgeConnection.typeName = "domainconfig.v1.QueryTerm.EdgeConnection";
QueryTerm_EdgeConnection.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "from_id", kind: "message", T: QueryTerm_Connection },
    { no: 2, name: "to_id", kind: "message", T: QueryTerm_Connection },
]);
/**
 * @generated from message domainconfig.v1.QueryTerm.LeafConnection
 */
class QueryTerm_LeafConnection extends protobuf_1.Message {
    constructor(data) {
        super();
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new QueryTerm_LeafConnection().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new QueryTerm_LeafConnection().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new QueryTerm_LeafConnection().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(QueryTerm_LeafConnection, a, b);
    }
}
exports.QueryTerm_LeafConnection = QueryTerm_LeafConnection;
QueryTerm_LeafConnection.runtime = protobuf_1.proto3;
QueryTerm_LeafConnection.typeName = "domainconfig.v1.QueryTerm.LeafConnection";
QueryTerm_LeafConnection.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: QueryTerm_Connection },
]);
/**
 * @generated from message domainconfig.v1.ArchetypeDefinition
 */
class ArchetypeDefinition extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: repeated domainconfig.v1.QueryTerm query_terms = 2;
         */
        this.queryTerms = [];
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new ArchetypeDefinition().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new ArchetypeDefinition().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new ArchetypeDefinition().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(ArchetypeDefinition, a, b);
    }
}
exports.ArchetypeDefinition = ArchetypeDefinition;
ArchetypeDefinition.runtime = protobuf_1.proto3;
ArchetypeDefinition.typeName = "domainconfig.v1.ArchetypeDefinition";
ArchetypeDefinition.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "metadata", kind: "message", T: shared_pb_js_1.Metadata },
    { no: 2, name: "query_terms", kind: "message", T: QueryTerm, repeated: true },
]);
