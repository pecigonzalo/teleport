/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter eslint_disable,add_pb_suffix,server_grpc1,ts_nocheck
// @generated from protobuf file "teleport/userpreferences/v1/userpreferences.proto" (package "teleport.userpreferences.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import { Empty } from "../../../google/protobuf/empty_pb";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { DiscoverResourcePreferences } from "./discover_resource_preferences_pb";
import { SideNavDrawerMode } from "./sidenav_preferences_pb";
import { AccessGraphUserPreferences } from "./access_graph_pb";
import { UnifiedResourcePreferences } from "./unified_resource_preferences_pb";
import { ClusterUserPreferences } from "./cluster_preferences_pb";
import { OnboardUserPreferences } from "./onboard_pb";
import { Theme } from "./theme_pb";
/**
 * UserPreferences is a collection of different user changeable preferences for the frontend.
 *
 * @generated from protobuf message teleport.userpreferences.v1.UserPreferences
 */
export interface UserPreferences {
    /**
     * theme is the theme of the frontend.
     *
     * @generated from protobuf field: teleport.userpreferences.v1.Theme theme = 2;
     */
    theme: Theme;
    /**
     * onboard is the preferences from the onboarding questionnaire.
     *
     * @generated from protobuf field: teleport.userpreferences.v1.OnboardUserPreferences onboard = 3;
     */
    onboard?: OnboardUserPreferences;
    /**
     * cluster_preferences are user preferences saved per cluster.
     *
     * @generated from protobuf field: teleport.userpreferences.v1.ClusterUserPreferences cluster_preferences = 4;
     */
    clusterPreferences?: ClusterUserPreferences;
    /**
     * unified_resource_preferences are user preferences saved for the Unified Resource web UI
     *
     * @generated from protobuf field: teleport.userpreferences.v1.UnifiedResourcePreferences unified_resource_preferences = 5;
     */
    unifiedResourcePreferences?: UnifiedResourcePreferences;
    /**
     * access_graph is the preferences for Access Graph.
     *
     * @generated from protobuf field: teleport.userpreferences.v1.AccessGraphUserPreferences access_graph = 6;
     */
    accessGraph?: AccessGraphUserPreferences;
    /**
     * side_nav_drawer_mode is the sidenav drawer behavior preference in the frontend.
     *
     * @generated from protobuf field: teleport.userpreferences.v1.SideNavDrawerMode side_nav_drawer_mode = 7;
     */
    sideNavDrawerMode: SideNavDrawerMode;
    /**
     * discover_resource_preferences are user preferences saved for the discover resource web UI.
     *
     * @generated from protobuf field: teleport.userpreferences.v1.DiscoverResourcePreferences discover_resource_preferences = 8;
     */
    discoverResourcePreferences?: DiscoverResourcePreferences;
}
/**
 * GetUserPreferencesRequest is a request to get the user preferences.
 *
 * @generated from protobuf message teleport.userpreferences.v1.GetUserPreferencesRequest
 */
export interface GetUserPreferencesRequest {
}
/**
 * GetUserPreferencesResponse is a response to get the user preferences.
 *
 * @generated from protobuf message teleport.userpreferences.v1.GetUserPreferencesResponse
 */
export interface GetUserPreferencesResponse {
    /**
     * preferences is the user preferences.
     *
     * @generated from protobuf field: teleport.userpreferences.v1.UserPreferences preferences = 1;
     */
    preferences?: UserPreferences;
}
/**
 * UpsertUserPreferencesRequest is a request to create or update the user preferences.
 *
 * @generated from protobuf message teleport.userpreferences.v1.UpsertUserPreferencesRequest
 */
