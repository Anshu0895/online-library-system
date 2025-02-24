module.exports = {
    testEnvironment: 'jsdom',
    setupFilesAfterEnv: ['<rootDir>/setupTests.js'],
    moduleNameMapper: {
      '\\.(css|less|sass|scss)$': 'identity-obj-proxy', // Mock CSS/SCSS files
      '\\.(gif|ttf|eot|svg)$': '<rootDir>/__mocks__/fileMock.js', // Mock static assets
    },
    transform: {
      '^.+\\.(js|jsx|ts|tsx)$': 'babel-jest', // Transform JavaScript/TypeScript files using Babel
    },
  };
  