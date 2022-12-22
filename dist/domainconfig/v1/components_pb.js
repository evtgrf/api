"use strict";
// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file domainconfig/v1/components.proto (package domainconfig.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck
Object.defineProperty(exports, "__esModule", { value: true });
exports.ComponentDefinition_Index = exports.ComponentDefinition_IndexEntry = exports.ComponentDefinition = exports.FieldDefinition = void 0;
const protobuf_1 = require("@bufbuild/protobuf");
const shared_pb_js_1 = require("./shared_pb.js");
const properties_pb_js_1 = require("./properties_pb.js");
/**
 * @generated from message domainconfig.v1.FieldDefinition
 */
class FieldDefinition extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: int64 property_id = 2;
         */
        this.propertyId = protobuf_1.protoInt64.zero;
        /**
         * @generated from field: repeated domainconfig.v1.DataType property_overrides = 3;
         */
        this.propertyOverrides = [];
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new FieldDefinition().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new FieldDefinition().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new FieldDefinition().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(FieldDefinition, a, b);
    }
}
exports.FieldDefinition = FieldDefinition;
FieldDefinition.runtime = protobuf_1.proto3;
FieldDefinition.typeName = "domainconfig.v1.FieldDefinition";
FieldDefinition.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "metadata", kind: "message", T: shared_pb_js_1.Metadata },
    { no: 2, name: "property_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "property_overrides", kind: "message", T: properties_pb_js_1.DataType, repeated: true },
]);
/**
 * @generated from message domainconfig.v1.ComponentDefinition
 */
class ComponentDefinition extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: bool is_edge = 2;
         */
        this.isEdge = false;
        /**
         * @generated from field: repeated domainconfig.v1.FieldDefinition fields = 3;
         */
        this.fields = [];
        /**
         * @generated from field: repeated domainconfig.v1.ComponentDefinition.Index indices = 4;
         */
        this.indices = [];
        /**
         * @generated from field: bool disable_history = 5;
         */
        this.disableHistory = false;
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new ComponentDefinition().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new ComponentDefinition().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new ComponentDefinition().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(ComponentDefinition, a, b);
    }
}
exports.ComponentDefinition = ComponentDefinition;
ComponentDefinition.runtime = protobuf_1.proto3;
ComponentDefinition.typeName = "domainconfig.v1.ComponentDefinition";
ComponentDefinition.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "metadata", kind: "message", T: shared_pb_js_1.Metadata },
    { no: 2, name: "is_edge", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "fields", kind: "message", T: FieldDefinition, repeated: true },
    { no: 4, name: "indices", kind: "message", T: ComponentDefinition_Index, repeated: true },
    { no: 5, name: "disable_history", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
]);
/**
 * @generated from message domainconfig.v1.ComponentDefinition.IndexEntry
 */
class ComponentDefinition_IndexEntry extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: int64 value_index = 1;
         */
        this.valueIndex = protobuf_1.protoInt64.zero;
        /**
         * @generated from field: bool is_descending = 2;
         */
        this.isDescending = false;
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new ComponentDefinition_IndexEntry().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new ComponentDefinition_IndexEntry().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new ComponentDefinition_IndexEntry().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(ComponentDefinition_IndexEntry, a, b);
    }
}
exports.ComponentDefinition_IndexEntry = ComponentDefinition_IndexEntry;
ComponentDefinition_IndexEntry.runtime = protobuf_1.proto3;
ComponentDefinition_IndexEntry.typeName = "domainconfig.v1.ComponentDefinition.IndexEntry";
ComponentDefinition_IndexEntry.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "value_index", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "is_descending", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
]);
/**
 * @generated from message domainconfig.v1.ComponentDefinition.Index
 */
class ComponentDefinition_Index extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: repeated domainconfig.v1.ComponentDefinition.IndexEntry entries = 1;
         */
        this.entries = [];
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new ComponentDefinition_Index().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new ComponentDefinition_Index().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new ComponentDefinition_Index().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(ComponentDefinition_Index, a, b);
    }
}
exports.ComponentDefinition_Index = ComponentDefinition_Index;
ComponentDefinition_Index.runtime = protobuf_1.proto3;
ComponentDefinition_Index.typeName = "domainconfig.v1.ComponentDefinition.Index";
ComponentDefinition_Index.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "entries", kind: "message", T: ComponentDefinition_IndexEntry, repeated: true },
]);
