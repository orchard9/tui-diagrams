# Example Markdown with Mermaid Diagrams

This document demonstrates Mermaid diagram support in tui-diagrams.

## User Authentication Flow

Here's how user authentication works:

```mermaid
graph TD
    A[Start] --> B[Enter Credentials]
    B --> C{Valid Format?}
    C -->|Yes| D[Check Database]
    C -->|No| E[Show Error]
    D --> F{User Exists?}
    F -->|Yes| G[Verify Password]
    F -->|No| E
    G --> H{Match?}
    H -->|Yes| I[Generate Token]
    H -->|No| E
    I --> J[Success]
```

## API Request Flow

The sequence of events when making an API request:

```mermaid
sequenceDiagram
    Client->>API: POST /login
    API->>Database: Query user
    Database-->>API: User data
    API->>API: Verify password
    API-->>Client: JWT token
```

## CI/CD Pipeline

Our deployment pipeline:

```mermaid
graph LR
    A(Commit) --> B[Build]
    B --> C[Test]
    C --> D{Pass?}
    D -->|Yes| E[Deploy]
    D -->|No| F[Notify]
```

## Microservices Communication

```mermaid
sequenceDiagram
    participant User
    participant Gateway
    participant Auth
    participant Service
    User->>Gateway: Request
    Gateway->>Auth: Validate token
    Auth-->>Gateway: OK
    Gateway->>Service: Forward request
    Service-->>Gateway: Response
    Gateway-->>User: Response
```

## Decision Tree

```mermaid
flowchart TD
    Start[Receive Request] --> Auth{Authenticated?}
    Auth -->|No| Reject[Return 401]
    Auth -->|Yes| Role{Check Role}
    Role -->|Admin| Full[Full Access]
    Role -->|User| Limited[Limited Access]
    Role -->|Guest| ReadOnly[Read Only]
```
