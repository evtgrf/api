"use strict";
// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file domainconfig/v1/shared.proto (package domainconfig.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck
Object.defineProperty(exports, "__esModule", { value: true });
exports.Metadata = void 0;
const protobuf_1 = require("@bufbuild/protobuf");
/**
 * @generated from message domainconfig.v1.Metadata
 */
class Metadata extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: int64 id = 1;
         */
        this.id = protobuf_1.protoInt64.zero;
        /**
         * @generated from field: string label = 2;
         */
        this.label = "";
        /**
         * @generated from field: string description = 3;
         */
        this.description = "";
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new Metadata().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new Metadata().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new Metadata().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(Metadata, a, b);
    }
}
exports.Metadata = Metadata;
Metadata.runtime = protobuf_1.proto3;
Metadata.typeName = "domainconfig.v1.Metadata";
Metadata.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "label", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
]);
