/*
 * For a detailed explanation regarding each configuration property and type check, visit:
 * https://jestjs.io/docs/configuration
 */

export default {
    collectCoverage: true,
    coverageDirectory: 'coverage',
    coverageProvider: 'v8',
    testEnvironment: 'jsdom',
    preset: 'ts-jest',
    transform: {},
    moduleNameMapper: {
        '\\.(css)$': 'identity-obj-proxy',
    },
    collectCoverageFrom: [
        'src/pages/**/*.tsx',
        'src/utils/**/!(index).ts',
        'src/components/**/*.tsx',
        'src/hooks/!(index).ts',
    ],
    coverageThreshold: {
        global: {
            branches: 90,
            functions: 90,
            lines: 90,
        },
    },
    coverageReporters: ['text', 'html'],
};
