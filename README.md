# SSO Protos

This repository contains Protocol Buffers definitions and generated Go code for the Single Sign-On (SSO) service. The service provide authentication, user management, and application management functionality via gRPC.

## Related Projects

This repository is used as a dependency for the main project — [SSO Service](https://github.com/rshelekhov/sso) - gRPC service implementation

## Architecture

The SSO system is split into **3 services** for better separation of concerns:

- **AuthService** - Authentication and token management
- **UserService** - User profile and role management
- **AppService** - Application registration and management

## Proto Files Structure

```
proto/api/
├── auth/v1/auth.proto    # Authentication service
├── user/v1/user.proto    # User management service
└── app/v1/app.proto      # Application management service
```

## Generated Code Usage

Import the generated Go code for each service:

```go
// Authentication service
import authv1 "github.com/rshelekhov/sso-protos/gen/go/api/auth/v1"

// User management service
import userv1 "github.com/rshelekhov/sso-protos/gen/go/api/user/v1"

// Application management service
import appv1 "github.com/rshelekhov/sso-protos/gen/go/api/app/v1"
```

## Service Methods

### AuthService (Authentication & Token Management)

- `RegisterUser` - Register new user with email verification
- `VerifyEmail` - Verify user's email address
- `Login` - Authenticate user and receive tokens
- `Logout` - Terminate user session
- `RefreshTokens` - Refresh authentication tokens
- `RefreshTokensForApp` - Refresh tokens for specific application
- `CheckSession` - Verify session validity
- `GetJWKS` - Retrieve JSON Web Key Set for token verification
- `ResetPassword` - Initiate password reset process
- `ChangePassword` - Change user password
- `ConfirmCrossServicePasswordChange` - Confirm password change across services

### UserService (User Management)

- `GetUser` - Retrieve current user information
- `GetUserByID` - Retrieve user information by ID
- `UpdateUser` - Update user profile details
- `DeleteUser` - Delete current user account
- `DeleteUserByID` - Delete user account by ID

### AppService (Application Management)

- `RegisterApp` - Register new application in the SSO system

## Features

- **Multi-device support** with device tracking via UserDeviceData
- **JWT token management** with access/refresh token pairs
- **Email verification** workflow for new users
- **Password reset/change** functionality
- **Cross-service token refresh** for multi-app environments
- **JWKS endpoint** for token validation

## Code Generation

This repository uses [buf](https://buf.build) for Protocol Buffers code generation.

To generate code:

Install buf:

```bash
brew install bufbuild/buf/buf
```

Install dependencies:

```bash
make .bin-deps
```

Generate code:

```bash
buf generate .
```

## License

MIT License - see the LICENSE file for details
