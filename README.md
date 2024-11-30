# SSO Protos
This repository contains Protocol Buffers definitions and generated Go code for the Single Sign-On (SSO) service. The service provides authentication and user management functionality via gRPC and REST API endpoints.

## Related Projects
This repository is used as a dependency for the main project — [SSO Service](https://github.com/rshelekhov/sso) - gRPC service implementation

## Overview
The SSO service supports the following features:

- Application registration and management
- User registration and email verification
- Authentication (login/logout)
- Password management (reset/change)
- JWT token management with JWKS endpoint
- User profile management (get/update/delete)
- Token refresh functionality
- Multi-device support with device tracking

## Proto Files Structure
The main service definition includes:

- Auth service with all necessary RPC methods
- Request/Response messages for each method
- Common types (UserDeviceData, TokenData, JWK)

## Generated Code Usage
The generated Go code can be imported into your projects:

``` go
import "github.com/rshelekhov/sso-protos/gen/go/sso"
```

## Service Methods

### Application Management

- `RegisterApp` - Register new application in the SSO system

### User Management

- `RegisterUser` - Register new user with email verification
- `VerifyEmail` - Verify user's email address
- `GetUser` - Retrieve user information
- `UpdateUser` - Update user details
- `DeleteUser` - Delete user account

### Authentication

- `Login` - Authenticate user and receive tokens
- `Logout` - Terminate user session
- `Refresh` - Refresh authentication tokens
- `GetJWKS` - Retrieve JSON Web Key Set for token verification

### Password Management

- `ResetPassword` - Initiate password reset process
- `ChangePassword` - Change user password

## License

MIT License - see the LICENSE file for details