export interface UpsertUserPreferencesRequest {
    /**
     * preferences is the new user preferences to set.
     *
     * @generated from protobuf field: teleport.userpreferences.v1.UserPreferences preferences = 1;
     */
    preferences?: UserPreferences;
}
// @generated message type with reflection information, may provide speed optimized methods
class UserPreferences$Type extends MessageType<UserPreferences> {
    constructor() {
        super("teleport.userpreferences.v1.UserPreferences", [
            { no: 2, name: "theme", kind: "enum", T: () => ["teleport.userpreferences.v1.Theme", Theme, "THEME_"] },
            { no: 3, name: "onboard", kind: "message", T: () => OnboardUserPreferences },
            { no: 4, name: "cluster_preferences", kind: "message", T: () => ClusterUserPreferences },
            { no: 5, name: "unified_resource_preferences", kind: "message", T: () => UnifiedResourcePreferences },
            { no: 6, name: "access_graph", kind: "message", T: () => AccessGraphUserPreferences },
            { no: 7, name: "side_nav_drawer_mode", kind: "enum", T: () => ["teleport.userpreferences.v1.SideNavDrawerMode", SideNavDrawerMode, "SIDE_NAV_DRAWER_MODE_"] },
            { no: 8, name: "discover_resource_preferences", kind: "message", T: () => DiscoverResourcePreferences }
        ]);
    }
    create(value?: PartialMessage<UserPreferences>): UserPreferences {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.theme = 0;
        message.sideNavDrawerMode = 0;
        if (value !== undefined)
            reflectionMergePartial<UserPreferences>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserPreferences): UserPreferences {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* teleport.userpreferences.v1.Theme theme */ 2:
                    message.theme = reader.int32();
                    break;
                case /* teleport.userpreferences.v1.OnboardUserPreferences onboard */ 3:
                    message.onboard = OnboardUserPreferences.internalBinaryRead(reader, reader.uint32(), options, message.onboard);
                    break;
                case /* teleport.userpreferences.v1.ClusterUserPreferences cluster_preferences */ 4:
                    message.clusterPreferences = ClusterUserPreferences.internalBinaryRead(reader, reader.uint32(), options, message.clusterPreferences);
                    break;
                case /* teleport.userpreferences.v1.UnifiedResourcePreferences unified_resource_preferences */ 5:
                    message.unifiedResourcePreferences = UnifiedResourcePreferences.internalBinaryRead(reader, reader.uint32(), options, message.unifiedResourcePreferences);
                    break;
                case /* teleport.userpreferences.v1.AccessGraphUserPreferences access_graph */ 6:
                    message.accessGraph = AccessGraphUserPreferences.internalBinaryRead(reader, reader.uint32(), options, message.accessGraph);
                    break;
                case /* teleport.userpreferences.v1.SideNavDrawerMode side_nav_drawer_mode */ 7:
                    message.sideNavDrawerMode = reader.int32();
                    break;
                case /* teleport.userpreferences.v1.DiscoverResourcePreferences discover_resource_preferences */ 8:
                    message.discoverResourcePreferences = DiscoverResourcePreferences.internalBinaryRead(reader, reader.uint32(), options, message.discoverResourcePreferences);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UserPreferences, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* teleport.userpreferences.v1.Theme theme = 2; */
        if (message.theme !== 0)
            writer.tag(2, WireType.Varint).int32(message.theme);
        /* teleport.userpreferences.v1.OnboardUserPreferences onboard = 3; */
        if (message.onboard)
            OnboardUserPreferences.internalBinaryWrite(message.onboard, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* teleport.userpreferences.v1.ClusterUserPreferences cluster_preferences = 4; */
        if (message.clusterPreferences)
            ClusterUserPreferences.internalBinaryWrite(message.clusterPreferences, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* teleport.userpreferences.v1.UnifiedResourcePreferences unified_resource_preferences = 5; */
        if (message.unifiedResourcePreferences)
            UnifiedResourcePreferences.internalBinaryWrite(message.unifiedResourcePreferences, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* teleport.userpreferences.v1.AccessGraphUserPreferences access_graph = 6; */
        if (message.accessGraph)
            AccessGraphUserPreferences.internalBinaryWrite(message.accessGraph, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* teleport.userpreferences.v1.SideNavDrawerMode side_nav_drawer_mode = 7; */
        if (message.sideNavDrawerMode !== 0)
            writer.tag(7, WireType.Varint).int32(message.sideNavDrawerMode);
        /* teleport.userpreferences.v1.DiscoverResourcePreferences discover_resource_preferences = 8; */
        if (message.discoverResourcePreferences)
            DiscoverResourcePreferences.internalBinaryWrite(message.discoverResourcePreferences, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.userpreferences.v1.UserPreferences
 */
export const UserPreferences = new UserPreferences$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetUserPreferencesRequest$Type extends MessageType<GetUserPreferencesRequest> {
    constructor() {
        super("teleport.userpreferences.v1.GetUserPreferencesRequest", []);
    }
    create(value?: PartialMessage<GetUserPreferencesRequest>): GetUserPreferencesRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetUserPreferencesRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetUserPreferencesRequest): GetUserPreferencesRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: GetUserPreferencesRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.userpreferences.v1.GetUserPreferencesRequest
 */
export const GetUserPreferencesRequest = new GetUserPreferencesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetUserPreferencesResponse$Type extends MessageType<GetUserPreferencesResponse> {
    constructor() {
        super("teleport.userpreferences.v1.GetUserPreferencesResponse", [
            { no: 1, name: "preferences", kind: "message", T: () => UserPreferences }
        ]);
    }
    create(value?: PartialMessage<GetUserPreferencesResponse>): GetUserPreferencesResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetUserPreferencesResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetUserPreferencesResponse): GetUserPreferencesResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* teleport.userpreferences.v1.UserPreferences preferences */ 1:
                    message.preferences = UserPreferences.internalBinaryRead(reader, reader.uint32(), options, message.preferences);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetUserPreferencesResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* teleport.userpreferences.v1.UserPreferences preferences = 1; */
        if (message.preferences)
            UserPreferences.internalBinaryWrite(message.preferences, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.userpreferences.v1.GetUserPreferencesResponse
 */
export const GetUserPreferencesResponse = new GetUserPreferencesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpsertUserPreferencesRequest$Type extends MessageType<UpsertUserPreferencesRequest> {
    constructor() {
        super("teleport.userpreferences.v1.UpsertUserPreferencesRequest", [
            { no: 1, name: "preferences", kind: "message", T: () => UserPreferences }
        ]);
    }
    create(value?: PartialMessage<UpsertUserPreferencesRequest>): UpsertUserPreferencesRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<UpsertUserPreferencesRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpsertUserPreferencesRequest): UpsertUserPreferencesRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* teleport.userpreferences.v1.UserPreferences preferences */ 1:
                    message.preferences = UserPreferences.internalBinaryRead(reader, reader.uint32(), options, message.preferences);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UpsertUserPreferencesRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* teleport.userpreferences.v1.UserPreferences preferences = 1; */
        if (message.preferences)
            UserPreferences.internalBinaryWrite(message.preferences, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.userpreferences.v1.UpsertUserPreferencesRequest
 */
export const UpsertUserPreferencesRequest = new UpsertUserPreferencesRequest$Type();
/**
 * @generated ServiceType for protobuf service teleport.userpreferences.v1.UserPreferencesService
 */
export const UserPreferencesService = new ServiceType("teleport.userpreferences.v1.UserPreferencesService", [
    { name: "GetUserPreferences", options: {}, I: GetUserPreferencesRequest, O: GetUserPreferencesResponse },
    { name: "UpsertUserPreferences", options: {}, I: UpsertUserPreferencesRequest, O: Empty }
]);
