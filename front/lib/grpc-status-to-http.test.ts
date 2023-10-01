import { expect } from '@jest/globals';
import grpcStatusToHttp from "@/lib/grpc-status-to-http";

describe("Test grpcStatusToHttp", () => {
    const data = [
        {grpc: 13, http: 400},
    ]

    for (const test of data) {
        it(`gRPC: ${test.grpc}`, () => {
            expect(grpcStatusToHttp(test.grpc)).toBe(test.http)
        })
    }
})