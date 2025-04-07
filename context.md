# GoUpRoot - System Context and Architecture

## Project Overview
GoUpRoot is a specialized Linux privilege escalation enumeration tool written in Go. It's designed primarily for CTF (Capture The Flag) environments and security assessments, automating the discovery of potential privilege escalation vectors on Linux systems.

## Core Purpose
The tool serves as an automated reconnaissance utility that systematically checks various system components and configurations that could potentially be exploited for privilege escalation. It's inspired by popular tools like LinPEAS but implemented in Go for improved performance and maintainability.

## System Architecture

### Main Components
1. **Core Engine** (`main.go`)
   - Orchestrates the execution flow
   - Provides a clean CLI interface
   - Sequentially runs all enumeration modules

### Enumeration Modules
Located in the `modules/` directory, each focusing on specific attack vectors:

1. **Kernel Information** (`kernel.go`)
   - Identifies kernel version and potential vulnerabilities

2. **User Management** (`users.go`)
   - Enumerates system users
   - Checks user privileges and groups

3. **Environment Analysis** (`env.go`)
   - Examines environment variables
   - Identifies potential security misconfigurations

4. **File Permissions** (`fileperms.go`)
   - Scans for files with dangerous permissions
   - Identifies writable sensitive files

5. **SUID Binary Detection** (`suid.go`)
   - Locates SUID/SGID binaries
   - Identifies potentially exploitable setuid programs

6. **Cron Job Analysis** (`cron.go`)
   - Examines scheduled tasks
   - Identifies writable cron jobs

7. **Password Security** (`passwords.go`)
   - Checks for readable password files
   - Identifies potential password-related vulnerabilities

8. **Credential Search** (`creds.go`)
   - Searches for exposed credentials
   - Identifies config files with sensitive information

9. **Process Enumeration** (`processes.go`)
   - Lists running processes
   - Identifies processes running with elevated privileges

## Execution Flow
1. Display banner and initialization
2. System information gathering
3. File and permission checks
4. Credential and secret scanning
5. Process enumeration
6. Results presentation

## Technical Notes
- **Platform Support**: Currently Linux-only
- **Language**: Written in Go for performance and cross-compilation capabilities
- **Dependencies**: Minimal external dependencies for better security and portability
- **Execution Model**: Runs with current user privileges, no elevation required

## Security Considerations
- Tool is designed for authorized security assessments only
- Should only be run on systems with explicit permission
- Focuses on passive enumeration rather than active exploitation

## Project Structure
```
GoUpRoot/
├── main.go           # Main orchestration and execution
├── modules/         # Core enumeration modules
│   ├── kernel.go    # Kernel information checks
│   ├── users.go     # User enumeration
│   ├── env.go       # Environment analysis
│   ├── fileperms.go # File permission checks
│   ├── suid.go      # SUID binary detection
│   ├── cron.go      # Cron job analysis
│   ├── passwords.go # Password security checks
│   ├── creds.go     # Credential scanning
│   └── processes.go # Process enumeration
└── README.md        # Project documentation
```
