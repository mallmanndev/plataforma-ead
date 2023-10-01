import getNameInitials from "@/lib/name-initials";
import {expect, it, describe} from '@jest/globals'

describe("Test getNameInitials", () => {
    const tests = [
        {name: "Matheus Mallmann", expected: "MM"},
        {name: "Matheus", expected: "MA"},
        {name: "Fulano de Tal", expected: "FT"},
    ]

    it.each(tests)(`Name: %s`, (test) => {
        const initial = getNameInitials(test.name)
        expect(initial).toBe(test.expected)
    })
})